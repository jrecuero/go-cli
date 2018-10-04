package battle

import (
	"bytes"
	"fmt"
)

// ITechnique represents ...
type ITechnique interface {
	IBase
	GetStyles() []IStyle
	AddStyle(...IStyle) bool
	RemoveStyle(...IStyle) bool
	RemoveStyleByName(string) bool
	GetStyleByName(string) IStyle
}

// Technique represents ...
type Technique struct {
	*Base
	styles []IStyle
}

// GetStyles is ...
func (tech *Technique) GetStyles() []IStyle {
	return tech.styles
}

// AddStyle is ...
func (tech *Technique) AddStyle(styles ...IStyle) bool {
	tech.styles = append(tech.styles, styles...)
	return true
}

// RemoveStyle is ...
func (tech *Technique) RemoveStyle(styles ...IStyle) bool {
	for _, style := range styles {
		if !tech.RemoveStyleByName(style.GetName()) {
			return false
		}
	}
	return true
}

// RemoveStyleByName is ...
func (tech *Technique) RemoveStyleByName(name string) bool {
	for i, style := range tech.styles {
		if style.GetName() == name {
			tech.styles = append(tech.styles[:i], tech.styles[i+1:]...)
			return true
		}
	}
	return false
}

// GetStyleByName is ...
func (tech *Technique) GetStyleByName(name string) IStyle {
	for _, style := range tech.styles {
		if style.GetName() == name {
			return style
		}
	}
	return nil
}

// String is ...
func (tech *Technique) String() string {
	var buf bytes.Buffer
	//buf.WriteString(fmt.Sprintf("> Technique Name: %s", tech.GetName()))
	buf.WriteString(fmt.Sprintf("> Technique %s", tech.Base))
	for _, style := range tech.styles {
		buf.WriteString(fmt.Sprintf("\n  : %s", style))
	}
	return buf.String()
}

// NewTechnique is ...
func NewTechnique(name string) *Technique {
	return &Technique{
		Base: NewBase(name),
	}
}

// NewFullTechnique is  ...
func NewFullTechnique(name string, desc string, ustats *UStats) *Technique {
	return &Technique{
		Base: NewFullBase(name, desc, ustats),
	}
}

// ITechniqueHandler represents ...
type ITechniqueHandler interface {
	GetTechnique() ITechnique
	SetTechnique(ITechnique) bool
	SetTechniqueByName(string) bool
	GetTechniques() []ITechnique
	SetTechniques([]ITechnique)
	AddTechnique(...ITechnique) bool
	RemoveTechnique(...ITechnique) bool
	RemoveTechniqueByName(string) bool
	GetTechniqueByName(string) ITechnique
}

// TechniqueHandler represents ...
type TechniqueHandler struct {
	techniques []ITechnique
	technique  ITechnique
}

// GetTechnique is ...
func (th *TechniqueHandler) GetTechnique() ITechnique {
	return th.technique
}

// SetTechnique is ...
func (th *TechniqueHandler) SetTechnique(tech ITechnique) bool {
	return th.SetTechniqueByName(tech.GetName())
}

// SetTechniqueByName is ...
func (th *TechniqueHandler) SetTechniqueByName(name string) bool {
	tech := th.GetTechniqueByName(name)
	if tech != nil {
		th.technique = tech
		return true
	}
	return false
}

// GetTechniques is ...
func (th *TechniqueHandler) GetTechniques() []ITechnique {
	return th.techniques
}

// SetTechniques iss ...
func (th *TechniqueHandler) SetTechniques(techs []ITechnique) {
	th.techniques = techs
}

// AddTechnique is ...
func (th *TechniqueHandler) AddTechnique(techs ...ITechnique) bool {
	th.techniques = append(th.techniques, techs...)
	return true
}

// RemoveTechnique is ...
func (th *TechniqueHandler) RemoveTechnique(techs ...ITechnique) bool {
	for _, tech := range techs {
		if !th.RemoveTechniqueByName(tech.GetName()) {
			return false
		}
	}
	return true
}

// RemoveTechniqueByName is ...
func (th *TechniqueHandler) RemoveTechniqueByName(name string) bool {
	for i, tech := range th.techniques {
		if tech.GetName() == name {
			th.techniques = append(th.techniques[:i], th.techniques[i+1:]...)
			return true
		}
	}
	return false
}

// GetTechniqueByName is ...
func (th *TechniqueHandler) GetTechniqueByName(name string) ITechnique {
	for _, tech := range th.techniques {
		if tech.GetName() == name {
			return tech
		}
	}
	return nil
}

// NewTechniqueHandler is ...
func NewTechniqueHandler() *TechniqueHandler {
	return &TechniqueHandler{}
}
