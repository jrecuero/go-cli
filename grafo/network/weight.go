package network

import (
	"bytes"
	"fmt"

	"github.com/jrecuero/go-cli/grafo"
)

// Weight represents ...
type Weight struct {
	*grafo.Edge
	weight int
}

// GetWeight is ...
func (w *Weight) GetWeight() int {
	return w.weight
}

// ToMermaid is ...
func (w *Weight) ToMermaid() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s-- %d -->%s\n", w.GetParent().Label, w.GetWeight(), w.GetChild().Label))
	return buffer.String()
}

// NewWeight is ...
func NewWeight(parent *grafo.Vertex, child *grafo.Vertex, w int) *Weight {
	return &Weight{
		Edge: grafo.NewEdge(parent,
			child,
			func(parent *grafo.Vertex, child *grafo.Vertex, params ...interface{}) (interface{}, bool) {
				return w, true
			}),
		weight: w,
	}
}
