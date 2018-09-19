package monster

import "fmt"

// IActor to be replaced ...
type IActor interface {
	GetName() string
	GetSpeed() int
	GetNext() int
	SetNext(int) IActor
	GetProcessing() int
	SetProcessing(int) IActor
	Reset() IActor
}

// Actor represents ...
type Actor struct {
	name       string
	speed      int
	next       int
	processing int
}

// GetName is ...
func (actor *Actor) GetName() string {
	return actor.name
}

// GetSpeed is ...
func (actor *Actor) GetSpeed() int {
	return actor.speed
}

// GetNext is ...
func (actor *Actor) GetNext() int {
	return actor.next
}

// SetNext is ...
func (actor *Actor) SetNext(next int) IActor {
	actor.next = next
	return actor
}

// GetProcessing is ...
func (actor *Actor) GetProcessing() int {
	return actor.processing
}

// SetProcessing is ...
func (actor *Actor) SetProcessing(processing int) IActor {
	actor.processing = processing
	actor.SetNext(processing)
	return actor
}

// Reset is ...
func (actor *Actor) Reset() IActor {
	actor.SetProcessing(actor.GetSpeed())
	return actor
}

// String is ...
func (actor *Actor) String() string {
	return fmt.Sprintf("%#v:%d", actor.GetName(), actor.GetSpeed())
}

// NewActor is ...
func NewActor(name string, speed int) *Actor {
	return &Actor{
		name:       name,
		speed:      speed,
		next:       speed,
		processing: speed,
	}
}
