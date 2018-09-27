package battle

// IActor represents
type IActor interface {
	ITechniqueHandler
	GetName() string
	GetStats() *Stats
	GetStyle() IStyle
	GetStances() IStance
	Styles() []IStyle
	Stances() []IStance
}

// Actor represents ...
type Actor struct {
	*TechniqueHandler
	name    string
	stats   *Stats
	styles  []IStyle
	stances []IStance
	style   IStyle
	stance  IStance
}

// NewActor is ...
func NewActor(name string) *Actor {
	return &Actor{
		TechniqueHandler: NewTechniqueHandler(),
		name:             name,
		stats:            NewStats(),
	}
}
