package freeway

import "fmt"

// IDriver represents ...
type IDriver interface {
	GetName() string
}

// Driver represents ...
type Driver struct {
	name string
}

// GetName is ...
func (driver *Driver) GetName() string {
	return driver.name
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
