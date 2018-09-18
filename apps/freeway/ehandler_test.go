package freeway_test

import (
	"testing"
	"time"

	"github.com/jrecuero/go-cli/apps/freeway"
)

func TestEHandler_EHandler(t *testing.T) {
	ehdlr := freeway.NewEHandler()
	if ehdlr.GetRace() != nil {
		t.Errorf("NewEHandler: race mismatch: exp: %v got: %v\n", nil, ehdlr.GetRace())
	}
}

func TestEHandler_SetupAndStart(t *testing.T) {
	ehdlr := freeway.NewEHandler()
	race := freeway.NewRace()
	fway := freeway.NewFreeway()
	fway.AddSection(freeway.NewSection(100, 1, "line", nil, nil, nil))
	fway.AddSection(freeway.NewSection(50, 1, "turn", nil, nil, nil))
	fway.AddSection(freeway.NewSection(100, 1, "line", nil, nil, nil))
	fway.AddSection(freeway.NewSection(50, 1, "turn", nil, nil, nil))
	race.SetFreeway(fway)
	race.AddDevice(freeway.NewFullDevice("dev-80", "dev-class", "dev-sub", 80))
	race.SetLaps(10)
	ehdlr.SetRace(race)
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
	ehdlr.AddWorkerCb(func(dev freeway.IDevice) {
		isect, pos := dev.Location().GetLocationIndex()
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

func TestEHandler_EndRace(t *testing.T) {
	ehdlr := freeway.NewEHandler()
	race := freeway.NewRace()
	fway := freeway.NewFreeway()
	fway.AddSection(freeway.NewSection(100, 1, "line", nil, nil, nil))
	fway.AddSection(freeway.NewSection(50, 1, "turn", nil, nil, nil))
	fway.AddSection(freeway.NewSection(100, 1, "line", nil, nil, nil))
	fway.AddSection(freeway.NewSection(50, 1, "turn", nil, nil, nil))
	race.SetFreeway(fway)
	race.AddDevice(freeway.NewFullDevice("dev-80", "dev-class", "dev-sub", 80))
	race.AddDevice(freeway.NewFullDevice("dev-75", "dev-class", "dev-sub", 75))
	race.AddDevice(freeway.NewFullDevice("dev-50", "dev-class", "dev-sub", 50))
	race.AddDevice(freeway.NewFullDevice("dev-90", "dev-class", "dev-sub", 90))
	race.AddDevice(freeway.NewFullDevice("dev-45", "dev-class", "dev-sub", 45))
	race.AddDevice(freeway.NewFullDevice("dev-55", "dev-class", "dev-sub", 55))
	race.AddDevice(freeway.NewFullDevice("dev-85", "dev-class", "dev-sub", 85))
	race.AddDevice(freeway.NewFullDevice("dev-60", "dev-class", "dev-sub", 60))
	race.AddDevice(freeway.NewFullDevice("dev-65", "dev-class", "dev-sub", 65))
	race.AddDevice(freeway.NewFullDevice("dev-70", "dev-class", "dev-sub", 70))
	race.SetLaps(5)
	ehdlr.SetRace(race)
	ehdlr.Setup()
	ehdlr.SetDelay(1000)
	go func() {
		time.Sleep(100 * 1000 * time.Millisecond)
		ehdlr.Stop()
	}()
	ehdlr.Start()
}
