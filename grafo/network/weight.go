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
	buffer.WriteString(fmt.Sprintf("%s-- %d -->%s\n", w.GetParent().GetLabel(), w.GetWeight(), w.GetChild().GetLabel()))
	return buffer.String()
}

// NewWeight is ...
func NewWeight(parent grafo.IVertex, child grafo.IVertex, w int) *Weight {
	return &Weight{
		Edge: grafo.NewEdge(parent,
			child,
			func(parent grafo.IVertex, child grafo.IVertex, params ...interface{}) (interface{}, bool) {
				return w, true
			}),
		weight: w,
	}
}
