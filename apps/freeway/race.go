package freeway

import (
	"github.com/jrecuero/go-cli/tools"
)

// Race represents ...
type Race struct {
	freeway *Freeway
	devices map[string]IDevice
	laps    int
}

// SetFreeway is ...
func (race *Race) SetFreeway(freeway *Freeway) *Race {
	race.freeway = freeway
	return race
}

// GetFreeway is ...
func (race *Race) GetFreeway() *Freeway {
	return race.freeway
}

// SetLaps is ...
func (race *Race) SetLaps(laps int) *Race {
	race.laps = laps
	return race
}

// GetLaps is ...
func (race *Race) GetLaps() int {
	return race.laps
}

// AddDevice is ...
func (race *Race) AddDevice(dev IDevice) *Race {
	race.devices[dev.GetName()] = dev
	return race
}

// GetDevices is ...
func (race *Race) GetDevices() map[string]IDevice {
	return race.devices
}

// GetDeviceByName is ...
func (race *Race) GetDeviceByName(name string) IDevice {
	return race.devices[name]
}

// Setup is ...
func (race *Race) Setup() {
	for _, device := range race.devices {
		device.NewLocation(race.GetFreeway())
		tools.ToDisplay("setup device %s\n", device)
	}
}

// NewRace is ...
func NewRace() *Race {
	return &Race{
		devices: make(map[string]IDevice),
	}
}
