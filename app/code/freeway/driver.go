package freeway

import "fmt"

// IDriver represents ...
type IDriver interface {
	GetName() string
	Traversing(spec Spec) int
	Entering(spec Spec, speed int) int
	Exiting(spec Spec, speed int) int
}

// Driver represents ...
type Driver struct {
	name string
}

// GetName is ...
func (driver *Driver) GetName() string {
	return driver.name
}

// Traversing is ...
func (driver *Driver) Traversing(spec Spec) int {
	return 1
}

// Entering is ...
func (driver *Driver) Entering(spec Spec, speed int) int {
	return 1
}

// Exiting is ...
func (driver *Driver) Exiting(spec Spec, speed int) int {
	return 1
}

// String is ...
func (driver *Driver) String() string {
	return fmt.Sprintf("name: %#v\n", driver.GetName())
}

// NewDriver is ...
func NewDriver(name string) *Driver {
	return &Driver{
		name: name,
	}
}
