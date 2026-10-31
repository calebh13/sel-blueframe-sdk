# Sprint 5 Report (Dates from Sprint 10/1 to Sprint 10/31)

## YouTube link of Sprint 5 Video (Make this video unlisted)
[Video link placeholder, to be added after recording]

## What's New (User Facing)
 * The bridge transport survives a crash. With `auto_restart=True` a dead bridge is restarted, subscriptions pick up from the last timestamp with no gaps or duplicates, and publishers report the restart once
 * Clear Python errors for a duplicate client name, a forked child process, and a timeout, and a Python timeout now cancels the call in Go
 * A local demo app with publishers, the hub, and subscribers, including a button that kills the bridge to show recovery
 * Python.NET spike results for calling C# from Python ([results](../../code/pythonnet-spike/README.md))
 * First user guide with a pub/sub example that was run against the fake Signal Hub ([guide](../../docs/Sprint5/01_User_Guide.md))
 * Compatibility matrix with the agreed Python version and operating systems ([matrix](../../docs/Sprint5/02_Compatibility_Matrix.md))
 * Deployment plan for Blueframe devices ([plan](../../docs/Sprint5/03_Deployment_Plan.md))

## Work Summary (Developer Facing)
Sprint 5 was planned around building the Python bindings on SEL's C# client, since Joe agreed to that move on September 28. The C# client never made it onto the team's machines, so Genevieve's hands-on evaluation (#36) could not start and the Go bridge stayed as the working SDK. Lucas hardened the bridge in the first week of October, adding crash recovery, fork detection, call deadlines and duplicate name handling with 30 new end to end tests, and built a demo app for the client. The team presented a Go and C# comparison to Joe on October 12. To get C# evidence without the client, Lucas ran a Python.NET spike with a stand-in C# class. It showed that generic methods and callbacks from a C# thread work from Python with no glue code, but a fatal .NET error still ends the Python session, so the bridge stays as the fallback. The team CI pipeline was run on the Sprint 5 branch, and the full test suite of 2,161 tests passed on all three transports on October 30. Claude was used for the spike, document drafts, and slides, and for cleaning up meeting transcripts, and the team checked every result against test logs. The main lesson was that a dependency on code the team does not control needs a fallback from the start.

## Unfinished Work
The hands-on C# evaluation (#36) did not start because SEL's C# client was not available on the team's machines. Progress is tracked in a comment on the issue, which explains why it could not be completed, and it moves to Sprint 6. The Python bindings work on the Go bridge but not yet on the C# client. The bridge has not been run on Windows or against a real Signal Hub, and nothing has been deployed to a Blueframe device, since that needs SEL hardware and app signing.

## Completed Issues/User Stories
Here are links to the issues that we completed in this sprint:

 * https://github.com/calebh13/sel-blueframe-sdk/issues/44
 * https://github.com/calebh13/sel-blueframe-sdk/issues/45
 * https://github.com/calebh13/sel-blueframe-sdk/issues/46
 * https://github.com/calebh13/sel-blueframe-sdk/issues/48
 * https://github.com/calebh13/sel-blueframe-sdk/issues/49
 * https://github.com/calebh13/sel-blueframe-sdk/issues/50

 ## Incomplete Issues/User Stories
 Here are links to issues we worked on but did not complete in this sprint:

 * https://github.com/calebh13/sel-blueframe-sdk/issues/36 SEL's C# client was not available on the team's machines, so only the stand-in spike could be run.
 * https://github.com/calebh13/sel-blueframe-sdk/issues/52 Caleb's review of the CI evidence was still open at the end of the sprint.
 * https://github.com/calebh13/sel-blueframe-sdk/issues/53 Genevieve's walkthrough of the user guide on a clean machine was still open at the end of the sprint.
 * https://github.com/calebh13/sel-blueframe-sdk/issues/54 Darron's cross-platform roadmap work carries into Sprint 6, alongside the Windows run of the bridge.
 * https://github.com/calebh13/sel-blueframe-sdk/issues/55 The video is recorded after the report and slides are final.

## Code Files for Review
Please review the following code files, which were actively developed during this sprint, for quality:
 * [spike.py](https://github.com/calebh13/sel-blueframe-sdk/blob/sprint-5-pythonnet-spike/code/pythonnet-spike/spike.py)
 * [test_spike.py](https://github.com/calebh13/sel-blueframe-sdk/blob/sprint-5-pythonnet-spike/code/pythonnet-spike/test_spike.py)
 * [LocalHub.cs](https://github.com/calebh13/sel-blueframe-sdk/blob/sprint-5-pythonnet-spike/code/pythonnet-spike/csharp/LocalHub.cs)
 * [pubsub_example.py](https://github.com/calebh13/sel-blueframe-sdk/blob/sprint-5-pythonnet-spike/docs/Sprint5/examples/pubsub_example.py)

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
