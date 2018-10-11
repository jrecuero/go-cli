package battle

import (
	"bytes"
	"fmt"
)

// ActorInfo represents ...
type ActorInfo struct {
	name string
	desc string
}

// GetName is ...
func (info *ActorInfo) GetName() string {
	return info.name
}

// GetDescription is ...
func (info *ActorInfo) GetDescription() string {
	return info.desc
}

// NewActorInfo is ...
func NewActorInfo(name string, desc string) *ActorInfo {
	return &ActorInfo{name, desc}
}

// IActor represents
type IActor interface {
	ITechniqueHandler
	IStyleHandler
	IStanceHandler
	IAmoveHandler
	GetName() string
	GetDescription() string
	GetStats() *Stats
	SetDefaultAsTechnique()
}

// Actor represents ...
type Actor struct {
	*TechniqueHandler
	*StyleHandler
	*StanceHandler
	*AmoveHandler
	name  string
	desc  string
	stats *Stats
}

// GetName is ...
func (actor *Actor) GetName() string {
	return actor.name
}

// GetDescription is ...
func (actor *Actor) GetDescription() string {
	return actor.desc
}

// GetStats is ...
func (actor *Actor) GetStats() *Stats {
	return actor.stats
}

// SetDefaultAsTechnique is ...
func (actor *Actor) SetDefaultAsTechnique() {
	for _, tech := range actor.GetTechniques() {
		if tech.IsDefault() {
			actor.SetTechnique(tech)
			return
		}
	}
}

// AddTechnique is ...
func (actor *Actor) AddTechnique(techs ...ITechnique) bool {
	for _, tech := range techs {
		if len(actor.GetTechniques()) == 0 {
			tech.SetAsDefault(true)
		}
		if !actor.TechniqueHandler.AddTechnique(tech) {
			return false
		}
		if !actor.AddStyle(tech.GetStyles()...) {
			return false
		}
	}
	return true
}

// AddStyle is ...
func (actor *Actor) AddStyle(styles ...IStyle) bool {
	for _, style := range styles {
		if !actor.StyleHandler.AddStyle(style) {
			return false
		}
		if !actor.AddStance(style.GetStances()...) {
			return false
		}
	}
	return true
}

// AddStance is ...
func (actor *Actor) AddStance(stances ...IStance) bool {
	for _, stance := range stances {
		if !actor.StanceHandler.AddStance(stance) {
			return false
		}
		if !actor.AddAmove(stance.GetAmoves()...) {
			return false
		}
	}
	return true
}

// AddAmove is ...
func (actor *Actor) AddAmove(amoves ...IAmove) bool {
	for _, amove := range amoves {
		if !actor.AmoveHandler.AddAmove(amove) {
			return false
		}
	}
	return true
}

// String is ...
func (actor *Actor) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("-->%s\n", actor.GetName()))
	buf.WriteString(fmt.Sprintf("%s", actor.GetStats()))
	return buf.String()
}

// NewActor is ...
func NewActor(name string, desc string, stats *Stats) *Actor {
	return &Actor{
		TechniqueHandler: NewTechniqueHandler(),
		StyleHandler:     NewStyleHandler(),
		StanceHandler:    NewStanceHandler(),
		AmoveHandler:     NewAmoveHandler(),
		name:             name,
		desc:             desc,
		stats:            stats,
	}
}

// NewBasicActor is ...
func NewBasicActor(name string, desc string) *Actor {
	return &Actor{
		name: name,
		desc: desc,
	}
}
