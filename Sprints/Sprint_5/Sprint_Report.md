# Sprint 5 Report (Dates from Sprint 10/1 to Sprint 10/31)

## YouTube link of Sprint 5 Video (Make this video unlisted)
(add the unlisted link after recording)

## What's New (User Facing)
 * The bridge transport now survives a crash. With `auto_restart=True` a dead bridge is restarted, subscriptions pick up from the last timestamp with no gaps or duplicates, and publishers report the restart once ([#44](https://github.com/calebh13/sel-blueframe-sdk/issues/44))
 * Clear Python errors for a duplicate client name, a forked child process, and a timeout, and a Python timeout now cancels the call in Go ([#44](https://github.com/calebh13/sel-blueframe-sdk/issues/44))
 * A local demo app with publishers, the hub, and subscribers, including a button that kills the bridge to show recovery ([#45](https://github.com/calebh13/sel-blueframe-sdk/issues/45))
 * Python.NET spike with measured results for calling C# from Python ([#46](https://github.com/calebh13/sel-blueframe-sdk/issues/46), [results](../../code/pythonnet-spike/README.md))
 * A CI workflow that runs real tests on Linux and Windows and fails when a test fails ([#47](https://github.com/calebh13/sel-blueframe-sdk/issues/47))
 * First user guide with a pub/sub example that was run against the fake Signal Hub ([#48](https://github.com/calebh13/sel-blueframe-sdk/issues/48), [guide](../../docs/Sprint5/01_User_Guide.md))
 * Compatibility matrix with links to test results ([#49](https://github.com/calebh13/sel-blueframe-sdk/issues/49), [matrix](../../docs/Sprint5/02_Compatibility_Matrix.md))
 * Deployment plan for Blueframe devices ([plan](../../docs/Sprint5/03_Deployment_Plan.md))
 * The COMTRADE Go code now has a `go.mod`, so `go test` works without extra setup

## Work Summary (Developer Facing)
Sprint 5 was planned around building the Python bindings on SEL's C# client, since Joe agreed to that move on September 28. The C# client never made it onto the team's machines this sprint, so the hands-on C# evaluation in #36 did not start. The Go bridge from Sprint 4 was kept as the working SDK in the meantime. Lucas hardened it in the first week of October, working through the open findings from the September 28 review and adding crash recovery, fork detection, call deadlines and duplicate name handling, with 30 new end to end tests. He then built a demo app and prepared the Go half of a Go versus C# comparison talk. To get some C# evidence without the client, Lucas ran a Python.NET spike against a small stand-in C# class. It showed that generic methods and callbacks from a C# thread work from Python with no glue code, which was most of the Go work. It also showed that a fatal .NET error still kills the Python session, the same problem the ctypes transport had. While setting up CI, the team found that the old `ci.yml` only printed messages and passed even when a build failed, so it was replaced with a workflow that runs the COMTRADE and Python.NET tests on Linux and Windows. The full 2,161 test suite was rerun on October 30 and passed on all three transports. The biggest lesson was that a dependency on code we do not control (the C# client) needs a fallback plan from day one. Having the bridge ready meant the sprint still ended with a working, tested SDK.

### Team contributions

| Member | Sprint 5 work on the board |
|---|---|
| Lucas Phillips (lawja12) | #44 bridge hardening (8), #45 demo app (3), #46 Python.NET spike (3), #47 CI (3), #48 user guide (3), #49 compatibility matrix (2), #50 report and slides (3) |
| Caleb Hansen (calebh13) | Not recorded on the board yet. Caleb, add your issues and points here. |
| Genevieve Kochel (gigiellen10) | #36 C# evaluation (5), still open. Genevieve, add any progress here. |
| Darron Li (darronese) | Not recorded on the board yet. Darron, add your issues and points here. |

## Unfinished Work
The hands-on C# evaluation (#36) did not start because SEL's C# client was not available on the team's machines. The Python.NET spike covers the parts that do not need the client, and its results are in a comment on #36. The issue stays in the Sprint 5 milestone with that comment and moves to Sprint 6.

The planned Sprint 5 deliverables stand as follows.

| Deliverable (Sprint 4 to 6 plan) | Status |
|---|---|
| 1. Python bindings, pub/sub from an imported wrapper | Done on the Go bridge. Not started on the C# client. |
| 2. Automated test harness | Done. 2,161 tests on the bridge prototype and a public CI workflow on Linux and Windows. The prototype suite is not in CI yet because its code is private. |
| 3. Cross-platform and compatibility validation | Partly done. Python.NET spike and COMTRADE pass on Linux and Windows, Python 3.10 to 3.13. The Go bridge has not been run on Windows. |
| 4. Initial user documentation | Done as a first draft. |

Issues #1 and #3 to #7 (Python SDK wrapper, Discover, Subscribe and Publish) are still in the old "Blueframe SDK" milestone. The Go bridge already covers all of them, so they should be closed or moved to Sprint 6 for the C# version.

Nothing has been deployed to a Blueframe device or tested against a real Signal Hub. Both need SEL hardware or a virtual Blueframe instance and SEL's app signing.

## Completed Issues/User Stories
Here are links to the issues that we completed in this sprint:

 * https://github.com/calebh13/sel-blueframe-sdk/issues/44
 * https://github.com/calebh13/sel-blueframe-sdk/issues/45

These are finished and waiting for a reviewer on the Sprint 5 pull request. They close when it is merged.

 * https://github.com/calebh13/sel-blueframe-sdk/issues/46
 * https://github.com/calebh13/sel-blueframe-sdk/issues/47
 * https://github.com/calebh13/sel-blueframe-sdk/issues/48
 * https://github.com/calebh13/sel-blueframe-sdk/issues/49
 * https://github.com/calebh13/sel-blueframe-sdk/issues/50

 ## Incomplete Issues/User Stories
 Here are links to issues we worked on but did not complete in this sprint:

 * https://github.com/calebh13/sel-blueframe-sdk/issues/36 SEL's C# client was not available on the team's machines, so only the stand-in spike could be run.

## Code Files for Review
Please review the following code files, which were actively developed during this sprint, for quality:
 * [spike.py](https://github.com/calebh13/sel-blueframe-sdk/blob/sprint-5-pythonnet-spike/code/pythonnet-spike/spike.py)
 * [test_spike.py](https://github.com/calebh13/sel-blueframe-sdk/blob/sprint-5-pythonnet-spike/code/pythonnet-spike/test_spike.py)
 * [LocalHub.cs](https://github.com/calebh13/sel-blueframe-sdk/blob/sprint-5-pythonnet-spike/code/pythonnet-spike/csharp/LocalHub.cs)
 * [tests.yml](https://github.com/calebh13/sel-blueframe-sdk/blob/sprint-5-pythonnet-spike/.github/workflows/tests.yml)
 * [pubsub_example.py](https://github.com/calebh13/sel-blueframe-sdk/blob/sprint-5-pythonnet-spike/docs/Sprint5/examples/pubsub_example.py)

The bridge hardening and demo app code is in the private gosignals repository because it is built on SEL's confidential client.

## Testing and Validation

| What | How | Result |
|---|---|---|
| Python SDK on the Go bridge, all transports | Full suite against the fake Signal Hub, October 30 | 2,161 passed in 300 s (ctypes 708, cffi 715, bridge 699, shared 39) |
| Bridge crash recovery | 17 tests in `test_bridge_hardening.py`, including resume with no loss and no duplicates | Passed |
| Names and cancellation | 13 tests in `test_bridge_names_and_cancel.py` | Passed |
| Pub/sub as a user would write it | `docs/Sprint5/examples/pubsub_example.py` | 600 published, 600 received by callback, 600 fetched |
| Python.NET | 6 tests on Linux and Windows, Python 3.10, 3.12, 3.13 ([run](https://github.com/calebh13/sel-blueframe-sdk/actions/runs/41144016907)) | 6 of 6 passed on all 6 jobs |
| COMTRADE | `go test` on Linux and Windows, Go 1.21 and stable | 3 of 3 passed on all 4 jobs |

Python.NET spike measurements (Ubuntu, Python 3.12, .NET 10).

| Check | Result |
|---|---|
| Plain method call | 2.7 to 2.9 us |
| 1M samples, block copy | 20 to 135 million samples/s |
| 1M samples, element by element | about 4.6 million samples/s |
| Callbacks from a C# thread | 316k to 355k per second |
| C# exception | Python exception, session keeps running |
| Fatal .NET error | Python process exits with code -6 |

User testing with the client is not recorded yet. See the Sprint 5 MoMs once they are filled in.

## Responsible Use of GenAI
Claude (through Claude Code) was used this sprint to write the Python.NET spike and its tests, the CI workflow, the first drafts of the user guide, compatibility matrix and deployment plan, this report, and the slides. It was also used to rerun the test suite and pull numbers out of the logs. Meeting minutes were written from AI transcripts and cleaned up with Claude, as in Sprint 4.

The AI output was wrong in ways that only showed up when it was run. The first spike run failed because the C# source folder had the same name as the .NET namespace, so Python imported the folder instead. Every number in this report comes from a test log or a CI run, not from the AI's summary. During this review the team also found that the old `ci.yml` passed no matter what, which is why it is no longer cited as test evidence. SEL client code was not given to any public tool or put in this repo. The team reviewed every deliverable and is responsible for all of it.

## Achievements and Challenges
Achievements since Sprint 4.
  * The bridge went from a prototype to something that recovers from its own crash, with tests to prove it.
  * The test suite grew from 2,155 to 2,161 tests and still passes on all three transports.
  * CI now runs real tests, on two operating systems and three Python versions.
  * There is a first user guide and a working pub/sub example.

Challenges.
  * The C# client was not available. The bridge was kept as the working SDK and the spike measured what it could without the client.
  * There is still no real Signal Hub or Blueframe device. All testing is against the fake hub.
  * The SDK code is private, so it cannot run in the public CI.

## Retrospective Summary
Here's what went well:
  * Keeping the Go bridge as a fallback meant the sprint still ended with a working SDK.
  * Writing tests for each bridge failure before fixing it caught regressions right away.
  * Rerunning everything near the end of the sprint gave numbers the report can stand on.

Here's what we'd like to improve:
   * The C# client should have been requested in writing in the first week, with a date.
   * Most Sprint 5 work sat on one laptop and only one member's work shows on the board.
   * The old CI should have been checked much earlier.

Here are changes we plan to implement in the next sprint:
   * Get the C# client and a virtual Blueframe instance from Joe in the first week of Sprint 6.
   * Every member opens an issue for their work on day one, and pushes code weekly.
   * Add the private prototype suite to CI in the private repo.
