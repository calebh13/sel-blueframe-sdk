# Compatibility Matrix (Sprint 5)

This replaces the planned matrix in the [Sprint 4 compatibility plan](../Sprint4/02_Compatibility_Maintenance_Performance_Plan.md). Every "Tested" row below links to a log or a CI run. Rows that were not run say so.

Agreed baseline with the client: Python 3.12 (kickoff meeting, September 14). Supported operating systems: Linux x86_64 first, Windows x86_64 second. Blueframe devices are the deployment target and have not been tested.

## Python SDK over the Go client (bridge prototype, private repo)

| Platform | Python | Transport | Result | Evidence |
|---|---|---|---|---|
| Ubuntu 24.04 x86_64 | 3.12.3 | ctypes | 708 passed | Local run, October 30, 2026 |
| Ubuntu 24.04 x86_64 | 3.12.3 | cffi | 715 passed | Local run, October 30, 2026 |
| Ubuntu 24.04 x86_64 | 3.12.3 | bridge | 699 passed | Local run, October 30, 2026 |
| Ubuntu 24.04 x86_64 | 3.12.3 | shared tests | 39 passed | Local run, October 30, 2026 |
| Windows x86_64 | any | bridge | Not tested | Needs a Windows machine with the private repo |
| Linux aarch64 | any | bridge | Not tested | |
| Blueframe device | any | bridge | Not tested | Needs SEL hardware, see the [deployment plan](03_Deployment_Plan.md) |

Total 2,161 passed, 0 failed, in 300 s. numpy 2.5.4, grpcio 1.85.1, protobuf 7.36.2, pytest 9.2.0. The test log stays with the private repo because test names include client internals.

## Python.NET spike (public, `code/pythonnet-spike`)

Six tests, plus the timing script. [CI run 41144016907](https://github.com/calebh13/sel-blueframe-sdk/actions/runs/41144016907).

| Platform | Python | .NET | Tests | Call cost | Callbacks/s |
|---|---|---|---|---|---|
| Ubuntu 24.04 (local) | 3.12.3 | 10.0.14 | 6 passed | 2.7 to 2.9 us | 316k to 355k |
| ubuntu-latest (CI) | 3.10 | 10.0.x | 6 passed | | |
| ubuntu-latest (CI) | 3.12 | 10.0.x | 6 passed | | |
| ubuntu-latest (CI) | 3.13 | 10.0.x | 6 passed | | |
| windows-latest (CI) | 3.10 | 10.0.x | 6 passed | | |
| windows-latest (CI) | 3.12 | 10.0.x | 6 passed | 7.0 us | 122k |
| windows-latest (CI) | 3.13 | 10.0.x | 6 passed | | |

On Windows a fatal .NET error ends Python with exit code 2148734499 (0x80131623) instead of -6, so the test only checks that the exit code is not zero there. CI runners are shared machines, so their timings are rough.

## COMTRADE reader and writer (public, `code/`)

| Platform | Go | Result |
|---|---|---|
| ubuntu-latest | 1.21 | 3 tests passed |
| ubuntu-latest | stable | 3 tests passed |
| windows-latest | 1.21 | 3 tests passed |
| windows-latest | stable | 3 tests passed |

Same CI run as above.

## Still open

* The Go bridge on Windows. The Windows handshake fix is in, but nobody has run the suite there.
* SEL's C# client. Only the stand-in class has been tested with Python.NET ([issue #36](https://github.com/calebh13/sel-blueframe-sdk/issues/36)).
* numpy 1.26. Only numpy 2.x has been run.
* A real Signal Hub. Everything so far is against the fake hub.
