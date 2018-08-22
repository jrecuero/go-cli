package grafo

import (
	"bytes"
	"fmt"
)

// VtoV represents ...
type VtoV struct {
	id     Ider
	Parent *Vertex
	Child  *Vertex
}

// NewVtoV is ...
func NewVtoV(parent *Vertex, child *Vertex) *VtoV {
	return &VtoV{
		id:     nextIder(),
		Parent: parent,
		Child:  child,
	}
}

// Edge is ...
type Edge struct {
	*VtoV
	Clearance ClearanceCb
}

// String is ...
func (edge *Edge) String() string {
	return fmt.Sprintf("%s -> %s\n", edge.GetParent().Label, edge.GetChild().Label)
}

// GetParent is ...
func (edge *Edge) GetParent() *Vertex {
	return edge.Parent
}

// SetParent is ...
func (edge *Edge) SetParent(parent *Vertex) {
	edge.Parent = parent
}

// GetChild is ...
func (edge *Edge) GetChild() *Vertex {
	return edge.Child
}

// SetChild is ...
func (edge *Edge) SetChild(child *Vertex) {
	edge.Child = child
}

// GetVtoV is ...
func (edge *Edge) GetVtoV() *VtoV {
	return edge.VtoV
}

// Check is ...
func (edge *Edge) Check(params ...interface{}) (interface{}, bool) {
	return edge.Clearance(edge.GetParent(), edge.GetChild(), params...)
}

// ToMermaid is ...
func (edge *Edge) ToMermaid() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s-->%s\n", edge.GetParent().Label, edge.GetChild().Label))
	return buffer.String()
}

// NewEdge is ...
func NewEdge(parent *Vertex, child *Vertex, clearance ClearanceCb) *Edge {
	return &Edge{
		VtoV:      NewVtoV(parent, child),
		Clearance: clearance,
	}
}

// StaticEdge is ...
func StaticEdge(parent *Vertex, child *Vertex) *Edge {
	return NewEdge(parent, child, func(parent *Vertex, child *Vertex, params ...interface{}) (interface{}, bool) {
		return nil, true
	})
}
