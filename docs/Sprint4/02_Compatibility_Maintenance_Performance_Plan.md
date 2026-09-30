# Compatibility, Maintenance, and Performance Plan

**Sprint 4 Deliverable 2**  
**Owner:** Caleb Hansen  
**Contributors:** Lucas Phillips (benchmarks and failure tests)  
**Status:** Reviewed with the client on September 28, 2026. Joe agreed that moving to the C# implementation is the correct move.

> The SignalHub client code is SEL Confidential. This plan describes findings and targets without SEL source code or internal names.

## 1. Purpose

This plan sets the Python and OS targets for the wrapper, explains how multithreading and performance were checked, and lays out how the wrapper can be maintained without duplicating the Go client. It also records the CGo investigation that changed the team's original plan.

## 2. CGo Investigation

The first plan was to add `//export` lines to Go functions with a small script, build the client as a C shared library, and call it from Python. Several problems made that unworkable for this client.

| Problem | What happens | What the wrapper has to do instead |
|---|---|---|
| The whole Go runtime loads into Python | Building with `-buildmode=c-shared` does not turn Go into C. The shared library carries its own garbage collector and scheduler, which start when Python loads it. | Treat the library as a separate runtime living inside Python, with its own threads |
| Generics cannot be exported | `//export` on a generic function is a compile error because there is no single function to export | Write one non-generic entry point per operation and pick the concrete type with a switch on a type code |
| Callbacks run on Go threads | A Go thread calling into Python has to take the GIL. At high data rates that means tens of thousands of GIL handoffs per second, and a slow Python callback stalls the thread reading from the Hub socket. | Buffer data in Go and let Python pull it in batches |
| Go objects cannot be handed to Python | Once a Go function returns a pointer to C, nothing on the Go side references it, so the garbage collector can free it. Pinning does not help because the client object holds many other Go pointers. | Keep objects in a table on the Go side and give Python integer handles |
| Object lifetimes | Go's garbage collector frees memory but does not close sockets or stop goroutines. Python code can also close things twice or use them after closing. | Explicit open and close with a defined order. Double close and use after close return errors. |
| Panics | The client panics on some ordinary input mistakes. A panic that reaches a C frame aborts the Python interpreter with no traceback. | Recover panics at every exported call and return them as errors |

Most of these problems also apply to a subprocess bridge. That is why the prototype puts them in a shared Go facade that every transport uses. The difference is that under CGo a mistake crashes Python with a segfault, while in a separate process it becomes a Python exception.

## 3. Compatibility Matrix

| Component | Target | Status |
|---|---|---|
| Python | 3.12 (client baseline) | Tested |
| Python | 3.10, 3.11, 3.13 | Planned for Sprint 5 |
| Python | 3.14 | Planned. Its default multiprocessing start method changed on Linux, which helps the fork issue in Section 4. |
| numpy | 1.26 and 2.x | 2.x tested, 1.26 planned |
| Go toolchain for building the wrapper | 1.25 | Tested |
| Go client | Unmodified SEL client | Tested against a fake Signal Hub only |
| Linux x86_64 | Primary platform | Tested |
| Linux aarch64 | Supported | Planned. The bridge cross-compiles without a C toolchain. |
| Windows x86_64 | Supported | Planned. The bridge builds for Windows from Linux. ctypes and cffi would need a MinGW toolchain on Windows. |
| macOS | Developer machines only | Not planned unless the client asks |
| Blueframe devices | Deployment target | Open question. Depends on whether apps can start a child process. |

The subprocess bridge is the easiest option to keep OS-independent. It is a plain Go binary with no C code, so every platform build can be made from one Linux machine and shipped inside a Python wheel.

## 4. Multithreading

