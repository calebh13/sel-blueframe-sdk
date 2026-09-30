# Documentation and Extensibility Plan

**Sprint 4 Deliverable 4**  
**Owner:** Darron Li  
**Status:** Reviewed with the client on September 28, 2026. Structure and conventions are defined, and pages will be written in Sprints 5 and 6.

## 1. Purpose

The client asked for documentation to be a major part of the project and wants future developers to be able to extend the wrapper without the original team. This plan defines who the docs are for, how they are organized, the conventions they follow, and how a new function gets added to the wrapper.

## 2. Audiences

| Audience | What they need |
|---|---|
| Engineers using the SDK | Install steps, a working example in a few minutes, how to publish and subscribe, what errors mean |
| Developers maintaining the wrapper | How the layers fit together, how to build and test, how to add a function, how to handle a client update |
| SEL reviewers | Supported platforms, known limitations, test results |

Engineers using the SDK should never need to read Go or know that a bridge process exists.

## 3. Documentation Structure

```
docs/
  index.md                  Overview and a 10 line example
  getting-started/
    installation.md         Supported Python versions and platforms, pip install
    quickstart.md           Connect, publish, fetch, subscribe
  user-guide/
    signals.md              Signals, rates, dtypes, metadata
    publishing.md           Fixed-rate, timestamped, and control publishers
    subscribing.md          Single, multi, time-aligned, callbacks vs pulling
    history.md              fetch and previous_sample
    discovery.md            Finding signals
    time.md                 Nanosecond timestamps, datetime and numpy helpers
    threads-and-processes.md  Threading rules and the fork rule
    errors.md               Every exception and what to do about it
    troubleshooting.md      Name already in use, dropped samples, bridge exits
  reference/                Generated from docstrings and type stubs
  maintainers/
    architecture.md         The five layers and the design rules
    building.md             Toolchain, build script, running tests
    adding-a-function.md    Step by step, see Section 6
    client-updates.md       What to do when the SEL client changes
    releasing.md            Versioning, wheels, changelog
  limitations.md            Known gaps and untested platforms
  changelog.md
```

The README stays short and links into this tree.

## 4. Tooling

| Need | Proposed tool | Reason |
|---|---|---|
| Site generator | MkDocs with the Material theme | Markdown source, simple to host inside SEL |
| API reference | mkdocstrings | Builds reference pages from Python docstrings |
| Docstring style | NumPy style | Familiar to engineers who already use numpy and scipy |
| Type information | Type hints plus `.pyi` stubs | Editors show signatures and the reference pages stay accurate |
| Example checking | Examples run as tests in CI | Examples in the docs cannot go stale |

Sphinx is the fallback if SEL already standardizes on it.

## 5. Conventions

- Every public function has a docstring with a one line summary, parameters, return value, exceptions, and a short example.
- Names in Python follow Python style (`snake_case`) and are chosen per operation. They do not copy Go names.
- Timestamps are always int64 nanoseconds in the API. The docs call this out wherever a time appears.
- Units and dtypes are stated for every value.
- Every example in the docs is a runnable script that is also run by the test suite.
- A limitation is documented in `limitations.md` in the same pull request that finds it.
- Docs changes ship in the same pull request as the code change they describe.
- The changelog lists each release with added, changed, deprecated, and removed items.

## 6. Extending the Wrapper

### Adding a function today

1. Add a non-generic function to the Go facade that calls the client and recovers panics.
2. Add the RPC to the bridge `.proto` file and regenerate the Go and Python stubs.
3. Add the handler on the bridge side.
4. Add the Python method with a docstring.
5. Add tests that run on every transport.
6. Add the function to the reference docs and, if it is common, the user guide.

This is about 40 to 60 lines, mostly mechanical.

### Adding a function after the planned generator

1. Add one entry to the list of exposed functions (client symbol, Python name, adapter type).
2. Regenerate and rebuild.
3. Add tests.

The generator stops with a clear error if a listed function is missing, changes signature, or has a shape it does not support.

### When the SEL client changes

- A changed signature fails the Go build in the facade.
- A committed snapshot of the exposed API shows the diff in CI so changes are reviewed.
- Canary tests fail if a client behavior the wrapper works around is fixed, so the workaround can be removed.
- `client-updates.md` walks through these steps.

### Deprecation

Renamed or removed Python functions stay for one minor release with a `DeprecationWarning` and a changelog entry.

## 7. Open Questions

- Where will SEL host the docs internally?
- Does SEL have a required docs tool or style guide?
- Now that the SDK is moving to the C# client, can its XML doc comments feed the Python reference pages?
