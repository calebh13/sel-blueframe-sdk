# Sprint 2 Report (Dates from Sprint 2/22 to Sprint 4/9)

Sprint Video: https://youtu.be/2cOXRF26fVw

## What's New (User Facing)
* Go-based COMTRADE reader for parsing .cfg and .dat files
* Go-based COMTRADE writer for generating valid COMTRADE files
* COMTRADE testing suite for validating correctness and reliability
* Proof-of-concept integration demonstrating calling Go code from Python libraries
* Blueframe development environment setup using Kubernetes, Docker, and Tilt

## Work Summary (Developer Facing)
This sprint focused on both infrastructure setup and development of a standalone COMTRADE processing toolchain. After receiving access to the field SDK, the team successfully configured a local development environment using Kubernetes, Docker, and Tilt to run Blueframe containers. In parallel, the team designed and implemented a COMTRADE reader, writer, and testing suite in Go, which serves both as a production-ready utility and as a proof of concept for cross-language integration with Python. Despite significant external blockers—primarily lack of access to the SignalHub Go repository—the team adapted by redirecting efforts toward independent development work. This required learning and working in unfamiliar languages and ecosystems, demonstrating strong adaptability and technical growth.

## Unfinished Work
The team was unable to complete planned functionality related to SignalHub, including subscribing to signals, publishing signals, and discovering available signals. These features were blocked due to lack of access to the existing Go SignalHub repository, preventing any meaningful development or integration. Progress on these issues has been documented, and they will be carried forward into the next sprint once access is granted. Each of the issues relating to these features are recorded below.

## Completed Issues/User Stories
Here are links to the issues that we completed in this sprint:
* https://github.com/calebh13/sel-blueframe-sdk/issues/14 -> COMTRADE Functionality
* https://github.com/calebh13/sel-blueframe-sdk/issues/13 -> COMTRADE Read / Write
* https://github.com/calebh13/sel-blueframe-sdk/issues/12 -> Write COMTRADE
* https://github.com/calebh13/sel-blueframe-sdk/issues/11 -> Read COMTRADE

## Incomplete Issues/User Stories
Here are links to issues we worked on but did not complete in this sprint:
* https://github.com/calebh13/sel-blueframe-sdk/issues/7 -> Due to blockers with getting secure access to the Go SignalHub repository, we could not implement the publishing workflow
* https://github.com/calebh13/sel-blueframe-sdk/issues/6 -> For the same blockers which prevented the full publishing workflow, the publishing function  could not be completed this sprint.
* https://github.com/calebh13/sel-blueframe-sdk/issues/5 -> When we began sprint 2, we thought we would quickly have access to the existing codebase but did not anticipate all of the secure access issues. This prevented us from completeing the subscribe workflow and functions
* https://github.com/calebh13/sel-blueframe-sdk/issues/4 -> When we began sprint 2, we thought we would quickly have access to the existing codebase but did not anticipate all of the secure access issues. This prevented us from completeing the subscribe workflow and functions

## Code Files for Review
Please review the following code files, which were actively developed during this sprint, for quality:
* [COMTRADE Reader (Go)] https://github.com/calebh13/sel-blueframe-sdk/blob/main/code/comtrade.go
* [COMTRADE Writer (Go)] https://github.com/calebh13/sel-blueframe-sdk/blob/main/code/comtrade.go
* [COMTRADE Test Suite (Go)] https://github.com/calebh13/sel-blueframe-sdk/blob/main/code/comtrade_test.go

## Retrospective Summary

Here's what went well:
* Strong team coordination and productivity once blockers were removed  
  Once access constraints were lifted, the team was able to quickly align on priorities and execute effectively. Work was distributed efficiently, and team members were proactive in picking up tasks, which allowed us to make meaningful progress in a short amount of time.

* Effective communication with stakeholders, including client, SEL mentors, and capstone professor  
  Despite ongoing challenges, we maintained consistent and transparent communication with all stakeholders. We made a deliberate effort to keep everyone informed of blockers, progress, and risks, which helped manage expectations and ensured alignment throughout the sprint.

* Successful adaptation to unfamiliar languages (Go) and technologies  
  The team demonstrated strong adaptability by working in Go and with infrastructure tools like Kubernetes, Docker, and Tilt—many of which were new to us. We were able to ramp up quickly and still produce functional, well-structured code despite the learning curve.

* Delivery of a meaningful and functional COMTRADE toolchain despite constraints  
  Even with major external blockers, we delivered a working COMTRADE reader, writer, and testing suite. This not only adds value as a standalone tool but also serves as a proof of concept for integrating Go-based services with the Python SDK.

Here's what we'd like to improve:
* Establishing more direct and efficient communication channels within SEL  
  Our current communication pathways within SEL proved to be slow and somewhat opaque, which significantly impacted our ability to move forward. Improving access to the right contacts and reducing dependency on indirect communication will be critical going forward.

* Improving meeting scheduling and consistency in preparation for future sprints  
  While meeting structure was less critical during this blocked sprint, it will become increasingly important as development ramps up. Establishing a more consistent schedule and clearer agendas will help ensure alignment and maximize productivity.

* Better communication and presentation of results to highlight the value of completed work  
  Although we made substantial technical progress, we could have done a better job presenting and “selling” our accomplishments. Improving how we communicate outcomes—especially to non-technical stakeholders—will help ensure our work is fully understood and appreciated.

Here are changes we plan to implement in the next sprint:
* Establish direct contacts with SEL security team to improve access turnaround and visibility  
  We are actively working to build more direct relationships with the SEL security team. This should give us better insight into access requirements and timelines, and help reduce delays caused by unclear or indirect communication channels.

* Implement a consistent meeting schedule and standardized update format  
  We plan to introduce a regular meeting cadence along with a structured update format. This will help keep both internal team members and external stakeholders aligned, while also improving accountability and clarity.

* Improve documentation with visuals, diagrams, and stronger demos to better communicate progress and impact  
  To better showcase our work, we will focus on enhancing documentation with diagrams, visual workflows, and more polished demos. This will make our progress easier to understand and help communicate the value of our contributions more effectively.