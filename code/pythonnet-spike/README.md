# Python.NET interop spike

This spike checks how Python.NET behaves when Python calls a C# hub client. SEL's C# client is confidential and is not in this repo, so `csharp/LocalHub.cs` is a small in-memory stand-in with the same call shapes the SDK needs. It has a generic publish, a history fetch, callbacks sent from a background thread, and methods that throw.

The numbers below measure Python.NET itself. They say nothing about SEL's client or a real Signal Hub.

## Run it

Needs the .NET 10 SDK and Python 3.12.

```bash
cd code/pythonnet-spike
python3 -m venv .venv
.venv/bin/pip install pythonnet numpy
dotnet build csharp -c Release
.venv/bin/python spike.py
```

## Results

Ubuntu 24.04, AMD Ryzen AI 7 350, Python 3.12.3, pythonnet 3.2.1, .NET 10.0.14. Four runs on October 30, 2026.

| Check | Result |
|---|---|
| Start the .NET runtime and load the assembly | 0.15 to 0.17 s |
| Plain method call from Python | 2.7 to 2.9 us per call |
| Publish 1M float64 samples, copying element by element | about 4.6 million samples/s |
| Publish 1M float64 samples, one block copy with `Marshal.Copy` | 20 to 135 million samples/s (120 million or more in 3 of 4 runs) |
| Generic `Publish[Single]` called from Python | works, 1,000,000 samples stored |
| 100,000 callbacks from a C# thread into Python | all delivered, 316k to 355k callbacks/s |
| Missing channel | raises `KeyNotFoundException` in Python, session keeps running |
| Mismatched array lengths | raises `ArgumentException` in Python, session keeps running |
| `Environment.FailFast` inside the C# code | the Python process dies with exit code -6 |

## What it means for the SDK

Generic methods can be called straight from Python with `Publish[Single]`. The Go prototype needed a hand-written shim with 16 type cases for the same thing.

Callbacks from a C# thread reach Python without extra glue. Python.NET takes the GIL for each call. At about 330k callbacks per second, per-sample callbacks are fine for slow signals, but fast signals should be handed to Python in batches.

Arrays have to be copied as one block. Copying element by element is about 25 times slower.

Normal C# exceptions turn into Python exceptions. A fatal error in the .NET runtime still kills the Python session, which is the same problem the ctypes and cffi transports had. For that reason the Go subprocess bridge stays as the fallback until the C# client is tested against a real Signal Hub.
