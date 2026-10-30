import sys
import threading

import numpy as np
import gosignals as gs


def main() -> None:
    hub_port: int = int(sys.argv[1])
    sample_count: int = 600
    received_count: list[int] = [0]
    all_received = threading.Event()

    def on_data(timestamps: np.ndarray, values: np.ndarray) -> None:
        received_count[0] = received_count[0] + len(values)
        if received_count[0] >= sample_count:
            all_received.set()

    voltage = gs.Signal("SUB1.BUS1", "VA", rate=60, dtype="float32")
    with gs.connect("127.0.0.1", hub_port, name="example-sub", transport="bridge", unique_name=True) as subscriber_client:
        with gs.connect("127.0.0.1", hub_port, name="example-pub", transport="bridge", unique_name=True) as publisher_client:
            start_time: int = publisher_client.rates.floor(1, publisher_client.now())
            with subscriber_client.subscribe(voltage, start=start_time, on_data=on_data):
                seconds: np.ndarray = np.arange(sample_count) / 60.0
                with publisher_client.publisher(voltage, start=start_time) as publisher:
                    publisher.publish(start_time, np.sin(2 * np.pi * seconds))
                finished: bool = all_received.wait(timeout=10)

            timestamps, values = publisher_client.fetch(voltage, start=start_time, end=start_time + 10_000_000_000)

    print("published", sample_count, "samples")
    print("subscriber received", received_count[0], "samples, finished:", finished)
    print("fetch returned", len(values), "samples, first three", values[:3])


if __name__ == "__main__":
    main()
