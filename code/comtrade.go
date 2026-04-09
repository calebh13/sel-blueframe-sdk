package comtrade

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// ========================= TYPES =========================


type AnalogChannel struct {
	Index int
	Name  string
	Phase string
	Unit  string
	A     float64
	B     float64
}

type DigitalChannel struct {
	Index int
	Name  string
	Phase string
}

type COMTRADE struct {
	StationName string
	DeviceID    string
	RevYear     string

	AnalogChannels  []AnalogChannel
	DigitalChannels []DigitalChannel

	Frequency  float64
	SampleRate float64

	StartTime   time.Time
	TriggerTime time.Time

	Data [][]float64
}

// ========================= WRITER =========================

func writeCFG(filename string, ct COMTRADE) error {
	var sb strings.Builder

	totalChannels := len(ct.AnalogChannels) + len(ct.DigitalChannels)

	// Header
	sb.WriteString(fmt.Sprintf("%s,%s,%s\n", ct.StationName, ct.DeviceID, ct.RevYear))
	sb.WriteString(fmt.Sprintf("%d,%dA,%dD\n", totalChannels, len(ct.AnalogChannels), len(ct.DigitalChannels)))

	// Analog channels
	for _, ch := range ct.AnalogChannels {
		sb.WriteString(fmt.Sprintf("%d,%s,%s,,%s,%f,%f,0,0,0,1,1,P\n",
			ch.Index, ch.Name, ch.Phase, ch.Unit, ch.A, ch.B))
	}

	// Digital channels
	for _, ch := range ct.DigitalChannels {
		sb.WriteString(fmt.Sprintf("%d,%s,%s,,0\n",
			ch.Index, ch.Name, ch.Phase))
	}

	// Frequency & sampling
	sb.WriteString(fmt.Sprintf("%f\n", ct.Frequency))
	sb.WriteString("1\n") // number of sample rates
	sb.WriteString(fmt.Sprintf("%f,0\n", ct.SampleRate))

	// Timestamps
	timeFormat := "02/01/2006,15:04:05.000000"
	sb.WriteString(ct.StartTime.Format(timeFormat) + "\n")
	sb.WriteString(ct.TriggerTime.Format(timeFormat) + "\n")

	// File type
	sb.WriteString("ASCII\n")
	sb.WriteString("1.0\n")

	return os.WriteFile(filename, []byte(sb.String()), 0644)
}

func writeDAT(filename string, ct COMTRADE) error {
	var sb strings.Builder

	for i, row := range ct.Data {
		line := fmt.Sprintf("%d,%f", i+1, row[0]) // sample#, timestamp

		for _, v := range row[1:] {
			line += fmt.Sprintf(",%f", v)
		}
		sb.WriteString(line + "\n")
	}

	return os.WriteFile(filename, []byte(sb.String()), 0644)
}

// ========================= READER =========================

func readCFG(filename string) (*COMTRADE, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ct := &COMTRADE{}
	scanner := bufio.NewScanner(file)

	lineNum := 0
	var analogCount, digitalCount int

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineNum++

		switch lineNum {
		case 1:
			parts := strings.Split(line, ",")
			ct.StationName = parts[0]
			ct.DeviceID = parts[1]
			ct.RevYear = parts[2]

		case 2:
			parts := strings.Split(line, ",")
			for _, p := range parts {
				if strings.HasSuffix(p, "A") {
					analogCount, _ = strconv.Atoi(strings.TrimSuffix(p, "A"))
				}
				if strings.HasSuffix(p, "D") {
					digitalCount, _ = strconv.Atoi(strings.TrimSuffix(p, "D"))
				}
			}

		default:
			if len(ct.AnalogChannels) < analogCount {
				parts := strings.Split(line, ",")
				idx, _ := strconv.Atoi(parts[0])
				a, _ := strconv.ParseFloat(parts[5], 64)
				b, _ := strconv.ParseFloat(parts[6], 64)
				ct.AnalogChannels = append(ct.AnalogChannels, AnalogChannel{
					Index: idx,
					Name:  parts[1],
					Phase: parts[2],
					Unit:  parts[4],
					A:     a,
					B:     b,
				})
				continue
			}

			if len(ct.DigitalChannels) < digitalCount {
				parts := strings.Split(line, ",")
				idx, _ := strconv.Atoi(parts[0])
				ct.DigitalChannels = append(ct.DigitalChannels, DigitalChannel{
					Index: idx,
					Name:  parts[1],
					Phase: parts[2],
				})
				continue
			}

			if ct.Frequency == 0 {
				ct.Frequency, _ = strconv.ParseFloat(line, 64)
				continue
			}

			if line == "1" {
				continue
			}

			if ct.SampleRate == 0 {
				parts := strings.Split(line, ",")
				ct.SampleRate, _ = strconv.ParseFloat(parts[0], 64)
				continue
			}

			timeFormat := "02/01/2006,15:04:05.000000"
			if ct.StartTime.IsZero() {
				ct.StartTime, _ = time.Parse(timeFormat, line)
				continue
			}

			if ct.TriggerTime.IsZero() {
				ct.TriggerTime, _ = time.Parse(timeFormat, line)
				continue
			}
		}
	}

	return ct, scanner.Err()
}

func readDAT(filename string, ct *COMTRADE) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) < 2 {
			continue
		}

		row := []float64{}
		timestamp, _ := strconv.ParseFloat(parts[1], 64)
		row = append(row, timestamp)

		for _, p := range parts[2:] {
			val, _ := strconv.ParseFloat(p, 64)
			row = append(row, val)
		}

		ct.Data = append(ct.Data, row)
	}

	return scanner.Err()
}

// test helpers

func mockCOMTRADE() COMTRADE {
	return COMTRADE{
		StationName: "TEST_STATION",
		DeviceID:    "DEVICE_1",
		RevYear:     "2013",

		Frequency:  60.0,
		SampleRate: 1000.0,

		StartTime:   time.Now(),
		TriggerTime: time.Now(),

		AnalogChannels: []AnalogChannel{
			{1, "VA", "A", "V", 1.0, 0.0},
			{2, "VB", "B", "V", 1.0, 0.0},
		},
		DigitalChannels: []DigitalChannel{
			{1, "BREAKER", "A"},
		},

		Data: [][]float64{
			{0.000001, 120.1, 121.2, 1},
			{0.000002, 119.8, 120.9, 0},
		},
	}
}
