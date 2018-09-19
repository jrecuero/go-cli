package monster

import "fmt"

// IActor to be replaced ...
type IActor interface {
	GetName() string
	GetSpeed() int
}

// Actor represents ...
type Actor struct {
	name  string
	speed int
}

// GetName is ...
func (actor *Actor) GetName() string {
	return actor.name
}

// GetSpeed is ...
func (actor *Actor) GetSpeed() int {
	return actor.speed
}

// String is ...
func (actor *Actor) String() string {
	return fmt.Sprintf("%#v:%d", actor.GetName(), actor.GetSpeed())
}

// NewActor is ...
func NewActor(name string, speed int) *Actor {
	return &Actor{
		name:  name,
		speed: speed,
	}
}
