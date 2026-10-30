# Deployment Plan and Status (Sprint 5)

## Where the SDK has to run

Blueframe is SEL's container based platform for operational technology. It runs on SEL-3350, SEL-3355 and SEL-3360 computers, or virtualized on other servers, and every application has to be signed and verified before it runs ([SEL Blueframe product page](https://selinc.com/products/blueframe), [SEL-3350 product page](https://selinc.com/products/3350/)). An engineer's Python script that uses the SDK would ship inside a Blueframe application container. Signal Hub runs on the same device, so the SDK connects to it on the device.

## Deployment path

1. Developer machine. The team's VS Code extension (`code/extension`, built in Sprint 3) extracts the Blueframe Field SDK, runs its SDK manager, and opens the Field SDK devcontainer. This is the environment SEL developers use to build Blueframe apps off site.
2. SDK inside the devcontainer. The engineer installs the Python SDK in that container. Today that means the bridge build from the private repo (see the [user guide](01_User_Guide.md)). After the move to C#, it will be pythonnet plus SEL's C# client assembly.
3. Blueframe app. The engineer's script and the SDK are packaged as a Blueframe application with the Field SDK tooling, signed, and installed on the device.
4. Device. The app connects to the device's Signal Hub.

## What was done in Sprint 5

| Step | Status |
|---|---|
| CI that runs real tests on every push, Linux and Windows | Done. [Tests workflow](../../.github/workflows/tests.yml), [run 41144016907](https://github.com/calebh13/sel-blueframe-sdk/actions/runs/41144016907) |
| SDK runs end to end against the fake Signal Hub | Done. 2,161 tests and the pub/sub example, October 30 |
| Demo app for client walkthroughs | Done, local only ([issue #45](https://github.com/calebh13/sel-blueframe-sdk/issues/45)) |
| SDK installed inside the Field SDK devcontainer | Not done |
| Packaged and signed as a Blueframe app | Not done |
| Run on a Blueframe device against a real Signal Hub | Not done |

The last three steps need SEL's signing process and either a device or a virtual Blueframe instance. The team does not have either yet. These are the first things to ask Joe for in Sprint 6.

## Open questions for SEL

* Can a Blueframe app start a child process? The Go bridge needs this. The C# path does not.
* Is the .NET runtime available inside a Blueframe app container, or does the app have to bring its own?
* Which Python version is in the Field SDK base image?
