package grafo

import (
	"bytes"
	"fmt"

	"github.com/jrecuero/go-cli/tools"
)

// Grafo represents ...
type Grafo struct {
	id       Ider
	Label    string
	root     IVertex
	anchor   IVertex
	path     *Path
	vertices map[Ider]IVertex
}

// GetRoot is ...
func (grafo *Grafo) GetRoot() IVertex {
	return grafo.root
}

// GetAnchor is ...
func (grafo *Grafo) GetAnchor() IVertex {
	return grafo.anchor
}

// GetVertices is ...
func (grafo *Grafo) GetVertices() map[Ider]IVertex {
	return grafo.vertices
}

// AddEdge adds the given edge to the given parent. If parent is nil, use
// the grafo root vertex. Parent attribute in the Child vertex is set properly.
func (grafo *Grafo) AddEdge(parent IVertex, edge IEdge) error {
	if parent == nil {
		parent = grafo.GetRoot()
		edge.SetParent(parent)
	}
	if !parent.GetHooked() {
		return tools.ERROR(nil, false, "Parent not found in grafo: %#v\n", parent)
	}
	if err := parent.AddEdge(edge); err != nil {
		return err
	}
	child := edge.GetChild()
	if err := child.AddParent(parent); err != nil {
		return err
	}
	child.SetHooked(true)
	grafo.vertices[child.GetID()] = child
	return nil
}

// AddVertex adds an static edge fromt the given parent to the given child.
func (grafo *Grafo) AddVertex(parent IVertex, child IVertex) error {
	if parent == nil {
		parent = grafo.GetRoot()
	}
	var edge IEdge = StaticEdge(parent, child)
	return grafo.AddEdge(parent, edge)
}

// ExistPathTo checks if there is edge from the given anchor to the given
// child. If not anchor vertex is provided, the grafo anchor is used instead of.
func (grafo *Grafo) ExistPathTo(anchor IVertex, dest IVertex) (IEdge, bool) {
	if anchor == nil {
		anchor = grafo.anchor
	}
	for _, edge := range anchor.GetEdges() {
		if edge.GetChild() == dest {
			return edge, true
		}
	}
	return nil, false
}

// IsPathTo check if there is a edge from the given anchor to the given child
// ana if the path is possible. If not anchor is vertex is provided, the grafo
// anchor is used instead of.
func (grafo *Grafo) IsPathTo(anchor IVertex, dest IVertex, params ...interface{}) (IEdge, bool) {
	if anchor == nil {
		anchor = grafo.anchor
	}
	if edge, ok := grafo.ExistPathTo(anchor, dest); ok {
		if _, bok := edge.Check(params...); bok {
			return edge, true
		}
	}
	return nil, false
}

// PathsFrom returns all existance and possible Edges from the given anchor.
func (grafo *Grafo) PathsFrom(anchor IVertex, params ...interface{}) []IVertex {
	var children []IVertex
	if anchor == nil {
		anchor = grafo.anchor
	}
	for _, edge := range anchor.GetEdges() {
		if _, ok := edge.Check(params...); ok {
			children = append(children, edge.GetChild())
		}
	}
	return children
}

// setAnchor is ..
func (grafo *Grafo) setAnchor(anchor IVertex) IVertex {
	grafo.anchor = anchor
	return grafo.GetAnchor()
}

// AddVtoV adds a edge to the grafo traverse.
func (grafo *Grafo) AddVtoV(edge IEdge) error {
	if edge.GetParent() == nil {
		edge.SetParent(grafo.GetRoot())
	}
	if grafo.GetAnchor() != edge.GetParent() {
		return tools.ERROR(nil, false, "parent is not the anchor: %#v\n", edge.GetParent())
	}
	grafo.setAnchor(edge.GetChild())
	grafo.path.Edges = append(grafo.path.Edges, edge)
	return nil
}

// SetAnchorTo moves the anchor to the destination vertex and adds the edge to
// the grafo traverse.
func (grafo *Grafo) SetAnchorTo(dest IVertex) IVertex {
	for _, edge := range grafo.anchor.GetEdges() {
		if edge.GetChild() == dest {
			if err := grafo.AddVtoV(edge); err != nil {
				return nil
			}
			return grafo.GetAnchor()
		}
	}
	return nil
}

// ToMermaid is ...
func (grafo *Grafo) ToMermaid() string {
	var buffer bytes.Buffer
	buffer.WriteString("graph LR\n")
	for _, vertex := range grafo.GetVertices() {
		for _, edge := range vertex.GetEdges() {
			buffer.WriteString(edge.ToMermaid())
		}
	}
	return buffer.String()
}

// NewGrafo is ...
func NewGrafo(label string) *Grafo {
	root := NewVertex("root/0")
	root.SetHooked(true)
	grafo := &Grafo{
		id:       nextIder(),
		Label:    label,
		root:     root,
		path:     NewPath(fmt.Sprintf("%s/path", label)),
		vertices: make(map[Ider]IVertex),
	}
	grafo.anchor = grafo.GetRoot()
	return grafo
}
