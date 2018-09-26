package freeway_test

import (
	"reflect"
	"testing"

	"github.com/jrecuero/go-cli/app/freeway"
)

func TestRace_Race(t *testing.T) {
	race := freeway.NewRace()
	if race.GetFreeway() != nil {
		t.Errorf("NewRace: freeway mismatch: exp: %v got: %v\n", nil, race.GetFreeway())
	}
	if len(race.GetDevices()) != 0 {
		t.Errorf("NewRace: device length mismatch: exp %d got%d\n", 0, len(race.GetDevices()))
	}
	if race.GetLaps() != 0 {
		t.Errorf("NewRace: laps length mismatch: exp %d got%d\n", 0, race.GetLaps())
	}
}

func TestRace_Freeway(t *testing.T) {
	race := freeway.NewRace()
	fway := freeway.NewFreeway()
	fway.AddSection(freeway.NewSection(100, 1, freeway.Straight, nil, nil, nil))
	fway.AddSection(freeway.NewSection(50, 1, freeway.Turn, nil, nil, nil))
	fway.AddSection(freeway.NewSection(100, 1, freeway.Straight, nil, nil, nil))
	fway.AddSection(freeway.NewSection(50, 1, freeway.Turn, nil, nil, nil))
	if race.SetFreeway(fway) == nil {
		t.Errorf("SetFreeway: code return: nil\n")
	}
	if readFway := race.GetFreeway(); !reflect.DeepEqual(readFway, fway) {
		t.Errorf("GetFreeway: freeway mismatch:\n\texp: %#v\n\tgot: %#v\n", readFway, fway)
	}
}

func TestRace_Laps(t *testing.T) {
	race := freeway.NewRace()
	if race.SetLaps(10) == nil {
		t.Errorf("SetLaps: code return: nil\n")
	}
	if race.GetLaps() != 10 {
		t.Errorf("SetLaps: laps length mismatch: exp %d got%d\n", 10, race.GetLaps())
	}

}

func TestRace_Devices(t *testing.T) {
	race := freeway.NewRace()
	devices := []freeway.IDevice{
		freeway.NewFullDevice("dev-100", "dev-class", "dev-sub", 100),
		freeway.NewFullDevice("dev-200", "dev-class", "dev-sub", 200),
		freeway.NewFullDevice("dev-300", "dev-class", "dev-sub", 300),
	}
	for _, dev := range devices {
		if race.AddDevice(dev) == nil {
			t.Errorf("AddDevice: return code: nil\n")
		}
	}
	devnameMatch := []bool{false, false, false}
	devMatch := []bool{false, false, false}
	for devname, dev := range race.GetDevices() {
		for i, expdev := range devices {
			expname := expdev.GetName()
			if expname == devname {
				if devnameMatch[i] {
					t.Errorf("GetDevices: duplicated device name: %#v\n", devname)
				}
				devnameMatch[i] = true
			}
			if reflect.DeepEqual(expdev, dev) {
				if devMatch[i] {
					t.Errorf("GetDevices: duplicated device: %#v\n", dev)
				}
				devMatch[i] = true
			}
			if expname == devname && reflect.DeepEqual(expdev, dev) {
				break
			}
		}
	}
	for i, matched := range devnameMatch {
		if !matched {
			t.Errorf("GetDevices: device name not found: %#v\n", devices[i].GetName())
		}
	}
	for i, matched := range devMatch {
		if !matched {
			t.Errorf("GetDevices: device not found: %#v\n", devices[i])
		}
	}

	for _, expdev := range devices {
		if gotdev := race.GetDeviceByName(expdev.GetName()); !reflect.DeepEqual(expdev, gotdev) {
			t.Errorf("GetDeviceByName: device mismatch:\n\texp: %#v\n\tgot: %#v\n", expdev, gotdev)
		}
	}
}
