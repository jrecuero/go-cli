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
