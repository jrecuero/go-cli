package main

import (
	"bytes"
	"time"

	"github.com/jrecuero/go-cli/apps/freeway"
	"github.com/jrecuero/go-cli/tools"
)

func display(dots int) {
	var buffer bytes.Buffer
	for i := 0; i < dots; i++ {
		buffer.WriteString(".")
	}
	tools.ToDisplay("%s", freeway.MoveUpperLeft(0))
	tools.ToDisplay("one%s\n", buffer.String())
	tools.ToDisplay("two%s\n", buffer.String())
	tools.ToDisplay("three%s\n", buffer.String())
}

func displayDevices(devices []*freeway.Device) {
	tools.ToDisplay("%s", freeway.ClearEntireScreen())
	tools.ToDisplay("%s", freeway.MoveUpperLeft(0))
	for _, dev := range devices {
		//tools.ToDisplay("%s %d:%d\n", dev.GetName(), dev.Location().GetLaps(), dev.Location().Int())
		//tools.ToDisplay("%s %d:%d\n", dev.GetName(), dev.Location().GetLaps(), dev.Location().LapInt())
		var buffer bytes.Buffer
		for i := 0; i < dev.Location().LapInt()/3; i++ {
			buffer.WriteString("*")
		}
		tools.ToDisplay("%s [%d]%s\n", dev.GetName(), dev.Location().GetLaps(), buffer.String())
	}
}

func main() {
	tools.ToDisplay("%s", freeway.ClearEntireScreen())
	tools.ToDisplay("%s", freeway.MoveUpperLeft(0))
	tools.ToDisplay("%s", freeway.FgYellow())
	ehdlr := freeway.NewEHandler()
	race := freeway.NewRace()
	fway := freeway.NewFreeway()
	fway.AddSection(freeway.NewSection(100, 1, freeway.Straight, nil, nil, nil))
	fway.AddSection(freeway.NewSection(50, 1, freeway.Turn, nil, nil, nil))
	fway.AddSection(freeway.NewSection(100, 1, freeway.Straight, nil, nil, nil))
	fway.AddSection(freeway.NewSection(50, 1, freeway.Turn, nil, nil, nil))
	race.SetFreeway(fway)
	devices := []*freeway.Device{
		freeway.NewFullDevice("dev-80", "dev-class", "dev-sub", 80),
		freeway.NewFullDevice("dev-50", "dev-class", "dev-sub", 50),
		freeway.NewFullDevice("dev-90", "dev-class", "dev-sub", 90),
		freeway.NewFullDevice("dev-60", "dev-class", "dev-sub", 60),
		freeway.NewFullDevice("dev-70", "dev-class", "dev-sub", 70),
	}
	for _, d := range devices {
		race.AddDevice(d)
	}
	race.SetLaps(5)
	ehdlr.SetRace(race)
	ehdlr.Setup()
	ehdlr.SetDelay(100)
	displayDevices(devices)
	go func() {
		for ehdlr.IsRunning() {
			time.Sleep(1 * time.Second)
			displayDevices(devices)
		}
	}()
	ehdlr.Start()
	tools.ToDisplay("%s", freeway.Reset())
}
