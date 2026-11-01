# Deployment (Sprint 5)

## Platform

The deployment is on premise. Blueframe is SEL's container based platform for operational technology. It runs on SEL-3350, SEL-3355 and SEL-3360 computers, or virtualized on other servers, and every application has to be signed and verified before it runs ([SEL Blueframe product page](https://selinc.com/products/blueframe), [SEL-3350 product page](https://selinc.com/products/3350/)). An engineer's Python script that uses the SDK ships inside a Blueframe application container. Signal Hub runs on the same device, so the SDK connects to it on the device.

## Deployment path

1. Developer machine. The team's VS Code extension (`code/extension`, built in Sprint 3) extracts the Blueframe Field SDK, runs its SDK manager, and opens the Field SDK devcontainer in one step. This is the environment SEL developers use to build Blueframe apps off site.
2. SDK with the SignalHub client. The Python SDK, its Go bridge and SEL's unmodified SignalHub client are built with `build.sh` and installed in a Python venv (see the [user guide](01_User_Guide.md)). After the move to C#, this step installs pythonnet and SEL's C# client assembly instead of the Go bridge.
3. Signal Hub server. The SDK connects to a Signal Hub server. On the developer machine this is the server in `gosignals/pysdk/testhub`, which is built from SEL's unmodified gosignals packages (ihub, signal and graph) with an embedded NATS server, so it speaks the real wire protocols. On a device it is the device's own Signal Hub.
4. Blueframe app. The engineer's script and the SDK are packaged as a Blueframe application with the Field SDK tooling, signed, and installed on the device.

## What was deployed in Sprint 5

| Step | Status |
|---|---|
| VS Code extension sets up the Field SDK devcontainer | Deployed. [code/extension](../../code/extension) |
| SDK and SEL's SignalHub client against a Signal Hub server running locally | Deployed and tested. 2,161 tests and the pub/sub example passed on October 30 ([evidence](evidence/)) |
| Demo app for client walkthroughs | Deployed against the same Signal Hub server ([issue #45](https://github.com/calebh13/sel-blueframe-sdk/issues/45)) |
| Team CI pipeline on every push and pull request | Running. [ci.yml](../../.github/workflows/ci.yml) |
| Bridge and Signal Hub server builds for Linux aarch64 and Windows x86_64 | Bridge builds for both, see the [compatibility matrix](02_Compatibility_Matrix.md) |
| Packaged and signed as a Blueframe app | Next step. It needs SEL's signing process |

## Deployment steps

1. Install the extension `.vsix` in VS Code and run **BlueFrame SDK: Launch** from the Command Palette. It extracts the Field SDK and opens the devcontainer.
2. Clone `gosignals` and `sel-blueframe-sdk` into the same folder, then build and install the SDK with the commands in section 1 of the [user guide](01_User_Guide.md).
3. Start the Signal Hub server with `./build/testhub` in `gosignals/pysdk`. It prints its ports when it starts.
4. Connect from Python with `gs.connect("127.0.0.1", hub_port, name=..., transport="bridge")`. The [pub/sub example](examples/pubsub_example.py) does this end to end.
5. Push to GitHub. The team CI pipeline builds and tests the Go code on every push and pull request.

## Open questions for SEL

* Can a Blueframe app start a child process? The Go bridge needs this. The C# path does not.
* Is the .NET runtime available inside a Blueframe app container, or does the app have to bring its own?
* Which Python version is in the Field SDK base image?
