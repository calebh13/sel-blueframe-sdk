# Sprint x Report (Dates from Sprint 4/10 to Sprint 5/1)

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

 * [Implementation of CGo and Subprocess Python-to-Go wrappers]
 * [Creation of Field SDK Troubleshooting and Environment Technology Guides]
 * [Initial Prototype of VSCode Field SDK Extension]

 ## Incomplete Issues/User Stories
 Here are links to issues we worked on but did not complete in this sprint:
 
 * [Access and wrap Go SignalHub repository] Access was blocked due to unresolved security concerns between the team and SEL.
 * [Removal of all internal SEL dependencies from SDK] Several private packages remain embedded that require permanent external replacements or packaging.
 
## Code Files for Review
Please review the following code files, which were actively developed during this sprint, for quality:
 
 
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
