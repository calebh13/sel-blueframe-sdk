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
| Windows x86_64 | any | bridge | Builds, tests not run | Cross-compiled October 31, 2026, see the roadmap below |
| Linux aarch64 | any | bridge | Builds, tests not run | Cross-compiled October 31, 2026, see the roadmap below |
| Blueframe device | any | bridge | Not tested | Needs SEL hardware, see the [deployment plan](03_Deployment_Plan.md) |

Total 2,161 passed, 0 failed, in 300 s. numpy 2.5.4, grpcio 1.85.1, protobuf 7.36.2, pytest 9.2.0. The test log stays with the private repo because test names include client internals. The [run summary](evidence/suite_summary.png) only has the totals, and the [pub/sub example output](evidence/pubsub_example_output.txt) is from the same day. The team CI pipeline also passed on the Sprint 5 branch ([run 41146542375 log](evidence/ci_run_41146542375.log)).

## Python.NET spike (public, `code/pythonnet-spike`)

Six tests, plus the timing script. [CI run 41144016907](https://github.com/calebh13/sel-blueframe-sdk/actions/runs/41144016907). The local numbers come from [three timing runs](evidence/pythonnet_spike_runs.txt) and a [fresh clone test](evidence/fresh_clone_readme_test.txt).

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

## Roadmap for Windows x86_64 and Linux aarch64

The bridge and the fake hub are plain Go, so they can be cross-compiled from Linux. These builds were done on October 31, 2026 with Go 1.26.8 and `CGO_ENABLED=0`, from `gosignals` main at commit 49400d6.

| Target | Bridge | Fake hub |
|---|---|---|
| windows/amd64 | Builds | Fails because `testhub/main.go` uses `syscall.SIGUSR1` and `SIGUSR2`, which Windows does not have |
| linux/arm64 | Builds | Builds |

Steps for Sprint 6, in order.

1. Move the `SIGUSR1` and `SIGUSR2` handling in `testhub/main.go` into a file that only builds on Linux, so the fake hub builds for Windows. Until then the hub can run on a Linux machine and the Windows tests can connect to it over the network.
2. Run `tests/e2e` and `tests/transport/test_bridge_transport.py` with `GOSIGNALS_TRANSPORTS=bridge` on a Windows x86_64 machine with Python 3.12, then 3.10 and 3.13.
3. Run the same tests on a Linux aarch64 machine.
4. Add each run to the first table with its date and pass count.

The ctypes and cffi transports are left out. They load Go into Python through cgo, so each target would need a C cross compiler, and the user guide already tells engineers to use the bridge.

## Still open

* The Go bridge test suite on Windows and Linux aarch64. Both builds work, but no tests have been run there yet.
* SEL's C# client. Only the stand-in class has been tested with Python.NET ([issue #36](https://github.com/calebh13/sel-blueframe-sdk/issues/36)).
* numpy 1.26. Only numpy 2.x has been run.
* A real Signal Hub. Everything so far is against the fake hub.
