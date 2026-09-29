# SEL Blueframe Python SDK

## Project summary

### One-sentence description of the project

A Python SDK that wraps SEL's existing SignalHub client so engineers can publish, subscribe to, and look up Blueframe signals from Python without learning Go or the Signal Hub protocol.

### Additional information about the project

This is the WSU CptS 421-423 capstone project for Spring to Fall 2026, sponsored by Schweitzer Engineering Laboratories (SEL).

Blueframe is SEL's hardened, Linux-based operating system for automation and monitoring in substations and other operational technology sites. Applications on Blueframe move real-time data through Signal Hub, a service with a low-level binary TCP protocol. SEL already has Signal Hub clients in Go and C#, but the automation and electrical engineers who would use them mostly work in Python.

The SDK does not port the protocol to Python. A third copy of the client would drift from the other two and double the maintenance work. Instead, the Python SDK wraps an existing client and hides the Go or C# details behind a small Python API that takes and returns numpy arrays and raises normal Python exceptions.

Sprint 4 tested three ways of wrapping the Go client (ctypes, cffi, and a gRPC subprocess bridge). All three worked, but only the bridge keeps a crash in the Go code from taking down the engineer's Python session. The research also found that most of the effort with Go goes into glue code, so the C# client with Python.NET is being evaluated at the start of Sprint 5 before the bindings are built. The details are in the [Sprint 4 planning documents](docs/Sprint4/).

## Installation

### Prerequisites

SEL Blueframe Field SDK is required to develop for Blueframe off-site. The Vanguard SDK is required to develop on-site.

The Python SDK prototype also needs the following.

* Python 3.12
* Go 1.25 or newer
* gcc, only for the ctypes and cffi transports
* Access to the team's private SignalHub client repository, since the client is SEL Confidential

The COMTRADE code in `code/` needs Go 1.21 or newer. The VS Code extension in `code/extension` needs Node.js, VS Code, Docker, and the Dev Containers extension.

### Add-ons

* **numpy** holds sample values and timestamps on the Python side.
* **grpcio** and **protobuf** carry calls between Python and the Go bridge process.
* **cffi** loads the Go shared library for the cffi transport.
* **pytest** and **pytest-timeout** run the test suite and stop tests that hang.
* **@vscode/vsce** packages the VS Code extension.

### Installation Steps

Confidential. The Python SDK prototype and its build steps are kept in SEL's private repository.

## Functionality

**Python SDK prototype.** An engineer's script connects to Signal Hub, publishes a signal, and reads it back without touching Go.

```python
import numpy as np
import gosignals as gs

with gs.connect("127.0.0.1", 49000, name="relay-study", transport="bridge") as hub:
    va = gs.Signal("SUB1.BUS1", "VA", rate=60, dtype="float32")
    t0 = hub.rates.floor(1, hub.now()) - 20_000_000_000

    with hub.publisher(va, start=t0) as pub:
        pub.publish(t0, np.sin(np.arange(600) / 60))

    ts, vals = hub.fetch(va, start=t0, end=t0 + 2_000_000_000)
    last = hub.previous_sample(va, lookback=30)
```

Timestamps are int64 nanoseconds. Errors from the Go client come back as `gs.GoSignalsError`.

**COMTRADE reader and writer.** `code/comtrade.go` reads and writes IEEE COMTRADE `.cfg` and `.dat` files. The tests write a sample file, read it back, and compare the two.

```bash
cd code
go mod init comtrade
go test ./...
```

**VS Code extension.** `code/extension` extracts the Blueframe Field SDK, runs the SDK manager, and opens the devcontainer. Put `field-sdk.tgz` in `code/extension`, then package and install it.

```bash
cd code/extension
npm install
npm run package
```

Install the `.vsix` from the Extensions menu and run **BlueFrame SDK: Launch** from the Command Palette.

## Known Problems

* The Python SDK prototype has only been tested against a fake Signal Hub. It has not been run against a real Signal Hub yet.
* Windows has not been tested.
* The ctypes and cffi transports crash the Python process if the Go client exits or panics, and they hang in fork-based worker pools. ctypes is still the default, so pass `transport="bridge"` to `gs.connect`.
* `code/comtrade.go` has no `go.mod`, so `go test` fails until `go mod init` is run in `code/`.
* Building the VS Code extension is slow because it bundles the 2.5 GB Field SDK archive.

## Contributing

This repo will eventually be SEL confidential so it should not be open to public contribution.

## Additional Documentation

  * [Sprint 1 Report](https://github.com/calebh13/sel-blueframe-sdk/blob/main/Sprints/Sprint_1/Sprint_Report.md)
  * [Sprint 2 Report](https://github.com/calebh13/sel-blueframe-sdk/blob/main/Sprints/Sprint2/Sprint_Report.md)
  * [Sprint 3 Report](https://github.com/calebh13/sel-blueframe-sdk/blob/main/Sprints/Sprint_3/Sprint_Report.md)
  * [Sprint 4 Report](https://github.com/calebh13/sel-blueframe-sdk/blob/main/Sprints/Sprint_4/Sprint_Report.md)
  * [Sprint 4 Python Binding Requirements and Architecture Plan](docs/Sprint4/01_Python_Binding_Architecture_Plan.md)
  * [Sprint 4 Compatibility, Maintenance, and Performance Plan](docs/Sprint4/02_Compatibility_Maintenance_Performance_Plan.md)
  * [Sprint 4 Testing and Development Process Plan](docs/Sprint4/03_Testing_and_Development_Process_Plan.md)
  * [Sprint 4 Documentation and Extensibility Plan](docs/Sprint4/04_Documentation_and_Extensibility_Plan.md)
  * [Sprint 4 Minutes of Meetings](Sprints/Sprint_4/MoM/)
  * [Project Description](Reports/01_Project_Description.pdf) and [Requirements and Specifications](Reports/02_Requirements_and_Specifications.pdf)
  * Sprint videos for [Sprint 1](https://youtu.be/a1FaINWTAow), [Sprint 2](https://youtu.be/2cOXRF26fVw), and [Sprint 3](https://youtu.be/-EzL6DnFo0w)
  * [Python wrapper demo video](https://youtu.be/ylGSQu6aIVY)
  * [Google Drive with Additional Documentation](https://drive.google.com/drive/u/1/folders/1l-EKSdQvaO1z8Cdv-wOVbb2EH3CFL6S-)

## License

No open-source license is granted. The project is being built for SEL, and the code will move into SEL's confidential repositories, so a `LICENSE.txt` is not included.
