package battle

// IStance represents ...
type IStance interface {
	IBase
	GetStyle() IStyle
}

// Stance represents ...
type Stance struct {
	*Base
	parent      IStyle
	updatestats *UStats
}

// NewStance is ...
func NewStance(name string, parent IStyle) *Stance {
	return &Stance{
		Base:   NewBase(name),
		parent: parent,
	}
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
