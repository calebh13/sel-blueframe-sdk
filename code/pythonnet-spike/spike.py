import os
import subprocess
import sys
import threading
import time

import numpy as np
import pythonnet

HERE: str = os.path.dirname(os.path.abspath(__file__))
DLL_DIR: str = os.path.join(HERE, "csharp", "bin", "Release", "net10.0")


def load_hub_assembly() -> float:
    start: float = time.perf_counter()
    pythonnet.load("coreclr")
    import clr
    sys.path.append(DLL_DIR)
    clr.AddReference("StandInHub")
    return time.perf_counter() - start


def time_plain_calls(hub, call_count: int) -> float:
    start: float = time.perf_counter()
    for call_index in range(call_count):
        hub.Ping(call_index)
    elapsed: float = time.perf_counter() - start
    return elapsed / call_count * 1e6


def to_dotnet_doubles(values: np.ndarray):
    from System import Array, Double, IntPtr
    from System.Runtime.InteropServices import Marshal
    contiguous: np.ndarray = np.ascontiguousarray(values, dtype=np.float64)
    copied = Array.CreateInstance(Double, len(contiguous))
    Marshal.Copy(IntPtr(contiguous.ctypes.data), copied, 0, len(contiguous))
    return copied


def to_dotnet_longs(values: np.ndarray):
    from System import Array, Int64, IntPtr
    from System.Runtime.InteropServices import Marshal
    contiguous: np.ndarray = np.ascontiguousarray(values, dtype=np.int64)
    copied = Array.CreateInstance(Int64, len(contiguous))
    Marshal.Copy(IntPtr(contiguous.ctypes.data), copied, 0, len(contiguous))
    return copied


def time_publish(hub, sample_count: int) -> dict[str, float]:
    from System import Array, Double, Single
    timestamps: np.ndarray = np.arange(sample_count, dtype=np.int64)
    values: np.ndarray = np.sin(np.arange(sample_count) / 60.0)
    results: dict[str, float] = {}

    start: float = time.perf_counter()
    slow_values = Array[Double](values.tolist())
    hub.Publish[Double]("slow", to_dotnet_longs(timestamps), slow_values)
    results["element_copy_samples_per_s"] = sample_count / (time.perf_counter() - start)

    start = time.perf_counter()
    hub.Publish[Double]("fast", to_dotnet_longs(timestamps), to_dotnet_doubles(values))
    results["block_copy_samples_per_s"] = sample_count / (time.perf_counter() - start)

    float_values = Array[Single](values.astype(np.float32).tolist())
    stored_count: int = hub.Publish[Single]("float32", to_dotnet_longs(timestamps), float_values)
    results["generic_float32_stored"] = float(stored_count)

    fetched = hub.Fetch("fast")
    results["fetch_matches"] = float(len(fetched) == sample_count and abs(fetched[10] - values[10]) < 1e-12)
    return results


def time_callbacks(hub, sample_count: int) -> dict[str, float]:
    from System import Action, Int64, Double
    received_count: list[int] = [0]
    callback_threads: set[int] = set()
    all_received = threading.Event()

    def on_sample(timestamp: int, value: float) -> None:
        received_count[0] = received_count[0] + 1
        callback_threads.add(threading.get_ident())
        if received_count[0] == sample_count:
            all_received.set()

    hub.Subscribe("live", Action[Int64, Double](on_sample))
    timestamps: np.ndarray = np.arange(sample_count, dtype=np.int64)
    values: np.ndarray = np.ones(sample_count)
    start: float = time.perf_counter()
    hub.Publish[Double]("live", to_dotnet_longs(timestamps), to_dotnet_doubles(values))
    finished: bool = all_received.wait(timeout=30)
    elapsed: float = time.perf_counter() - start

    results: dict[str, float] = {}
    results["delivered"] = float(received_count[0])
    results["finished"] = float(finished)
    results["callbacks_per_s"] = received_count[0] / elapsed
    results["ran_off_main_thread"] = float(threading.main_thread().ident not in callback_threads)
    return results


def check_exceptions(hub) -> dict[str, str]:
    results: dict[str, str] = {}
    try:
        hub.Fetch("missing")
    except Exception as error:
        results["missing_channel"] = type(error).__name__ + ": " + str(error).splitlines()[0]
    from System import Array, Double, Int64
    try:
        hub.Publish[Double]("bad", Array[Int64]([1, 2]), Array[Double]([1.0]))
    except Exception as error:
        results["length_mismatch"] = type(error).__name__ + ": " + str(error).splitlines()[0]
    results["session_alive_after"] = str(hub.Ping(41) == 42)
    return results


def check_crash_isolation() -> int:
    crash_code: str = (
        "import sys; sys.path.insert(0, %r); import spike; spike.load_hub_assembly();"
        "from StandInHub import LocalHub; LocalHub().CrashProcess()" % HERE
    )
    finished = subprocess.run([sys.executable, "-c", crash_code], capture_output=True)
    return finished.returncode


def main() -> None:
    load_seconds: float = load_hub_assembly()
    from StandInHub import LocalHub
    hub = LocalHub()
    print("python", sys.version.split()[0], "pythonnet runtime coreclr")
    print("load_and_add_reference_s", round(load_seconds, 3))
    print("plain_call_us", round(time_plain_calls(hub, 100_000), 2))
    publish_results: dict[str, float] = time_publish(hub, 1_000_000)
    for name in publish_results:
        print(name, round(publish_results[name], 1))
    callback_results: dict[str, float] = time_callbacks(hub, 100_000)
    for name in callback_results:
        print(name, round(callback_results[name], 1))
    exception_results: dict[str, str] = check_exceptions(hub)
    for name in exception_results:
        print(name, exception_results[name])
    print("crash_in_child_exit_code", check_crash_isolation())


if __name__ == "__main__":
    main()
