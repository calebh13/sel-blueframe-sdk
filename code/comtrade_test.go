package comtrade

import (
	"math"
	"os"
	"path/filepath"
	"testing"
)

// ---------------- WRITE TEST ----------------

func TestWriteCOMTRADE(t *testing.T) {
	ct := mockCOMTRADE()

	dir := t.TempDir()
	cfgFile := filepath.Join(dir, "test.cfg")
	datFile := filepath.Join(dir, "test.dat")

	if err := writeCFG(cfgFile, ct); err != nil {
		t.Fatal(err)
	}
	if err := writeDAT(datFile, ct); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		t.Fatal("CFG file not created")
	}
	if _, err := os.Stat(datFile); os.IsNotExist(err) {
		t.Fatal("DAT file not created")
	}
}

// ---------------- READ TEST ----------------

func TestReadCOMTRADE(t *testing.T) {
	ct := mockCOMTRADE()

	dir := t.TempDir()
	cfgFile := filepath.Join(dir, "test.cfg")
	datFile := filepath.Join(dir, "test.dat")

	if err := writeCFG(cfgFile, ct); err != nil {
		t.Fatal(err)
	}
	if err := writeDAT(datFile, ct); err != nil {
		t.Fatal(err)
	}

	readCT, err := readCFG(cfgFile)
	if err != nil {
		t.Fatal(err)
	}
	if err := readDAT(datFile, readCT); err != nil {
		t.Fatal(err)
	}

	if readCT.StationName != ct.StationName {
		t.Errorf("StationName mismatch")
	}
	if len(readCT.AnalogChannels) != len(ct.AnalogChannels) {
		t.Errorf("Analog channel count mismatch")
	}
	if len(readCT.Data) != len(ct.Data) {
		t.Errorf("Data length mismatch")
	}
}

// ---------------- ROUND TRIP ----------------

func TestRoundTrip(t *testing.T) {
	original := mockCOMTRADE()

	dir := t.TempDir()
	cfgFile := filepath.Join(dir, "round.cfg")
	datFile := filepath.Join(dir, "round.dat")

	if err := writeCFG(cfgFile, original); err != nil {
		t.Fatal(err)
	}
	if err := writeDAT(datFile, original); err != nil {
		t.Fatal(err)
	}

	readCT, err := readCFG(cfgFile)
	if err != nil {
		t.Fatal(err)
	}
	if err := readDAT(datFile, readCT); err != nil {
		t.Fatal(err)
	}

	if original.StationName != readCT.StationName {
		t.Errorf("StationName mismatch")
	}

	if len(original.Data) != len(readCT.Data) {
		t.Fatalf("Data length mismatch")
	}

	for i := range original.Data {
		for j := range original.Data[i] {
			if math.Abs(original.Data[i][j]-readCT.Data[i][j]) > 1e-6 {
				t.Errorf("Data mismatch at [%d][%d]: got %f, want %f",
					i, j, readCT.Data[i][j], original.Data[i][j])
			}
		}
	}
}
