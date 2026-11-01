# Python SDK User Guide (Sprint 5 draft)

This guide is for an engineer who wants to publish or subscribe to Signal Hub signals from Python. It covers the current SDK, which runs SEL's Go SignalHub client in a separate bridge process. The move to the C# client is still in progress (see [issue #36](https://github.com/calebh13/sel-blueframe-sdk/issues/36)), and the Python API in this guide is meant to stay the same after that move.

The SDK source is in the team's private `gosignals` repository because the client is SEL Confidential. Only the API and examples are in this public repo.

## 1. Install

You need the following.

* Python 3.12 (3.10 and 3.13 are planned, see the [compatibility matrix](02_Compatibility_Matrix.md))
* Go 1.25 or newer, to build the bridge and the fake Signal Hub
* Read access to the private `gosignals` repository

Clone both repositories into the same folder. The example in section 3 is run from this repo with a relative path, so it fails if `sel-blueframe-sdk` is not next to `gosignals`.

```bash
git clone https://github.com/calebh13/gosignals.git
git clone https://github.com/calebh13/sel-blueframe-sdk.git
cd gosignals/pysdk
python3 -m venv .venv
.venv/bin/pip install numpy grpcio protobuf
./build.sh bridge testhub
export PYTHONPATH=$PWD/python
```

`build.sh` writes `build/gosignals-bridge` and `build/testhub`. The SDK finds the bridge there on its own.

## 2. Start a hub to test against

There is no real Signal Hub outside SEL, so the team's fake hub is used for development. It speaks the real wire protocol and prints its ports when it starts.

```bash
./build/testhub
# PORTS hub=33631 graph=45251 nats=45069 ctl=35739
```

Leave it running in its own terminal. The port numbers change each time.

## 3. Publish and subscribe

[`examples/pubsub_example.py`](examples/pubsub_example.py) connects one publisher client and one subscriber client. The subscriber gets each batch through a callback, and then the same data is fetched back as history.

```bash
.venv/bin/python ../../sel-blueframe-sdk/docs/Sprint5/examples/pubsub_example.py 33631
```

Output from a run on October 30, 2026, against the fake hub.

```
published 600 samples
subscriber received 600 samples, finished: True
fetch returned 600 samples, first three [0.         0.10452846 0.20791169]
```

The core of it is below.

```python
import gosignals as gs

voltage = gs.Signal("SUB1.BUS1", "VA", rate=60, dtype="float32")
with gs.connect("127.0.0.1", hub_port, name="example-pub", transport="bridge") as client:
    start_time = client.rates.floor(1, client.now())
    with client.publisher(voltage, start=start_time) as publisher:
        publisher.publish(start_time, values)
    timestamps, values = client.fetch(voltage, start=start_time, end=start_time + 10_000_000_000)
```

Timestamps are int64 nanoseconds since the Unix epoch. Values go in and come out as numpy arrays.

## 4. Public API

Everything below is imported from `gosignals`. Nothing in the API mentions Go.

| Name | What it does |
|---|---|
| `connect(host, port, name, transport="bridge", ...)` | Opens a `Client`. Options include `nats`, `graph`, `publish_signals`, `auto_restart` and `unique_name`. |
| `Signal(measurement_point, quantity, rate, dtype)` | Names a signal. `rate` is samples per second, or `VARIABLE_RATE`. |
| `Client.publisher(signal, start)` | Fixed rate publisher. `publish(t0, values)` sends a block of samples. |
| `Client.timestamped_publisher(signal)` | Publisher for samples that each carry their own timestamp. |
| `Client.control_publisher(signal)` | Publisher for control values. |
| `Client.subscribe(signal, start, end=LIVE, on_data=None)` | One signal. Read with `read()` and `read_all()`, or pass `on_data(timestamps, values)`. |
| `Client.subscribe_many(signals, start, end=LIVE)` | Several signals on one subscription. |
| `Client.subscribe_aligned(signals, start, end=LIVE)` | Several signals lined up on one time grid, one row per timestamp. |
| `Client.fetch(signal, start, end)` | History for `[start, end)` as `(timestamps, values)`. |
| `Client.previous_sample(signal, lookback=60)` | Newest sample in the lookback window, or `None`. |
| `Client.monitor_signals(query)` | Watches the signal directory (Graph) for signals that match. |
| `Client.rates` | Helpers for rate and time math, for example `floor`, `time_at_index`, `samples_in_period`. |

Errors are normal Python exceptions.

| Exception | When |
|---|---|
| `GoSignalsError` | The hub or client rejected the call, for example a signal with no measurement point. |
| `TransportError` | The bridge process died and `auto_restart` is off. |
| `BridgeRestartedError` | Raised once by a publisher after the bridge was restarted. Samples sent right before may be lost. |
| `ClientNameInUseError` | Another open client already uses that name. Use `unique_name=True` in notebooks and scripts that run more than once. |
| `CallbackError` | An `on_data` callback raised. The original exception is its `__cause__`, and the subscription lists them in `callback_errors`. |

## 5. Things to know

* Use `transport="bridge"`. The `ctypes` and `cffi` transports load Go into Python, and a Go crash or `os.Exit` takes the Python session down with it.
* With `auto_restart=True`, a dead bridge is restarted and subscriptions continue from the last timestamp received, with no gaps or duplicates in the tests.
* A process that forks after connecting gets a clear error in the child. Make a new client in the child instead.
* Each client starts its own bridge process (about 17 MB).
* The SDK has only been tested against the fake hub so far, not a real Signal Hub.
