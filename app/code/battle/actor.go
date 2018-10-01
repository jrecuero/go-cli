package battle

// IActor represents
type IActor interface {
	ITechniqueHandler
	IStyleHandler
	IStanceHandler
	IAmoveHandler
	GetName() string
	GetStats() *Stats
}

// Actor represents ...
type Actor struct {
	*TechniqueHandler
	*StyleHandler
	*StanceHandler
	*AmoveHandler
	name  string
	stats *Stats
}

// GetName is ...
func (actor *Actor) GetName() string {
	return actor.name
}

// GetStats is ...
func (actor *Actor) GetStats() *Stats {
	return actor.stats
}

// NewActor is ...
func NewActor(name string) *Actor {
	return &Actor{
		TechniqueHandler: NewTechniqueHandler(),
		StyleHandler:     NewStyleHandler(),
		StanceHandler:    NewStanceHandler(),
		AmoveHandler:     NewAmoveHandler(),
		name:             name,
		stats:            NewStats(),
	}
}
