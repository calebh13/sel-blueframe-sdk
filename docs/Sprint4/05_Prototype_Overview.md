# Sprint 4 Prototype Overview

**Contributors:** Lucas Phillips (prototype), Caleb Hansen (repo setup and onboarding), Genevieve Kochel (source access and Python.NET research)  
**Status:** Prototype working in Go and Python. Reviewed with the client, who agreed the SDK should move to the C# client next.

> The prototype is built directly on SEL's SignalHub client, which is SEL Confidential, so its source stays in SEL's private repository. This page describes what was built without SEL source code or internal names.

## 1. What Was Built

The prototype is a working Python SDK over SEL's Go client. It was written in two languages. The Go side does everything that has to happen next to the client, and the Python side is what engineers import.

### Go code

| Component | What it does | Size |
|---|---|---|
| Go facade | The only code that calls SEL's client. Replaces generic Go calls with one plain call per operation, picks the concrete type from a 16 entry encoding table, buffers subscription data so Go never calls Python, recovers every panic as an error, and works around a few client behaviors. | 964 lines |
| C interface | Exposes the facade as 23 C functions so ctypes and cffi can load it as a shared library | 352 lines |
| gRPC bridge | A standalone Go program that serves the facade over 21 gRPC calls on localhost. Python starts it as a child process. The Go and Python stubs are generated from one `.proto` file. | 370 lines of Go and 126 lines of proto |
| Fake Signal Hub | One Go process that acts as the Hub, the signal directory, and its message bus, with in-memory history. Test hooks can send a fatal error, drop every connection, or make a signal never answer. | Test tool |
| Build script | Builds the shared library, the bridge, and the fake Hub with one command | Shell |

### Python code

| Component | What it does | Size |
|---|---|---|
| Python API | `connect`, signals, publishers, subscriptions, history, discovery, and rate helpers. numpy in and out, Python exceptions for every error. | 710 lines |
| ctypes transport | Loads the shared library with Python's standard library | 265 lines |
| cffi transport | Loads the same library with declarations read from the generated C header | 438 lines |
| Bridge transport | Starts the bridge and talks to it over gRPC | 225 lines |
| Test suite | Transport tests plus end to end tests that run against the fake Hub on every transport | 2,155 tests |

## 2. Feature Coverage

Every operation in the agreed scope works end to end from Python on all three transports.

| Area | Status |
|---|---|
| Connect and close, client info | Working |
| Signals for all 16 encodings, with and without metadata | Working |
| Fixed-rate, timestamped, and control publishing | Working |
| Single, multi-signal, time-aligned, and control subscriptions | Working |
| History fetch and latest sample | Working |
| Signal discovery | Working |
| Rate and time helpers | Working |
| Callbacks on a Python thread | Working |
| Lazy production (the Hub asking the client to compute a signal) | Deferred. Needs a Python round trip for every timestamp. |
| Connection events surfaced to Python | Deferred |

The full suite passed 2,155 of 2,155 tests on September 29, 2026. Joe reviewed the test run offline and said it meets his preferences.

## 3. Source Access (Genevieve)

The SignalHub client had been out of reach since Sprint 2. Genevieve got access to the C# and C++ versions and separated the Go client from the server code it lived in. That gave Caleb and Lucas a standalone client to build against.

## 4. Repo Setup and Onboarding (Caleb)

Caleb set up the extracted Go client in a private repository with its module path changed so it builds outside SEL's network. He also wrote Claude skills that describe the client's architecture and the design rules for binding it to Python, so the rest of the team could get up to speed on a large unfamiliar codebase faster.

## 5. Python.NET Research (Genevieve)

Genevieve looked at wrapping SEL's C# client with Python.NET instead of the Go client. .NET keeps type information at runtime, so Python.NET can call C# classes and generic methods directly without an export layer or generated handlers. That removes most of the glue code the Go prototype needed. On September 28 Joe agreed that moving to the C# implementation is the correct move. The work that carries over is the Python API design, the fake Hub, and the test suite.

## 6. Related Documents

- [Python Binding Requirements and Architecture Plan](01_Python_Binding_Architecture_Plan.md)
- [Compatibility, Maintenance, and Performance Plan](02_Compatibility_Maintenance_Performance_Plan.md)
- [Testing and Development Process Plan](03_Testing_and_Development_Process_Plan.md)
- [Documentation and Extensibility Plan](04_Documentation_and_Extensibility_Plan.md)
