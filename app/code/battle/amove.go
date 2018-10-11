package battle

import (
	"bytes"
	"fmt"
)

// Amode represents
type Amode int

const (
	// AmodeNone represents ...
	AmodeNone Amode = 0
	// AmodeAttack represents ...
	AmodeAttack Amode = 1
	// AmodeDefence represents ...
	AmodeDefence Amode = 2
)

// IAmove represents ...
type IAmove interface {
	IBase
	GetStance() IStance
	GetStyle() IStyle
	GetTechnique() ITechnique
	GetAmode() Amode
}

// Amove represents ...
type Amove struct {
	*Base
	parent IStance
	amode  Amode
}

//GetStance is ...
func (amove *Amove) GetStance() IStance {
	return amove.parent
}

// GetStyle is ...
func (amove *Amove) GetStyle() IStyle {
	return amove.parent.GetStyle()
}

// GetTechnique is ...
func (amove *Amove) GetTechnique() ITechnique {
	return amove.parent.GetTechnique()
}

// GetDN is ...
func (amove *Amove) GetDN() string {
	return fmt.Sprintf("%s:%s",
		amove.GetStance().GetDN(),
		amove.Base.GetDN())
}

// GetAmode is ...
func (amove *Amove) GetAmode() Amode {
	return amove.amode
}

// String is ...
func (amove *Amove) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("Amove %s", amove.Base))
	return buf.String()
}

// NewAmove is ...
func NewAmove(name string, parent IStance) *Amove {
	amove := &Amove{
		Base:   NewBase(name),
		parent: parent,
	}
	parent.AddAmove(amove)
	return amove
}

// NewFullAmove is  ...
func NewFullAmove(name string, desc string, ustats *UStats, parent IStance) *Amove {
	amove := NewAmove(name, parent)
	amove.Base = NewFullBase(name, desc, ustats)
	return amove
}

// IAmoveHandler represents ...
type IAmoveHandler interface {
	GetAmove() IAmove
	SetAmove(IAmove) bool
	SetAmoveByName(string) bool
	GetAmoves() []IAmove
	SetAmoves([]IAmove)
	AddAmove(...IAmove) bool
	RemoveAmove(...IAmove) bool
	RemoveAmoveByName(string) bool
	GetAmoveByName(string) IAmove
}

// AmoveHandler represents ...
type AmoveHandler struct {
	amoves []IAmove
	amove  IAmove
}

// GetAmove is ...
func (amoveh *AmoveHandler) GetAmove() IAmove {
	return amoveh.amove
}

// SetAmove is ...
func (amoveh *AmoveHandler) SetAmove(amove IAmove) bool {
	return amoveh.SetAmoveByName(amove.GetName())
}

// SetAmoveByName is ...
func (amoveh *AmoveHandler) SetAmoveByName(name string) bool {
	amove := amoveh.GetAmoveByName(name)
	if amove != nil {
		amoveh.amove = amove
		return true
	}
	return false
}

// GetAmoves is ...
func (amoveh *AmoveHandler) GetAmoves() []IAmove {
	return amoveh.amoves
}

// SetAmoves iss ...
func (amoveh *AmoveHandler) SetAmoves(amoves []IAmove) {
	amoveh.amoves = amoves
}

// AddAmove is ...
func (amoveh *AmoveHandler) AddAmove(amoves ...IAmove) bool {
	amoveh.amoves = append(amoveh.amoves, amoves...)
	return true
}

// RemoveAmove is ...
func (amoveh *AmoveHandler) RemoveAmove(amoves ...IAmove) bool {
	for _, amove := range amoves {
		if !amoveh.RemoveAmoveByName(amove.GetName()) {
			return false
		}
	}
	return true
}

// RemoveAmoveByName is ...
func (amoveh *AmoveHandler) RemoveAmoveByName(name string) bool {
	for i, amove := range amoveh.amoves {
		if amove.GetName() == name {
			amoveh.amoves = append(amoveh.amoves[:i], amoveh.amoves[i+1:]...)
			return true
		}
	}
	return false
}

// GetAmoveByName is ...
func (amoveh *AmoveHandler) GetAmoveByName(name string) IAmove {
	for _, amove := range amoveh.amoves {
		if amove.GetName() == name {
			return amove
		}
	}
	return nil
}

// NewAmoveHandler is ...
func NewAmoveHandler() *AmoveHandler {
	return &AmoveHandler{}
}