- **GIL.** Go never calls Python. Data is buffered in Go and pulled by Python in batches, so Go threads never wait on the GIL.
- **User callbacks** run on a dedicated Python thread. An exception in a callback is collected and raised on close while delivery keeps going.
- **Thread safety.** The test suite runs 40 subscriptions and 40 publishers on one client, 25 publishers across 5 clients, and 16 threads sharing one client. All pass.
- **Races inside the client.** Go's race detector reports races between the client's public calls and its own background goroutines. Any Go program using the client hits the same races, so they are client behavior. The wrapper will lock each object it owns and run the race detector on its own code.
- **fork().** The Go runtime does not survive `fork()`. With ctypes and cffi, fork-based worker pools hang. With the bridge, each worker has to create its own client. The documented rule will be to use the `spawn` start method or create the client inside each worker.

## 5. Performance

Measured on the prototype over about 30 seconds per transport, two runs, on a 16 core Linux machine against a fake Signal Hub. Signals were float32 at a synthetic 1 MHz rate, sent in bulk publishes of 100,000 samples.

| Scenario | ctypes | cffi | Bridge | Go alone |
|---|---|---|---|---|
| Bulk publish to subscribe (samples/s) | 13.4 to 14.8 M | 13.5 to 14.9 M | 16.5 to 17.2 M | 23.3 M |
| One sample per call | 2.0 microseconds | 1.45 microseconds | 17 microseconds | not measured |
| Latency p50 / p99 | 5.2 / 10.3 ms | 5.2 / 10.2 ms | 5.5 / 10.6 ms | not measured |
| Memory (RSS) | about 320 MB | about 310 MB | 130 MB + 190 MB | not measured |

Latency is set by the client's 10 ms send interval. The bridge adds about 0.3 ms. For comparison, 500 signals at 60 Hz is about 120 KB/s of data, so throughput is far from being the bottleneck.

### Proposed targets

These are proposals since the client left benchmarks for the team to set.

| Metric | Target |
|---|---|
| Bulk throughput | At least 50% of Go alone |
| Added latency over Go alone | Under 1 ms at p50 |
| Samples dropped under normal load | None, and any drop is counted and reported |
| Time for Ctrl+C to take effect | Under 1 second |
| Memory per idle client | To be measured in Sprint 5 |

## 6. Maintenance Strategy

The point of the project is to avoid a second copy of the client. The plan has four parts.

1. **Wrap and do not port.** All protocol logic stays in SEL's client. The wrapper only translates calls and data.
2. **One facade.** A single Go facade is the only code that calls the client. When the client changes, the Go compiler flags the facade at build time.
3. **Generated interfaces.** The bridge protocol is one `.proto` file that generates both the Go and Python sides, so the two cannot silently disagree. A planned generator would read an explicit list of functions to expose and produce the handlers, so adding a function with a known shape is one entry and a rebuild.
4. **Drift checks in CI.** A committed snapshot of the exposed API will show a diff whenever the client changes. Canary tests will fail if a client behavior the wrapper works around gets fixed, so the workaround can be removed.

Estimated effort for a new operation today is 40 to 60 lines, mostly mechanical. With the generator it should be a single entry for common shapes.

## 7. Risks

| Risk | Where it shows up | Mitigation |
|---|---|---|
| A crash in the Go code | Any use of ctypes or cffi | Use the bridge |
| Terminal Ctrl+C also stopped the bridge | Interactive use | Fixed. The bridge now starts in its own process group. |
| Duplicate client names retry forever | Re-running a Jupyter cell | Clear error on connect and an optional unique suffix |
| Fake Signal Hub differs from the real one | First real deployment | Contract tests against a real Hub |
| Windows console signals and antivirus flags | Windows users | Job Object for cleanup and a signed bridge binary |
| Client update changes a behavior the wrapper relies on | Client releases | API snapshot and canary tests |

## 8. Open Questions

- Can Blueframe apps start child processes?
- Which Windows versions do engineers use?
- Is a real Signal Hub available for compatibility testing?
- The compatibility matrix needs to be redone for the C# client and the .NET runtime in Sprint 5, now that Joe has agreed to move to it.
