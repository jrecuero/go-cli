package freeway_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/jrecuero/go-cli/apps/freeway"
)

func TestEHandler_EHandler(t *testing.T) {
	ehdlr := freeway.NewEHandler()
	if ehdlr.GetFreeway() != nil {
		t.Errorf("NewEHandler: freeway mismatch: exp: %v got: %v\n", nil, ehdlr.GetFreeway())
	}
	if len(ehdlr.GetDevices()) != 0 {
		t.Errorf("NewEHandler: device length mismatch: exp %d got%d\n", 0, len(ehdlr.GetDevices()))
	}
}

func TestEHandler_Freeway(t *testing.T) {
	ehdlr := freeway.NewEHandler()
	fway := freeway.NewFreeway()
	fway.AddSection(freeway.NewSection(100, 1, "line", nil, nil, nil))
	fway.AddSection(freeway.NewSection(50, 1, "turn", nil, nil, nil))
	fway.AddSection(freeway.NewSection(100, 1, "line", nil, nil, nil))
	fway.AddSection(freeway.NewSection(50, 1, "turn", nil, nil, nil))
	if ehdlr.SetFreeway(fway) == nil {
		t.Errorf("SetFreeway: code return: nil\n")
	}
	if readFway := ehdlr.GetFreeway(); !reflect.DeepEqual(readFway, fway) {
		t.Errorf("GetFreeway: freeway mismatch:\n\texp: %#v\n\tgot: %#v\n", readFway, fway)
	}
}

func TestEHandler_Devices(t *testing.T) {
	ehdlr := freeway.NewEHandler()
	devices := []freeway.IDevice{
		freeway.NewFullDevice("dev-100", "dev-class", "dev-sub", 100),
		freeway.NewFullDevice("dev-200", "dev-class", "dev-sub", 200),
		freeway.NewFullDevice("dev-300", "dev-class", "dev-sub", 300),
	}
	for _, dev := range devices {
		if ehdlr.AddDevice(dev) == nil {
			t.Errorf("AddDevice: return code: nil\n")
		}
	}
	devnameMatch := []bool{false, false, false}
	devMatch := []bool{false, false, false}
	for devname, dev := range ehdlr.GetDevices() {
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
		if gotdev := ehdlr.GetDeviceByName(expdev.GetName()); !reflect.DeepEqual(expdev, gotdev) {
			t.Errorf("GetDeviceByName: device mismatch:\n\texp: %#v\n\tgot: %#v\n", expdev, gotdev)
		}
	}
}

func TestEHandler_SetupAndStart(t *testing.T) {
	ehdlr := freeway.NewEHandler()
	fway := freeway.NewFreeway()
	fway.AddSection(freeway.NewSection(100, 1, "line", nil, nil, nil))
	fway.AddSection(freeway.NewSection(50, 1, "turn", nil, nil, nil))
	fway.AddSection(freeway.NewSection(100, 1, "line", nil, nil, nil))
	fway.AddSection(freeway.NewSection(50, 1, "turn", nil, nil, nil))
	ehdlr.SetFreeway(fway)
	ehdlr.AddDevice(freeway.NewFullDevice("dev-80", "dev-class", "dev-sub", 80))
	ehdlr.Setup()
	ehdlr.SetDelay(1000)
	counter := 0
	expLocations := []struct {
		isect int
		pos   int
	}{
		{0, 80},
		{2, 10},
		{2, 90},
		{0, 20},
		{1, 0},
		{2, 30},
		{3, 10},
		{0, 40},
		{1, 20},
		{2, 50},
	}
	ehdlr.SetWorkerCb(func(dev freeway.IDevice) {
		isect, pos := dev.GetLocationIndex()
		//tools.ToDisplay("[%d] worker device: %d-%d (%v)\n", counter, isect, pos, expLocations[counter])
		if isect != expLocations[counter].isect {
			t.Errorf("SetupAndStart: section mismatch: exp: %d got:%d\n", expLocations[counter].isect, isect)
		}
		if pos != expLocations[counter].pos {
			t.Errorf("SetupAndStart: position mismatch: exp: %d got:%d\n", expLocations[counter].pos, pos)
		}
		counter++
	})
	go func() {
		time.Sleep(10 * 1000 * time.Millisecond)
		ehdlr.Stop()
	}()
	ehdlr.Start()
}
