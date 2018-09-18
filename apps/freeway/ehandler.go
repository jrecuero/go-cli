package freeway

import (
	"time"

	"github.com/jrecuero/go-cli/tools"
)

// RaceTime represents ...
type RaceTime int

// EHandler represents ...
type EHandler struct {
	race      *Race
	running   bool
	freeze    bool
	delay     int
	workerCbs []func(IDevice)
	rtime     RaceTime
	endgame   []IDevice
}

// tick is ...
func (ehdlr *EHandler) tick() {
	ehdlr.rtime++
}

// SetDelay is ...
func (ehdlr *EHandler) SetDelay(delay int) *EHandler {
	ehdlr.delay = delay
	return ehdlr
}

// AddWorkerCb is ...
func (ehdlr *EHandler) AddWorkerCb(cb func(IDevice)) *EHandler {
	ehdlr.workerCbs = append(ehdlr.workerCbs, cb)
	return ehdlr
}

// SetRace is ...
func (ehdlr *EHandler) SetRace(race *Race) *EHandler {
	ehdlr.race = race
	return ehdlr
}

// GetRace is ...
func (ehdlr *EHandler) GetRace() *Race {
	return ehdlr.race
}

// callWorkerCbs is ...
func (ehdlr *EHandler) callWorkerCbs(device IDevice) {
	for _, cb := range ehdlr.workerCbs {
		if cb != nil {
			cb(device)
		}
	}
}

// worker is ...
func (ehdlr *EHandler) worker(device IDevice) {
	device.SetRunning(true)
	for ehdlr.running {
		if !device.GetRunning() {
			return
		}
		if !ehdlr.freeze {
			device.FreewayTraverse()
			//tools.ToDisplay("worker device: %s\n", device)
		}
		ehdlr.callWorkerCbs(device)
		time.Sleep(time.Duration(ehdlr.delay) * time.Millisecond)
	}
}

// Setup is ...
func (ehdlr *EHandler) Setup() {
	for _, device := range ehdlr.GetRace().GetDevices() {
		device.NewLocation(ehdlr.GetRace().GetFreeway())
		tools.ToDisplay("setup device %s\n", device)
	}
}

// Start is ...
func (ehdlr *EHandler) Start() {
	ehdlr.running = true
	for _, device := range ehdlr.GetRace().GetDevices() {
		go ehdlr.worker(device)
	}
	tools.ToDisplay("EHandler started ...\n")
	ehdlr.AddWorkerCb(func(dev IDevice) {
		//tools.ToDisplay("device %#v: %d/%d\n", dev.GetName(), dev.Location().GetLaps(), ehdlr.GetRace().GetLaps())
		if dev.Location().GetLaps() == ehdlr.GetRace().GetLaps() {
			dev.SetRunning(false)
			ehdlr.endgame = append(ehdlr.endgame, dev)
			if len(ehdlr.endgame) == 1 {
				tools.ToDisplay("[%d] winner: %#v time: %#v\n", len(ehdlr.endgame), dev.GetName(), ehdlr.rtime)
			} else {
				tools.ToDisplay("[%d] device: %#v time: %#v\n", len(ehdlr.endgame), dev.GetName(), ehdlr.rtime)
			}
			if len(ehdlr.endgame) == len(ehdlr.GetRace().GetDevices()) {
				tools.ToDisplay("Race ended\n")
				ehdlr.Stop()
			}
		}
	})
	for ehdlr.running {
		time.Sleep(time.Duration(ehdlr.delay) * time.Millisecond)
		ehdlr.tick()
	}
	tools.ToDisplay("EHandler stopped\n")
	tools.ToDisplay("%s\n", ehdlr.endgame)
}

// Stop is ...
func (ehdlr *EHandler) Stop() {
	ehdlr.running = false
}

// SetFreeze is ...
func (ehdlr *EHandler) SetFreeze(freeze bool) {
	ehdlr.freeze = freeze
}

// NewEHandler is ...
func NewEHandler() *EHandler {
	return &EHandler{
		running: false,
		delay:   10,
	}
}
