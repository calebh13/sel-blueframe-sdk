using System;
using System.Collections.Generic;
using System.Threading;

namespace StandInHub
{
    // A small in-memory stand-in for a pub/sub hub client. It has the same call
    // shapes the Python SDK needs (generic publish, history fetch, callbacks on a
    // background thread, exceptions) without using any SEL code.
    public class LocalHub
    {
        private readonly object hubLock = new object();
        private readonly Dictionary<string, List<double>> samplesByChannel = new Dictionary<string, List<double>>();
        private readonly Dictionary<string, List<Action<long, double>>> subscribersByChannel = new Dictionary<string, List<Action<long, double>>>();

        public int Ping(int number)
        {
            return number + 1;
        }

        public int Publish<T>(string channelName, long[] timestamps, T[] values) where T : struct, IConvertible
        {
            if (timestamps.Length != values.Length)
            {
                throw new ArgumentException("timestamps and values must be the same length");
            }

            List<Action<long, double>> subscribers;
            lock (hubLock)
            {
                if (!samplesByChannel.ContainsKey(channelName))
                {
                    samplesByChannel[channelName] = new List<double>();
                }
                List<double> stored = samplesByChannel[channelName];
                for (int sampleIndex = 0; sampleIndex < values.Length; sampleIndex++)
                {
                    stored.Add(values[sampleIndex].ToDouble(null));
                }
                subscribers = GetSubscribers(channelName);
            }

            if (subscribers.Count > 0)
            {
                double[] copied = new double[values.Length];
                for (int sampleIndex = 0; sampleIndex < values.Length; sampleIndex++)
                {
                    copied[sampleIndex] = values[sampleIndex].ToDouble(null);
                }
                Thread sender = new Thread(() => SendToSubscribers(subscribers, timestamps, copied));
                sender.IsBackground = true;
                sender.Start();
            }
            return values.Length;
        }

        public double[] Fetch(string channelName)
        {
            lock (hubLock)
            {
                if (!samplesByChannel.ContainsKey(channelName))
                {
                    throw new KeyNotFoundException("no channel named " + channelName);
                }
                return samplesByChannel[channelName].ToArray();
            }
        }

        public void Subscribe(string channelName, Action<long, double> onSample)
        {
            lock (hubLock)
            {
                if (!subscribersByChannel.ContainsKey(channelName))
                {
                    subscribersByChannel[channelName] = new List<Action<long, double>>();
                }
                subscribersByChannel[channelName].Add(onSample);
            }
        }

        public void CrashProcess()
        {
            Environment.FailFast("stand-in for a fatal error inside the client");
        }

        private List<Action<long, double>> GetSubscribers(string channelName)
        {
            if (!subscribersByChannel.ContainsKey(channelName))
            {
                return new List<Action<long, double>>();
            }
            return new List<Action<long, double>>(subscribersByChannel[channelName]);
        }

        private static void SendToSubscribers(List<Action<long, double>> subscribers, long[] timestamps, double[] values)
        {
            for (int sampleIndex = 0; sampleIndex < values.Length; sampleIndex++)
            {
                foreach (Action<long, double> onSample in subscribers)
                {
                    onSample(timestamps[sampleIndex], values[sampleIndex]);
                }
            }
        }
    }
}
