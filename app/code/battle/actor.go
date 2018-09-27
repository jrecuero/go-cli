package battle

// IActor represents
type IActor interface {
	ITechniqueHandler
	IStyleHandler
	IStanceHandler
	GetName() string
	GetStats() *Stats
}

// Actor represents ...
type Actor struct {
	*TechniqueHandler
	*StyleHandler
	*StanceHandler
	name  string
	stats *Stats
}

// NewActor is ...
func NewActor(name string) *Actor {
	return &Actor{
		TechniqueHandler: NewTechniqueHandler(),
		StyleHandler:     NewStyleHandler(),
		StanceHandler:    NewStanceHandler(),
		name:             name,
		stats:            NewStats(),
	}
}
