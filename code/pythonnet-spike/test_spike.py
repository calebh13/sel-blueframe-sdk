import sys

import numpy as np
import pytest

import spike


@pytest.fixture(scope="module")
def hub():
    spike.load_hub_assembly()
    from StandInHub import LocalHub
    return LocalHub()


def test_plain_call_returns_value(hub) -> None:
    assert hub.Ping(41) == 42


def test_publish_and_fetch_round_trip(hub) -> None:
    from System import Double
    sample_count: int = 1000
    timestamps: np.ndarray = np.arange(sample_count, dtype=np.int64)
    values: np.ndarray = np.linspace(-1.0, 1.0, sample_count)
    stored_count: int = hub.Publish[Double]("round_trip", spike.to_dotnet_longs(timestamps), spike.to_dotnet_doubles(values))
    fetched = hub.Fetch("round_trip")
    assert stored_count == sample_count
    for sample_index in range(sample_count):
        assert fetched[sample_index] == values[sample_index]


def test_generic_float32_publish(hub) -> None:
    from System import Array, Single
    float_values = Array[Single]([1.5, 2.5, 3.5])
    stored_count: int = hub.Publish[Single]("float32_case", spike.to_dotnet_longs(np.arange(3)), float_values)
    assert stored_count == 3
    assert list(hub.Fetch("float32_case")) == [1.5, 2.5, 3.5]


def test_callbacks_from_dotnet_thread(hub) -> None:
    callback_results: dict[str, float] = spike.time_callbacks(hub, 10_000)
    assert callback_results["finished"] == 1.0
    assert callback_results["delivered"] == 10_000.0
    assert callback_results["ran_off_main_thread"] == 1.0


def test_dotnet_exceptions_become_python_exceptions(hub) -> None:
    exception_results: dict[str, str] = spike.check_exceptions(hub)
    assert exception_results["missing_channel"].startswith("KeyNotFoundException")
    assert exception_results["length_mismatch"].startswith("ArgumentException")
    assert exception_results["session_alive_after"] == "True"


def test_fatal_dotnet_error_kills_the_process() -> None:
    exit_code: int = spike.check_crash_isolation()
    assert exit_code != 0
    if sys.platform != "win32":
        assert exit_code == -6
