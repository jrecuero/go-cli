package freeway

import (
	"fmt"
	"math/rand"
)

// IDevice represents ...
type IDevice interface {
	GetName() string
	GetClass() string
	GetSubClass() string
	GetPower() int
	Location() *Location
	NewLocation(*Freeway) (ISection, int)
	Traversing() int
	Entering(int) int
	Exiting(int) int
	FreewayTraverse()
	GetDriver() IDriver
	SetDriver(IDriver) IDevice
	GetRunning() bool
	SetRunning(bool) IDevice
}

// Device represents ...
type Device struct {
	name      string
	dclass    string
	dsubclass string
	power     int
	location  *Location
	driver    IDriver
	running   bool
	Dry       bool
}

// GetName is ...
func (dev *Device) GetName() string {
	return dev.name
}

// GetClass is ...
func (dev *Device) GetClass() string {
	return dev.dclass
}

// GetSubClass is ...
func (dev *Device) GetSubClass() string {
	return dev.dsubclass
}

// GetPower is ...
func (dev *Device) GetPower() int {
	return dev.power
}

// Location is ...
func (dev *Device) Location() *Location {
	return dev.location
}

// getLocation is ...
func (dev *Device) getLocation() (ISection, int) {
	return dev.location.GetLocation()
}

// getLocationIndex is ...
func (dev *Device) getLocationIndex() (int, int) {
	return dev.location.GetLocationIndex()
}

// NewLocation is ...
func (dev *Device) NewLocation(freeway *Freeway) (ISection, int) {
	dev.location = NewLocation(freeway)
	return dev.getLocation()
}

// diceit is ...
func (dev *Device) diceit(todice int) int {
	if !dev.Dry {
		dice := rand.Float32()
		return int(float32(todice) * dice)
	}
	return todice
}

// Traversing is ...
func (dev *Device) Traversing() int {
	section, _ := dev.getLocation()
	spec := section.GetSpec()
	travTotal := 1
	if dev.GetDriver() != nil {
		travTotal = dev.GetDriver().Traversing(spec)
	}
	return travTotal * dev.diceit(section.Traversing()*dev.GetPower())
}

// Entering is ...
func (dev *Device) Entering(speed int) int {
	section, _ := dev.getLocation()
	spec := section.GetSpec()
	travTotal := 1
	if dev.GetDriver() != nil {
		travTotal = dev.GetDriver().Entering(spec, speed)
	}
	return travTotal * dev.diceit(section.Entering(speed)*speed)
}

// Exiting is ...
func (dev *Device) Exiting(speed int) int {
	section, _ := dev.getLocation()
	spec := section.GetSpec()
	travTotal := 1
	if dev.GetDriver() != nil {
		travTotal = dev.GetDriver().Exiting(spec, speed)
	}
	return travTotal * dev.diceit(section.Exiting(speed)*speed)
}

// GetDriver is ...
func (dev *Device) GetDriver() IDriver {
	return dev.driver
}

// SetDriver is ...
func (dev *Device) SetDriver(driver IDriver) IDevice {
	dev.driver = driver
	return dev
}

// GetRunning is ...
func (dev *Device) GetRunning() bool {
	return dev.running
}

// SetRunning is ...
func (dev *Device) SetRunning(running bool) IDevice {
	dev.running = running
	return dev
}

// FreewayTraverse is ..
func (dev *Device) FreewayTraverse() {
	section, devpos := dev.getLocation()
	devSpeed := dev.Traversing()
	//tools.ToDisplay("traversing %s : %d\n", dev.GetName(), devSpeed)
	position := devpos + devSpeed
	for position >= section.GetLen() {
		nextSpeed := position - section.GetLen()
		exitSpeed := dev.Exiting(nextSpeed)
		section, _ = dev.location.NextSection()
		position = dev.Entering(exitSpeed)
		//tools.ToDisplay("new position %s next: %d exit: %d position: %d\n", dev.GetName(), nextSpeed, exitSpeed, position)
	}
	dev.location.SetPos(position)
}

// String is ...
func (dev *Device) String() string {
	isect, pos := dev.getLocationIndex()
	return fmt.Sprintf("name: %#v class|sub: %#v|%#v power: %d loc: %d-%d driver: %#v\n",
		dev.GetName(), dev.GetClass(), dev.GetSubClass(), dev.GetPower(), isect, pos, dev.GetDriver())
}

// NewDevice is ...
func NewDevice(name string) *Device {
	return &Device{
		name: name,
	}
}

// NewFullDevice is ...
func NewFullDevice(name string, dclass string, dsubclass string, power int) *Device {
	d := NewDevice(name)
	d.name = name
	d.dclass = dclass
	d.dsubclass = dsubclass
	d.power = power
	return d
}
