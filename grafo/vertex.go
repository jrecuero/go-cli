package grafo

import (
	"bytes"
	"fmt"
	"strconv"
)

// IVertex represents ...
type IVertex interface {
	GetID() Ider
	GetLabel() string
	GetParents() []IVertex
	GetEdges() []IEdge
	GetTraversed() []*VtoV
	GetContent() IContent
	AddParent(IVertex) error
	AddEdge(IEdge) error
}

// Vertex represents ...
type Vertex struct {
	id        Ider
	Label     string
	Parents   []*Vertex
	Edges     []IEdge
	Traversed []*VtoV
	Content   IContent
	hooked    bool
}

// GetID is ...
func (vertex *Vertex) GetID() Ider {
	return vertex.id
}

// AddParent is ...
func (vertex *Vertex) AddParent(parent *Vertex) error {
	vertex.Parents = append(vertex.Parents, parent)
	return nil
}

// AddEdge is ...
func (vertex *Vertex) AddEdge(edge IEdge) error {
	vertex.Edges = append(vertex.Edges, edge)
	return nil
}

// String is ...
func (vertex *Vertex) String() string {
	return vertex.Label
}

// NewVertex is ...
func NewVertex(label string) *Vertex {
	return &Vertex{
		id:    nextIder(),
		Label: label,
	}
}

// Path represents ...
type Path struct {
	id    Ider
	Label string
	Edges []IEdge
}

// String is ...
func (path *Path) String() string {
	var buffer bytes.Buffer
	if len(path.Edges) != 0 {
		for _, b := range path.Edges {
			buffer.WriteString(fmt.Sprintf("%s : ", b.GetParent().Label))
		}
		buffer.WriteString(fmt.Sprintf("%s", path.Edges[len(path.Edges)-1].GetChild().Label))
	}
	return buffer.String()
}

// NewPath is ...
func NewPath(label string) *Path {
	id := nextIder()
	if label == "" {
		label = strconv.Itoa(int(id))
	}
	return &Path{
		id:    id,
		Label: label,
	}
}
