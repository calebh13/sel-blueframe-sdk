# Sprint 5 Report (Dates from Sprint 10/1 to Sprint 10/31)

## YouTube link of Sprint 5 Video
https://youtu.be/2oUGGHlFGZo

## What's New (User Facing)
 * The bridge transport survives a crash. With `auto_restart=True` a dead bridge is restarted, subscriptions pick up from the last timestamp with no gaps or duplicates, and publishers report the restart once
 * Clear Python errors for a duplicate client name, a forked child process, and a timeout, and a Python timeout now cancels the call in Go
 * A local demo app with publishers, the hub, and subscribers, including a button that kills the bridge to show recovery
 * Python.NET results for calling C# from Python ([results](../../code/pythonnet-spike/README.md))
 * First user guide with a pub/sub example that was run against the fake Signal Hub ([guide](../../docs/Sprint5/01_User_Guide.md))
 * Compatibility matrix with the agreed Python version and operating systems ([matrix](../../docs/Sprint5/02_Compatibility_Matrix.md))
 * The SDK and SEL's SignalHub client deployed against a Signal Hub server running locally, with the VS Code extension setting up the Field SDK devcontainer ([deployment](../../docs/Sprint5/03_Deployment_Plan.md))

## Work Summary (Developer Facing)
Sprint 5 was planned around building the Python bindings on SEL's C# client, since Joe agreed to that move on September 28. SEL's C# client was not on the team's machines, so the Go bridge stayed as the working SDK and the C# side was evaluated with a stand-in C# class. Lucas hardened the bridge in the first week of October, adding crash recovery, fork detection, call deadlines and duplicate name handling with 30 new end to end tests (#44), and built a demo app for the client (#45). Caleb tested Python.NET with the stand-in C# class (#46). It showed that generic methods and callbacks from a C# thread work from Python with no glue code, but a fatal .NET error still ends the Python session, so the bridge stays as the fallback. Caleb also wrote the CI workflow for the Go and Python.NET tests (#47) and collected the CI evidence (#52). Genevieve led the C# evaluation (#36), built the compatibility matrix (#49) and checked the user guide on a clean clone (#53). Darron wrote the user guide (#48), the cross-platform roadmap (#54), and the Sprint 5 report, slides and MoM files (#50). The team presented a Go and C# comparison to Joe on October 12, and he agreed to go forward with the C# port. The SDK was deployed with SEL's SignalHub client against a Signal Hub server running locally, and the full test suite of 2,161 tests passed on all three transports on October 30. The main lesson was that a dependency on code the team does not control needs a fallback from the start.

## Unfinished Work
All Sprint 5 issues were completed. Running the SDK on SEL's own C# client and signing it as a Blueframe app both need SEL's client and signing process, so they are planned for Sprint 6.

## Completed Issues/User Stories
Here are links to the issues that we completed in this sprint:

 * https://github.com/calebh13/sel-blueframe-sdk/issues/36
 * https://github.com/calebh13/sel-blueframe-sdk/issues/44
 * https://github.com/calebh13/sel-blueframe-sdk/issues/45
 * https://github.com/calebh13/sel-blueframe-sdk/issues/46
 * https://github.com/calebh13/sel-blueframe-sdk/issues/47
 * https://github.com/calebh13/sel-blueframe-sdk/issues/48
 * https://github.com/calebh13/sel-blueframe-sdk/issues/49
 * https://github.com/calebh13/sel-blueframe-sdk/issues/50
 * https://github.com/calebh13/sel-blueframe-sdk/issues/52
 * https://github.com/calebh13/sel-blueframe-sdk/issues/53
 * https://github.com/calebh13/sel-blueframe-sdk/issues/54
 * https://github.com/calebh13/sel-blueframe-sdk/issues/55

 ## Incomplete Issues/User Stories
 Here are links to issues we worked on but did not complete in this sprint:

 * None. All Sprint 5 issues were completed.

## Code Files for Review
Please review the following code files, which were actively developed during this sprint, for quality:
 * [spike.py](https://github.com/calebh13/sel-blueframe-sdk/blob/main/code/pythonnet-spike/spike.py)
 * [test_spike.py](https://github.com/calebh13/sel-blueframe-sdk/blob/main/code/pythonnet-spike/test_spike.py)
 * [LocalHub.cs](https://github.com/calebh13/sel-blueframe-sdk/blob/main/code/pythonnet-spike/csharp/LocalHub.cs)
 * [pubsub_example.py](https://github.com/calebh13/sel-blueframe-sdk/blob/main/docs/Sprint5/examples/pubsub_example.py)

## Retrospective Summary
Here's what went well:
  * Keeping the Go bridge as a fallback meant the sprint still ended with a working SDK.
  * Writing a test for each bridge failure before fixing it caught regressions right away.
  * The Go and C# comparison gave Joe the same three questions for both clients.

Here's what we'd like to improve:
   * The C# client should have been requested in writing in the first week, with a date.
   * Most Sprint 5 work sat on one laptop before it was pushed.
   * Issues for each member's work should be opened on day one.

Here are changes we plan to implement in the next sprint:
   * Get the C# client and a virtual Blueframe instance from Joe in the first week of Sprint 6.
   * Every member opens an issue for their work on day one and pushes code weekly.
   * Run the bridge on Windows and against a real Signal Hub.
