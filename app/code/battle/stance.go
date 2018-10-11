package battle

import (
	"bytes"
	"fmt"
)

// IStance represents ...
type IStance interface {
	IBase
	GetStyle() IStyle
	GetTechnique() ITechnique
	GetAmoves() []IAmove
	AddAmove(...IAmove) bool
	RemoveAmove(...IAmove) bool
	RemoveAmoveByName(string) bool
	GetAmoveByName(string) IAmove
}

// Stance represents ...
type Stance struct {
	*Base
	parent IStyle
	amoves []IAmove
}

// GetStyle is ...
func (stance *Stance) GetStyle() IStyle {
	return stance.parent
}

// GetTechnique is ...
func (stance *Stance) GetTechnique() ITechnique {
	return stance.parent.GetTechnique()
}

// GetDN is ...
func (stance *Stance) GetDN() string {
	return fmt.Sprintf("%s:%s",
		stance.GetStyle().GetDN(),
		stance.Base.GetDN())
}

// GetAmoves is ...
func (stance *Stance) GetAmoves() []IAmove {
	return stance.amoves
}

// AddAmove is ...
func (stance *Stance) AddAmove(amoves ...IAmove) bool {
	stance.amoves = append(stance.amoves, amoves...)
	return true
}

// RemoveAmove is ...
func (stance *Stance) RemoveAmove(amoves ...IAmove) bool {
	for _, amove := range amoves {
		if !stance.RemoveAmoveByName(amove.GetName()) {
			return false
		}
	}
	return true
}

// RemoveAmoveByName is ...
func (stance *Stance) RemoveAmoveByName(name string) bool {
	for i, amove := range stance.amoves {
		if amove.GetName() == name {
			stance.amoves = append(stance.amoves[:i], stance.amoves[i+1:]...)
			return true
		}
	}
	return false
}

// GetAmoveByName is ...
func (stance *Stance) GetAmoveByName(name string) IAmove {
	for _, amove := range stance.amoves {
		if amove.GetName() == name {
			return amove
		}
	}
	return nil
}

// String is ...
func (stance *Stance) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("Stance %s", stance.Base))
	for _, amove := range stance.amoves {
		buf.WriteString(fmt.Sprintf("\n      # %s", amove))
	}
	return buf.String()
}

// NewStance is ...
func NewStance(name string, parent IStyle) *Stance {
	stance := &Stance{
		Base:   NewBase(name),
		parent: parent,
	}
	parent.AddStance(stance)
	return stance
}

// NewFullStance is  ...
func NewFullStance(name string, desc string, ustats *UStats, parent IStyle) *Stance {
	stance := NewStance(name, parent)
	stance.Base = NewFullBase(name, desc, ustats)
	return stance
}

// IStanceHandler represents ...
type IStanceHandler interface {
	GetStance() IStance
	SetStance(IStance) bool
	SetStanceByName(string) bool
	GetStances() []IStance
	SetStances([]IStance)
	AddStance(...IStance) bool
	RemoveStance(...IStance) bool
	RemoveStanceByName(string) bool
	GetStanceByName(string) IStance
}

// StanceHandler represents ...
type StanceHandler struct {
	stances []IStance
	stance  IStance
}

// GetStance is ...
func (th *StanceHandler) GetStance() IStance {
	return th.stance
}

// SetStance is ...
func (th *StanceHandler) SetStance(stance IStance) bool {
	return th.SetStanceByName(stance.GetName())
}

// SetStanceByName is ...
func (th *StanceHandler) SetStanceByName(name string) bool {
	stance := th.GetStanceByName(name)
	if stance != nil {
		th.stance = stance
		return true
	}
	return false
}

// GetStances is ...
func (th *StanceHandler) GetStances() []IStance {
	return th.stances
}

// SetStances iss ...
func (th *StanceHandler) SetStances(stances []IStance) {
	th.stances = stances
}

// AddStance is ...
func (th *StanceHandler) AddStance(stances ...IStance) bool {
	th.stances = append(th.stances, stances...)
	return true
}

// RemoveStance is ...
func (th *StanceHandler) RemoveStance(stances ...IStance) bool {
	for _, stance := range stances {
		if !th.RemoveStanceByName(stance.GetName()) {
			return false
		}
	}
	return true
}

// RemoveStanceByName is ...
func (th *StanceHandler) RemoveStanceByName(name string) bool {
	for index, stance := range th.stances {
		if stance.GetName() == name {
			th.stances = append(th.stances[:index], th.stances[index+1:]...)
			return true
		}
	}
	return false
}

// GetStanceByName is ...
func (th *StanceHandler) GetStanceByName(name string) IStance {
	for _, stance := range th.stances {
		if stance.GetName() == name {
			return stance
		}
	}
	return nil
}

// NewStanceHandler is ...
func NewStanceHandler() *StanceHandler {
	return &StanceHandler{}
}
