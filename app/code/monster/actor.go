package monster

import "fmt"

// IActor to be replaced ...
type IActor interface {
	GetName() string
	GetSpeed() *Speed
	Reset() IActor
}

// Actor represents ...
type Actor struct {
	name  string
	speed *Speed
}

// GetName is ...
func (actor *Actor) GetName() string {
	return actor.name
}

// GetSpeed is ...
func (actor *Actor) GetSpeed() *Speed {
	return actor.speed
}

// Reset is ...
func (actor *Actor) Reset() IActor {
	actor.GetSpeed().Reset()
	return actor
}

// String is ...
func (actor *Actor) String() string {
	return fmt.Sprintf("%#v %s", actor.GetName(), actor.GetSpeed())
}

// NewActor is ...
func NewActor(name string, speed *Speed) *Actor {
	return &Actor{
		name:  name,
		speed: speed,
	}
}
