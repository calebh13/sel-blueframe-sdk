# Sprint 3 Report (Dates from Sprint 4/10 to Sprint 5/1)

## YouTube link of Sprint * Video (Make this video unlisted)

## What's New (User Facing)
 * A VSCode extension that provides a one-click installation and deployment process for the Blueframe Field SDK.
 * New user interface buttons in VSCode to automate environment tasks such as starting or stopping the local environment.
 * A functional Python interface for parsing and writing IEEE COMTRADE files.
 * Detailed troubleshooting guides and technology manuals for setting up the Blueframe development environment.

## Work Summary (Developer Facing)
This sprint focused on building the foundational bridge between Python and the Go-based Blueframe ecosystem. We developed two parallel wrapping strategies—CGo and Subprocess—to evaluate performance tradeoffs, using a custom "go-comtrade" repository as a complex test case for file I/O and multi-module integration. Our team overcame significant barriers regarding SEL-internal repository restrictions by identifying and documenting workarounds for missing Artifactory and Bitbucket dependencies. A major learning milestone involved mastering containerization tools like Docker, Kubernetes, and Helm to ensure the Field SDK could be deployed reliably in a simulated environment.

## Unfinished Work
We did not complete the full integration of the official Go SignalHub repository due to ongoing security and access negotiations with the SEL team. Additionally, the VSCode extension requires a structural split between production and development branches to resolve slow build times caused by bundling the large 2.5GB SDK artifact.

## Completed Issues/User Stories
Here are links to the issues that we completed in this sprint:

 * [Implementation of CGo and Subprocess Python-to-Go wrappers](https://github.com/calebh13/sel-blueframe-sdk/issues/19)
 * [Creation of Field SDK Troubleshooting and Environment Technology Guides](https://github.com/calebh13/sel-blueframe-sdk/issues/20)
 * [Initial Prototype of VSCode Field SDK Extension](https://github.com/calebh13/sel-blueframe-sdk/issues/21)
 * [Implementation of Python Wrapping Reading Functionality](https://github.com/calebh13/sel-blueframe-sdk/issues/22)
 * [Implementation of Python Wrapping Writing Functionality](https://github.com/calebh13/sel-blueframe-sdk/issues/23)

 ## Incomplete Issues/User Stories
 Here are links to issues we worked on but did not complete in this sprint:
 * https://github.com/calebh13/sel-blueframe-sdk/issues/7 -> Due to blockers with getting secure access to the Go SignalHub repository, we could not implement the publishing workflow
* https://github.com/calebh13/sel-blueframe-sdk/issues/6 -> For the same blockers which prevented the full publishing workflow, the publishing function  could not be completed this sprint.
* https://github.com/calebh13/sel-blueframe-sdk/issues/5 -> When we began sprint 2, we thought we would quickly have access to the existing codebase but did not anticipate all of the secure access issues. This prevented us from completeing the subscribe workflow and functions
* https://github.com/calebh13/sel-blueframe-sdk/issues/4 -> When we began sprint 2, we thought we would quickly have access to the existing codebase but did not anticipate all of the secure access issues. This prevented us from completeing the subscribe workflow and functions
 
 * [Access and wrap Go SignalHub repository] Access was blocked due to unresolved security concerns between the team and SEL.
 * [Removal of all internal SEL dependencies from SDK] Several private packages remain embedded that require permanent external replacements or packaging.
 
## Code Files for Review
Please review the following code files, which were actively developed during this sprint, for quality:
 * [VSCode 1-click Extension](https://github.com/calebh13/sel-blueframe-sdk/tree/main/code/extension)
 * [Python GO Wrapper](https://github.com/calebh13/sel-blueframe-sdk/tree/main/code/python-wrapper)
 * [COMTRADE Go](https://github.com/calebh13/sel-blueframe-sdk/blob/main/code/comtrade.go)
 * [COMTRADE Go Tests](https://github.com/calebh13/sel-blueframe-sdk/blob/main/code/comtrade_test.go)
 
## Retrospective Summary
Here's what went well:
  * Successfully built production-ready proofs-of-concept for cross-language integration.
  * Gained a deep architectural understanding of Blueframe and SignalHub systems through direct code interaction
  * Improved our ability to ask targeted technical questions and break down complex technological blockers.
 
Here's what we'd like to improve:
   * The speed of our development cycle when working with large 2.5GB SDK artifacts.
   * Reducing the time spent on technological onboarding for containerization tools like Kubernetes.
  
Here are changes we plan to implement in the next sprint:
   * Transition to externally accessible services like DockerHub for SDK dependencies to ensure out-of-the-box functionality.
   * Split the VSCode extension into production and development branches to improve build speed during the change-build-test cycle.
