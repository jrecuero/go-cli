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
	//SetEdges([]IEdge)
	GetTraversed() []*VtoV
	GetContent() IContent
	GetHooked() bool
	SetHooked(bool)
	AddParent(IVertex) error
	AddEdge(IEdge) error
}

// Vertex represents ...
type Vertex struct {
	id        Ider
	Label     string
	Parents   []IVertex
	Edges     []IEdge
	Traversed []*VtoV
	Content   IContent
	hooked    bool
}

// GetID is ...
func (vertex *Vertex) GetID() Ider {
	return vertex.id
}

// GetLabel is ...
func (vertex *Vertex) GetLabel() string {
	return vertex.Label
}

// GetParents is ...
func (vertex *Vertex) GetParents() []IVertex {
	return vertex.Parents
}

// GetEdges is ...
func (vertex *Vertex) GetEdges() []IEdge {
	return vertex.Edges
}

//// SetEdges is ...
//func (vertex *Vertex) SetEdges(edges []IEdge) {
//    vertex.Edges = edges
//}

// GetTraversed is ...
func (vertex *Vertex) GetTraversed() []*VtoV {
	return vertex.Traversed
}

// GetContent is ...
func (vertex *Vertex) GetContent() IContent {
	return vertex.Content
}

// GetHooked is ..
func (vertex *Vertex) GetHooked() bool {
	return vertex.hooked
}

// SetHooked is ...
func (vertex *Vertex) SetHooked(hooked bool) {
	vertex.hooked = hooked
}

// AddParent is ...
func (vertex *Vertex) AddParent(parent IVertex) error {
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
			buffer.WriteString(fmt.Sprintf("%s : ", b.GetParent().GetLabel()))
		}
		buffer.WriteString(fmt.Sprintf("%s", path.Edges[len(path.Edges)-1].GetChild().GetLabel()))
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
