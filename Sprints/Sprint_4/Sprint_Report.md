# Sprint 4 Report (Dates from Sprint 8/27 to Sprint 9/30)

## YouTube link of Sprint 4 Video
https://www.youtube.com/watch?v=aWu-euDfJv8

## What's New (User Facing)
 * Working Python prototype that publishes, subscribes, fetches history, and discovers signals through SEL's Go SignalHub client, covering every operation in the agreed scope ([overview](../../docs/Sprint4/05_Prototype_Overview.md))
 * Go code for the prototype, made up of a facade over SEL's client, a C interface with 23 functions, a gRPC bridge with 21 calls, and a fake Signal Hub for testing
 * Side by side comparison of three ways to call the Go client from Python (ctypes, cffi, and a gRPC subprocess bridge)
 * Fake Signal Hub and a 2,155 test suite that passes on all three transports
 * [Python Binding Requirements and Architecture Plan](../../docs/Sprint4/01_Python_Binding_Architecture_Plan.md) with a proposed Python API
 * [Compatibility, Maintenance, and Performance Plan](../../docs/Sprint4/02_Compatibility_Maintenance_Performance_Plan.md) with benchmark results
 * [Testing and Development Process Plan](../../docs/Sprint4/03_Testing_and_Development_Process_Plan.md)
 * [Documentation and Extensibility Plan](../../docs/Sprint4/04_Documentation_and_Extensibility_Plan.md)
 * C# with Python.NET chosen as the base for the SDK, with the client's agreement

## Work Summary (Developer Facing)
Sprint 4 was planned as an investigation sprint, and most of the research behind the Python SDK is now done. The access problem that blocked Sprints 2 and 3 was cleared early on. Genevieve got access to the C# and C++ versions of the client and pulled the Go client out of the server code, and she also set up our meetings with Joe. Caleb put the Go client in a private repo and wrote Claude skills to speed up onboarding to it. He then worked through why the original CGo plan breaks down for this client, looking at object lifetimes, callbacks, concurrency, and generics. Lucas took the bridge research further and built a working prototype with three transports behind one Python API, a fake Signal Hub, and a test suite that runs on every transport. That work also showed why the Go codebase is a poor fit for wrapping, since most of the effort went into glue code instead of features. Genevieve's Python.NET research showed the C# client avoids most of that glue. Joe was very happy with the research and agreed that moving to the C# implementation is the correct move, and after the sprint's final test run he told us it meets his preferences. Darron Li spearheaded cross-functional alignment on the documentation strategy, proactively socializing best-in-class conventions and driving stakeholder synergy across the extensibility roadmap. The biggest lesson for the team was that the approach that looked easiest on paper, a script that adds export lines to Go functions, was the most fragile one once it was tested.

## Unfinished Work
Hands-on work with the C# client has not started. Genevieve's Python.NET research is done and Joe agreed on the move to C#, but running the C# client against the fake Signal Hub and repeating the benchmarks did not fit in the sprint. That issue has a comment explaining why and was moved to the Sprint 5 milestone. The prototype has also not been run against a real Signal Hub or on Windows yet. Both are part of the Sprint 5 compatibility work.

## Completed Issues/User Stories
Here are links to the issues that we completed in this sprint:

There are 11 completed issues (#25 to #35). The Sprint 4 milestone page on GitHub shows 18 closed because it also counts the 7 pull requests.

 * https://github.com/calebh13/sel-blueframe-sdk/issues/25
 * https://github.com/calebh13/sel-blueframe-sdk/issues/26
 * https://github.com/calebh13/sel-blueframe-sdk/issues/27
 * https://github.com/calebh13/sel-blueframe-sdk/issues/28
 * https://github.com/calebh13/sel-blueframe-sdk/issues/29
 * https://github.com/calebh13/sel-blueframe-sdk/issues/30
 * https://github.com/calebh13/sel-blueframe-sdk/issues/31
 * https://github.com/calebh13/sel-blueframe-sdk/issues/32
 * https://github.com/calebh13/sel-blueframe-sdk/issues/33
 * https://github.com/calebh13/sel-blueframe-sdk/issues/34
 * https://github.com/calebh13/sel-blueframe-sdk/issues/35
 
 ## Incomplete Issues/User Stories
 Here are links to issues we worked on but did not complete in this sprint:
 
 * https://github.com/calebh13/sel-blueframe-sdk/issues/36 We did not get to the hands-on work because the Python.NET research took most of the sprint, so running the C# client against the fake Signal Hub moved to Sprint 5.

## Code Files for Review
Please review the following code files, which were actively developed during this sprint, for quality:

Note: the source code for the Python SDK prototype is private. It is built on SEL's confidential SignalHub client and lives in SEL's private gosignals repository, so only the documents below are in this public repo.
 * [Python Binding Requirements and Architecture Plan](https://github.com/calebh13/sel-blueframe-sdk/blob/main/docs/Sprint4/01_Python_Binding_Architecture_Plan.md)
 * [Compatibility, Maintenance, and Performance Plan](https://github.com/calebh13/sel-blueframe-sdk/blob/main/docs/Sprint4/02_Compatibility_Maintenance_Performance_Plan.md)
 * [Testing and Development Process Plan](https://github.com/calebh13/sel-blueframe-sdk/blob/main/docs/Sprint4/03_Testing_and_Development_Process_Plan.md)
 * [Documentation and Extensibility Plan](https://github.com/calebh13/sel-blueframe-sdk/blob/main/docs/Sprint4/04_Documentation_and_Extensibility_Plan.md)
 * [Sprint 4 Prototype Overview](https://github.com/calebh13/sel-blueframe-sdk/blob/main/docs/Sprint4/05_Prototype_Overview.md)
 
## Retrospective Summary
Here's what went well:
  * Access to the client code finally came through, which unblocked work that had been stuck since Sprint 2.
  * The prototype answered Joe's main question with test results instead of guesses.
  * Building all three transports behind one API made the comparison fair and left us with a test suite we can keep using.
 
Here's what we'd like to improve:
   * The CGo plan was assumed to be easy and took time to disprove. A small spike at the start of the sprint would have caught it sooner.
   * Prototype work sat on one laptop for most of the sprint before it was pushed anywhere.
   * Nothing has been tested against a real Signal Hub, so some findings may not hold up in deployment.
  
Here are changes we plan to implement in the next sprint:
   * Start moving the SDK onto the C# client in the first week of Sprint 5.
   * Push prototype work to the private repo as it happens.
   * Ask Joe for access to a real Signal Hub and a Windows machine for testing.
