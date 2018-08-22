package grafo

import (
	"bytes"
	"fmt"
)

// VtoV represents ...
type VtoV struct {
	id     Ider
	Parent IVertex
	Child  IVertex
}

// NewVtoV is ...
func NewVtoV(parent IVertex, child IVertex) *VtoV {
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
	return fmt.Sprintf("%s -> %s\n", edge.GetParent().GetLabel(), edge.GetChild().GetLabel())
}

// GetParent is ...
func (edge *Edge) GetParent() IVertex {
	return edge.Parent
}

// SetParent is ...
func (edge *Edge) SetParent(parent IVertex) {
	edge.Parent = parent
}

// GetChild is ...
func (edge *Edge) GetChild() IVertex {
	return edge.Child
}

// SetChild is ...
func (edge *Edge) SetChild(child IVertex) {
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
	buffer.WriteString(fmt.Sprintf("%s-->%s\n", edge.GetParent().GetLabel(), edge.GetChild().GetLabel()))
	return buffer.String()
}

// NewEdge is ...
func NewEdge(parent IVertex, child IVertex, clearance ClearanceCb) *Edge {
	return &Edge{
		VtoV:      NewVtoV(parent, child),
		Clearance: clearance,
	}
}

// StaticEdge is ...
func StaticEdge(parent IVertex, child IVertex) *Edge {
	return NewEdge(parent, child, func(parent IVertex, child IVertex, params ...interface{}) (interface{}, bool) {
		return nil, true
	})
}
