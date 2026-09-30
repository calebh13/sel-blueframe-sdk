# Testing and Development Process Plan

**Sprint 4 Deliverable 3**  
**Owner:** Genevieve Kochel  
**Contributors:** Lucas Phillips (test harness and test suite)  
**Status:** Reviewed with the client. 2,155 tests passing on September 29, 2026. Joe looked at the test run offline and said it meets his preferences.

> The SignalHub client code is SEL Confidential. The test code lives in the team's private repository next to the prototype. This plan describes it without SEL source code.

## 1. Purpose

This plan defines the test categories for the Python wrapper, the tools used to run them, what counts as passing, and the workflow the team follows during Sprints 5 and 6.

## 2. Guiding Rule

The SEL Go client is treated as the source of truth. When a test disagrees with the client, either the test is wrong or it has found a client behavior the wrapper needs to design around. Bugs in the wrapper itself are kept as expected failures with a written reason until they are fixed.

## 3. Test Categories

| Category | What it checks | Tool | Passing means |
|---|---|---|---|
| Unit | Python API input checks, dtype mapping, rate and time helpers | pytest | Every case matches a reference model of the client's own math |
| Transport | Each transport (ctypes, cffi, bridge) on its own | pytest | Contract tests pass for each transport |
| Integration | Python to Go to Signal Hub and back, for publish, subscribe, history, and discovery | pytest with the fake Signal Hub | Values come back bit-exact with exact timestamps |
| Interoperability | All 16 sample encodings, with and without metadata | pytest, parametrized | Every encoding round-trips on every transport |
| Concurrency | Many threads and clients sharing publishers and subscriptions | pytest, Go race detector | No lost or duplicated samples, no deadlocks |
| Failure | Fatal Hub errors, dropped connections, Ctrl+C, fork, callbacks that raise | pytest with child processes and fault hooks in the fake Hub | The wrapper raises a Python error and stays usable where the design says it should |
| Large data | 1 million sample publishes and fetches | pytest, marked slow | Bit-exact results within the time limit |
| Cross-platform | Linux x86_64, Linux aarch64, Windows x86_64 | CI matrix | Same suite passes on each supported platform |
| Performance | Throughput, per-call cost, latency, memory | Benchmark scripts | Meets the targets in the [compatibility plan](02_Compatibility_Maintenance_Performance_Plan.md) |

## 4. Test Harness

The client talks to real SEL services, which are not available outside SEL. A fake Signal Hub was written in Go for testing. It uses the client's own wire code so the messages match what the real Hub expects. It stores history in memory, replays and streams data, supports time-aligned and control channels, and answers signal discovery queries. Test hooks let a test send a fatal error to every client, drop all connections, or make a signal never answer so timeouts can be tested.

One script builds the Go pieces and runs the full suite, then prints a pass and fail summary for each transport. Every test has a timeout, 60 seconds by default and up to 180 seconds for slow tests.

## 5. Current Results

The full suite was run on September 29, 2026.

| Measure | Result |
|---|---|
| Tests collected | 2,155 |
| Passed | 2,155 |
| Failed, errors, skipped | 0 |
| Run time | about 5 minutes |
| Transports covered | ctypes, cffi, bridge |

| Area | What it covers |
|---|---|
| Rate arithmetic | Every rate helper against a reference model over 17 rates |
| Publish and subscribe | Live and history round trips for 16 encodings and 12 rates, gap filling, backward time rules |
| Signals | Describing and parsing signals, bad input |
| Fetch and previous sample | Exact time ranges, empty ranges, unknown signals, timeouts |
| Metadata | Quality and time quality bits, trigger times, structured arrays |
| Timestamped (variable rate) | Live and fetch, equal and backward timestamps |
| Time-aligned | Bounded and live aligned reads, gaps, mixed rates |
| Control | Control signal round trips and invalid cases |
| Lifecycle | Context managers, double close, use after close |
| Callbacks | Delivery thread, exceptions, closing from inside a callback |
| Concurrency | 40 subscriptions and 40 publishers on one client, threads sharing a client |
| Reconnect | Resuming after the Hub drops the connection |
| Process | Fatal errors, Ctrl+C during a read, fork after connect |

Tests found four wrapper bugs during Sprint 4 and all four were fixed. One example is metadata fields being copied by position instead of by name.

## 6. Sample Tests

A round trip over every encoding, run once per transport.

```python
@pytest.mark.parametrize("enc", ALL_ENCODINGS)
def test_live_roundtrip_every_encoding(client, mp, enc):
    sig = signal_for(mp, "VA", enc, rate=60)
    t0 = base_time(client)
    sub = client.subscribe(sig, start=t0)
    vals = make_records(enc, 150)
    with client.publisher(sig, start=t0) as pub:
        assert pub.publish(t0, vals) == 150
    ts, got = collect(sub, 150)
    assert_records_equal(got, vals)
    np.testing.assert_array_equal(ts, expected_times(60, t0, 150))
    sub.close()
```

History replay after the data was already published.

```python
@pytest.mark.parametrize("enc", ALL_ENCODINGS)
def test_history_replay_every_encoding(client, testhub, mp, enc):
    sig = signal_for(mp, "VA", enc, rate=30)
    t0 = base_time(client)
    vals = make_records(enc, 90, seed=7)
    with client.publisher(sig, start=t0) as pub:
        pub.publish(t0, vals)
    assert testhub.wait_count(sig_id(client, sig), 90) == 90
    with client.subscribe(sig, start=t0) as sub:
        ts, got = collect(sub, 90)
    assert_records_equal(got, vals)
```

## 7. Acceptance Criteria for Sprint 5 and 6 Work

- All existing tests pass on every supported transport before a pull request is merged.
- New features come with tests first, following the test-driven approach the client suggested.
- Race detector runs clean on the wrapper's own Go code.
- Every workaround for a client behavior has a canary test that fails if the behavior changes.
- The suite passes on each platform in the compatibility matrix, or the gap is written down as a known limitation.

## 8. Development Workflow

1. Every task is a GitHub issue with an owner, a milestone, and story points.
2. Work happens on a feature branch named after the issue.
3. Pull requests link the issue they close and include a short before and after video.
4. At least one other team member reviews each pull request.
5. The Kanban board is updated as issues move from Backlog to Done.
6. Findings from client meetings are written up as MoMs in the sprint folder within a day.

## 9. Gaps

- Nothing has been tested against a real Signal Hub yet.
- Windows has not been tested.
- Soak tests over several hours have not been run.
- The fake Hub simplifies a few behaviors, such as data retention.
