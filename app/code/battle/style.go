package battle

// IStyle represents ...
type IStyle interface {
	IBase
	GetTechnique() ITechnique
	GetStances() []IStance
	AddStance(...IStance) bool
	RemoveStance(...IStance) bool
	RemoveStanceByName(string) bool
	GetStanceByName(string) IStance
}

// Style represents ...
type Style struct {
	*Base
	stances []IStance
	parent  ITechnique
}

// GetStances is ...
func (style *Style) GetStances() []IStance {
	return style.stances
}

// AddStance is ...
func (style *Style) AddStance(stances ...IStance) bool {
	style.stances = append(style.stances, stances...)
	return true
}

// RemoveStance is ...
func (style *Style) RemoveStance(stances ...IStance) bool {
	for _, stance := range stances {
		if !style.RemoveStanceByName(stance.GetName()) {
			return false
		}
	}
	return true
}

// RemoveStanceByName is ...
func (style *Style) RemoveStanceByName(name string) bool {
	for i, stance := range style.stances {
		if stance.GetName() == name {
			style.stances = append(style.stances[:i], style.stances[i+1:]...)
			return true
		}
	}
	return false
}

// GetStanceByName is ...
func (style *Style) GetStanceByName(name string) IStance {
	for _, stance := range style.stances {
		if stance.GetName() == name {
			return stance
		}
	}
	return nil
}

// NewStyle is ...
func NewStyle(name string, parent ITechnique) *Style {
	return &Style{
		Base:   NewBase(name),
		parent: parent,
	}
}

// IStyleHandler represents ...
type IStyleHandler interface {
	GetStyle() IStyle
	SetStyle(IStyle) bool
	SetStyleByName(string) bool
	GetStyles() []IStyle
	SetStyles([]IStyle)
	AddStyle(...IStyle) bool
	RemoveStyle(...IStyle) bool
	RemoveStyleByName(string) bool
	GetStyleByName(string) IStyle
}

// StyleHandler represents ...
type StyleHandler struct {
	styles []IStyle
	style  IStyle
}

// GetStyle is ...
func (sh *StyleHandler) GetStyle() IStyle {
	return sh.style
}

// SetStyle is ...
func (sh *StyleHandler) SetStyle(style IStyle) bool {
	return sh.SetStyleByName(style.GetName())
}

// SetStyleByName is ...
func (sh *StyleHandler) SetStyleByName(name string) bool {
	style := sh.GetStyleByName(name)
	if style != nil {
		sh.style = style
		return true
	}
	return false
}

// GetStyles is ...
func (sh *StyleHandler) GetStyles() []IStyle {
	return sh.styles
}

// SetStyles iss ...
func (sh *StyleHandler) SetStyles(styles []IStyle) {
	sh.styles = styles
}

// AddStyle is ...
func (sh *StyleHandler) AddStyle(styles ...IStyle) bool {
	sh.styles = append(sh.styles, styles...)
	return true
}

// RemoveStyle is ...
func (sh *StyleHandler) RemoveStyle(styles ...IStyle) bool {
	for _, style := range styles {
		if !sh.RemoveStyleByName(style.GetName()) {
			return false
		}
	}
	return true
}

// RemoveStyleByName is ...
func (sh *StyleHandler) RemoveStyleByName(name string) bool {
	for index, style := range sh.styles {
		if style.GetName() == name {
			sh.styles = append(sh.styles[:index], sh.styles[index+1:]...)
			return true
		}
	}
	return false
}

// GetStyleByName is ...
func (sh *StyleHandler) GetStyleByName(name string) IStyle {
	for _, style := range sh.styles {
		if style.GetName() == name {
			return style
		}
	}
	return nil
}

// NewStyleHandler is ...
func NewStyleHandler() *StyleHandler {
	return &StyleHandler{}
}
