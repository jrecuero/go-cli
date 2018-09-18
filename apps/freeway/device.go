package freeway

import (
	"fmt"
)

// IDevice represents ...
type IDevice interface {
	GetName() string
	GetClass() string
	GetSubClass() string
	GetPower() int
	GetLocation() (ISection, int)
	GetLocationIndex() (int, int)
	NewLocation(*Freeway) (ISection, int)
	Traversing() int
	Entering(int) int
	Exiting(int) int
	FreewayTraverse()
	GetDriver() IDriver
	SetDriver(IDriver) IDevice
}

// Device represents ...
type Device struct {
	name      string
	dclass    string
	dsubclass string
	power     int
	location  *Location
	driver    IDriver
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

// NewLocation is ...
func (dev *Device) NewLocation(freeway *Freeway) (ISection, int) {
	dev.location = NewLocation(freeway)
	return dev.GetLocation()
}

// GetLocation is ...
func (dev *Device) GetLocation() (ISection, int) {
	return dev.location.GetLocation()
}

// GetLocationIndex is ...
func (dev *Device) GetLocationIndex() (int, int) {
	return dev.location.GetLocationIndex()
}

// Traversing is ...
func (dev *Device) Traversing() int {
	section, _ := dev.GetLocation()
	//spec := section.GetSpec()
	speed := section.Traversing() * dev.GetPower()
	return speed
}

// Entering is ...
func (dev *Device) Entering(speed int) int {
	section, _ := dev.GetLocation()
	//spec := section.GetSpec()
	return section.Entering(speed) * speed
}

// Exiting is ...
func (dev *Device) Exiting(speed int) int {
	section, _ := dev.GetLocation()
	//spec := section.GetSpec()
	return section.Exiting(speed) * speed
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

// FreewayTraverse is ..
func (dev *Device) FreewayTraverse() {
	section, devpos := dev.GetLocation()
	devSpeed := dev.Traversing()
	//tools.ToDisplay("traversing %s : %d\n", dev.GetName(), devSpeed)
	position := devpos + devSpeed
	for position >= section.GetLen() {
		nextSpeed := position - section.GetLen()
		section, _ = dev.location.NextSection()
		exitSpeed := dev.Exiting(nextSpeed)
		position = dev.Entering(exitSpeed)
		//tools.ToDisplay("new position %s next: %d exit: %d position: %d\n", dev.GetName(), nextSpeed, exitSpeed, position)
	}
	dev.location.SetPos(position)
}

// String is ...
func (dev *Device) String() string {
	isect, pos := dev.GetLocationIndex()
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
