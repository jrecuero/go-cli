package monster

import "fmt"

// Speed represents ...
type Speed struct {
	nominal    int
	next       int
	processing int
}

// Get is ...
func (speed *Speed) Get() int {
	return speed.nominal
}

// GetNext is ...
func (speed *Speed) GetNext() int {
	return speed.next
}

// GetProcessing is ...
func (speed *Speed) GetProcessing() int {
	return speed.processing
}

// SetNext is ...
func (speed *Speed) SetNext(next int) {
	speed.next = next
}

// SetProcessing is ...
func (speed *Speed) SetProcessing(processing int) {
	speed.processing = processing
	speed.next = processing
}

// Reset is ...
func (speed *Speed) Reset() {
	speed.next = speed.Get()
	speed.processing = speed.Get()
}

// String is ...
func (speed *Speed) String() string {
	return fmt.Sprintf("speed:%d-next:%d-proc:%d", speed.Get(), speed.GetNext(), speed.GetProcessing())
}

// NewSpeed is ...
func NewSpeed(nominal int) *Speed {
	return &Speed{
		nominal:    nominal,
		next:       nominal,
		processing: nominal,
	}
}
