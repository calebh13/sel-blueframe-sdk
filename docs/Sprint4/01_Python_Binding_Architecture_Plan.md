# Python Binding Requirements and Architecture Plan

**Sprint 4 Deliverable 1**  
**Owner:** Lucas Phillips  
**Contributors:** Genevieve Kochel (C# and Python.NET research)  
**Status:** Reviewed with the client on September 28, 2026

> The SignalHub client code is SEL Confidential. This document describes the approach and results without SEL source code or internal names. The prototype code is kept in the team's private repository.

## 1. Purpose

This plan covers how Python programs will call the existing SignalHub client. It lists the requirements from the client, the approaches that were tried, the proposed Python API, how data types and errors cross the language boundary, and what stays hidden from users.

## 2. Requirements

These come from the September 14 client meeting and the Sprint 4 to 6 project plan.

| # | Requirement | Source |
|---|---|---|
| R1 | Wrap the existing client. Do not write a Python port of the protocol. | Client, Sept 14 |
| R2 | Avoid double maintenance when the underlying client changes. | Client, Sept 14 |
| R3 | Leave the Go client code unmodified. | Team constraint |
| R4 | Support publish and subscribe as the main use case. | Client, Sept 14 |
| R5 | Target Python 3.12. | Client, Sept 14 |
| R6 | Hide Go details from SDK users. | Client, Sept 14 |
| R7 | Make the wrapper extensible by future developers. | Client, Sept 14 |
| R8 | A failure in the wrapped code should raise a Python error, not crash the user's program. | Team, from prototype results |

## 3. Approaches Evaluated

Three ways of calling the Go client from Python were built and tested behind the same Python API.

| | ctypes | cffi | Subprocess bridge (gRPC) |
|---|---|---|---|
| How it works | Go built as a C shared library and loaded with Python's standard ctypes module | Same shared library, loaded with cffi using declarations read from the generated header | Go runs as a child process and Python talks to it over gRPC on localhost |
| Works for common operations | Yes | Yes | Yes |
| Bulk throughput | 13 to 15 M samples/s | 13 to 15 M samples/s | 16 to 17 M samples/s |
| Cost per call | about 2 microseconds | about 1.5 microseconds | about 17 microseconds, so data is sent in batches |
| Median latency | 5.2 ms | 5.2 ms | 5.5 ms |
| Go crash, panic, or exit | Python process dies | Python process dies | Python gets an exception and can start a new bridge |
| Fork-based worker pools | Hang | Hang | Work if each worker makes its own client |
| Keeping declarations in sync | Written by hand | Read from the C header | Generated from one .proto file |
| Go changes needed to be safe | An exit hook in the Go code | An exit hook in the Go code | None |

For reference, the Go client alone reached about 23 M samples/s in the same benchmark. Latency is set by the client's 10 ms send interval, not by the wrapper.

The earlier plan of adding export lines to Go functions with a script and calling them through CGo does not work for this client. Caleb's findings on why are in the [compatibility plan](02_Compatibility_Maintenance_Performance_Plan.md).

## 4. Proposed Architecture

The design has five layers. Only the transport layer changes between the three approaches.

| Layer | Responsibility |
|---|---|
| 1. Python API | What engineers import. Takes and returns numpy arrays. Raises Python exceptions. |
| 2. Transport | Moves calls and data between Python and Go. The recommended transport is the subprocess bridge. |
| 3. Go facade | The only code that calls the SEL client. Turns generic Go calls into one call per operation, buffers subscription data, and turns panics into errors. |
| 4. SignalHub Go client | SEL's existing client, unchanged. |
| 5. Signal Hub | The server. In testing this is a fake Signal Hub written by the team. |

Three design rules removed most of the problems found in the first prototype.

1. Go never calls into Python. Subscription data is buffered in Go and Python pulls it in batches. User callbacks run on a Python thread.
2. Control information travels as structured messages and sample data travels as raw binary records. The same format works for all three transports.
3. No call blocks for more than 250 ms at a time, so Ctrl+C works within about half a second.

## 5. Proposed Python API

```python
import gosignals as gs

with gs.connect("127.0.0.1", 49000, name="relay-study", transport="bridge") as hub:
    va = gs.Signal("SUB1.BUS1", "VA", rate=60, dtype="float32")

    with hub.publisher(va, start=t0) as pub:
        pub.publish(t0, values)                   # numpy array

    ts, vals = hub.fetch(va, start=t0, end=t1)    # history as numpy arrays

    with hub.subscribe(va, start=t0) as sub:      # live data
        for ts, vals in sub:
            ...

    last = hub.previous_sample(va, lookback=600)
```

| Area | Python calls |
|---|---|
| Connecting | `gs.connect(...)`, `client.close()`, context manager |
| Signals | `gs.Signal(...)`, `Signal.describe()` |
| Publishing | `publisher`, `timestamped_publisher`, `control_publisher` |
| Subscribing | `subscribe`, `subscribe_many`, `subscribe_aligned`, `subscribe_control` |
| History | `fetch`, `previous_sample` |
| Discovery | `monitor_signals` |
| Rates and time | `client.rates` helpers, `client.now()` |

The 34 exported types in the client's core package collapse into a handful of Python classes. Most of the Go types exist to satisfy Go's type system and do not represent separate ideas for a user.

## 6. Data Types

- The client supports eight numeric sample types, each with and without quality metadata, for 16 encodings in total. Each maps to a numpy dtype.
- Metadata samples are numpy structured records with value, trigger time, time quality, and quality fields.
- Timestamps are int64 nanoseconds since the Unix epoch. Float seconds are not used for storage because a float64 cannot hold nanosecond timestamps exactly.
- Objects with a lifetime (clients, publishers, subscriptions) are handles on the Go side and Python objects with `close()` on the Python side.
- Plain records such as signal descriptions are copied by value.

## 7. Error Handling

| Python exception | Raised when |
|---|---|
| `GoSignalsError` | The Go client returns an error or a panic is recovered |
| `TransportError` | The transport fails, for example when the bridge process exits |
| `CallbackError` | A user callback raised an exception. It is reported on close while delivery continues. |
| `ValueError` / `TimeoutError` | Bad input from the user, or a call that ran out of time |

Every Go call recovers panics and returns them as errors. Go stack traces are not shown to users.

## 8. What Is Hidden from Users

- Go generics and the per-type dispatch table
- Handles, the transport, and the bridge process
- The wire protocol and packet timing
- Go logging, which will be routed into Python's `logging` module
- Workarounds for client behavior that the wrapper handles internally

## 9. Alternative Using the C# Client and Python.NET

Genevieve's research looked at wrapping SEL's C# client with Python.NET instead of the Go client. Several of the hardest Go problems do not exist in .NET.

| Problem with Go | How .NET compares |
|---|---|
| Generics only exist at compile time | .NET generics exist at runtime, so Python can pick the type when it makes the call |
| Functions cannot be called by name | Python.NET uses reflection to expose public classes and methods without an export layer |
| Objects cannot cross into Python | .NET objects are passed to Python as references |
| Panics and exits kill Python | Exceptions on the calling thread become Python exceptions |

Some problems carry over. Forking after the .NET runtime loads is also unsafe. An unhandled exception on a .NET background thread still ends the process. Engineers would need a .NET runtime installed. None of this has been tested against the real C# client yet.

## 10. Recommendation

1. Evaluate the C# client with Python.NET early in Sprint 5. This is a one to two day check against the fake Signal Hub, using the same benchmarks and failure tests.
2. If the C# client passes, build the SDK on it.
3. If the team stays with Go, build on the subprocess bridge. ctypes and cffi should not be used for production.
4. Either way, keep the Python API design, the fake Signal Hub, and the test suite. None of them depend on Go.

## 11. Unresolved Questions

- Does the C# client run on Linux with a modern .NET version?
- Can target machines start a child process for the bridge?
- Which operating systems do engineers use day to day?
- Is a real Signal Hub available for testing before Sprint 6?
- Does lazy production (the client computing signals on request) need to be supported? It needs a round trip to Python for every timestamp and is deferred for now.

## 12. Evidence

- Prototype SDK with three transports and a shared Python API (private repository)
- Test suite of 2,155 tests, all passing on September 29, 2026 ([testing plan](03_Testing_and_Development_Process_Plan.md))
- Benchmark and failure test results from the prototype
- [MoM September 14](../../Sprints/Sprint_4/MoM/MoM_2026-09-14.md) and [MoM September 28](../../Sprints/Sprint_4/MoM/MoM_2026-09-28.md)
