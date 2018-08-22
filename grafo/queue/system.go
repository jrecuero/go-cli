package queue

import "github.com/jrecuero/go-cli/grafo"

// System represents ..
type System struct {
	*grafo.Grafo
}

// AddQueue is ...
func (system *System) AddQueue(parent *grafo.Vertex, child *grafo.Vertex, limit int) error {
	if parent == nil {
		parent = system.GetRoot()
	}
	var edge grafo.IEdge = NewQueue(parent, child, limit)
	return system.AddEdge(parent, edge)
}

// NewSystem is ...
func NewSystem(label string) *System {
	return &System{
		grafo.NewGrafo(label),
	}
}
