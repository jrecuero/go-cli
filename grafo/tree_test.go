package grafo_test

import (
	"testing"

	"github.com/jrecuero/go-cli/grafo"
)

// TestGrafo_NewGrafo is ...
func TestGrafo_NewGrafo(t *testing.T) {
	tree := grafo.NewGrafo("tree/1")
	if tree == nil {
		t.Errorf("NewGrafo: tree can not be <nil>")
	}
	if tree.Label != "tree/1" {
		t.Errorf("NewGrafo: label mistmatch: exp: %#v got: %#v\n", "tree/1", tree.Label)
	}
	if tree.GetRoot() == nil {
		t.Errorf("NewGrafo: root can not be <nil>")
	}
	if tree.GetAnchor() == nil {
		t.Errorf("NewGrafo: anchor can not be <nil>")
	}
	if tree.GetAnchor() != tree.GetRoot() {
		t.Errorf("NewGrafo: anchor mismatch: exp: %v got: %v\n", tree.GetRoot(), tree.GetAnchor())
	}
}

// TestGrafo_AddEdge is ...
func TestGrafo_AddEdge(t *testing.T) {
	tree := grafo.NewGrafo("tree/1")
	root := tree.GetRoot()
	parent := grafo.NewVertex("parent/1")
	rootEdge := grafo.StaticEdge(nil, parent)
	if err := tree.AddEdge(nil, rootEdge); err != nil {
		t.Errorf("AddEdge: return code error: %#v\n", err)
	}
	if rootEdge.Parent == nil {
		t.Errorf("Grafo:AddEdge: edge parent mismatch: exp not <nil>")
	}
	if rootEdge.Parent != root {
		t.Errorf("Grafo:AddEdge: edge parent label mismatch: exp: %v got: %v\n", root, rootEdge.Parent)
	}
	if len(parent.Parents) != 1 {
		t.Errorf("Grafo:AddEdge: child parents length mismatch: exp %d got: %d\n", 1, len(parent.Parents))
	}
	if parent.Parents[0] != root {
		t.Errorf("Grafo:AddEdge: child parent mismatch: exp %v got: %v\n", root, parent.Parents[0])
	}
}

// TestGrafo_AddChild is ...
func TestGrafo_AddChild(t *testing.T) {
	tree := grafo.NewGrafo("tree/1")
	root := tree.GetRoot()
	child1 := grafo.NewVertex("child/1")
	if err := tree.AddVertex(nil, child1); err != nil {
		t.Errorf("Grafo:AddVertex: return code error: %#v\n", err)
	}
	if len(child1.Parents) != 1 {
		t.Errorf("Grafo:AddVertex: child parents length mismatch: exp %d got: %d\n", 1, len(child1.Parents))
	}
	if child1.Parents[0] != root {
		t.Errorf("Grafo:AddVertex: child mismatch: exp %v got: %v\n", root, child1.Parents[0])
	}

	child2 := grafo.NewVertex("child/2")
	if err := tree.AddVertex(nil, child2); err != nil {
		t.Errorf("Grafo:AddVertex: return code error: %#v\n", err)
	}
	if len(child2.Parents) != 1 {
		t.Errorf("Grafo:AddVertex: child parents length mismatch: exp %d got: %d\n", 1, len(child2.Parents))
	}
	if child2.Parents[0] != root {
		t.Errorf("Grafo:AddVertex: child mismatch: exp %v got: %v\n", root, child2.Parents[0])
	}
	if len(root.Edges) != 2 {
		t.Errorf("Grafo:AdChild: root edges length mismatch: exp: %d got: %d\n", 2, len(root.Edges))
	}
}

// TestGrafo_ExistPathTo is ...
func TestGrafo_ExistPathTo(t *testing.T) {
	tree := grafo.NewGrafo("tree/1")
	parent := grafo.NewVertex("parent/1")
	child1 := grafo.NewVertex("child/1")
	child2 := grafo.NewVertex("child/2")
	child3 := grafo.NewVertex("child/3")
	tree.AddVertex(nil, parent)
	tree.AddVertex(parent, child1)
	tree.AddVertex(parent, child2)
	tree.AddVertex(child1, child3)
	if edge, ok := tree.ExistPathTo(parent, child1); ok {
		if edge != parent.Edges[0] {
			t.Errorf("Grafo:ExistPathTo: incorrect edge: exp: %#v got: %#v\n", parent.Edges[0], edge)
		}
	} else {
		t.Errorf("Grafo:ExistPathTo: edge not found from: %#v to %#v\n", parent, child1)
	}
	if _, ok := tree.ExistPathTo(parent, child3); ok {
		t.Errorf("Grafo:ExistPathTo: edge found from: %#v to %#v\n", parent, child3)
	}
}

// TestGrafo_IsPathTo is ...
func TestGrafo_IsPathTo(t *testing.T) {
	tree := grafo.NewGrafo("tree/1")
	parent := grafo.NewVertex("parent/1")
	child1 := grafo.NewVertex("child/1")
	child2 := grafo.NewVertex("child/2")
	child3 := grafo.NewVertex("child/3")
	tree.AddVertex(nil, parent)
	tree.AddVertex(parent, child1)
	tree.AddVertex(parent, child2)
	tree.AddVertex(child1, child3)
	if edge, ok := tree.IsPathTo(parent, child1); ok {
		if edge != parent.Edges[0] {
			t.Errorf("Grafo:IsPathTo: incorrect edge: exp: %#v got: %#v\n", parent.Edges[0], edge)
		}
	} else {
		t.Errorf("Grafo:IsPathTo: edge not found from: %#v to %#v\n", parent, child1)
	}
	if _, ok := tree.IsPathTo(parent, child3); ok {
		t.Errorf("Grafo:IsPathTo: edge found from: %#v to %#v\n", parent, child3)
	}
}

// TestGrafo_PathsFrom is ...
func TestGrafo_PathsFrom(t *testing.T) {
	tree := grafo.NewGrafo("tree/1")
	parent := grafo.NewVertex("parent/1")
	child1 := grafo.NewVertex("child/1")
	child2 := grafo.NewVertex("child/2")
	child3 := grafo.NewVertex("child/3")
	tree.AddVertex(nil, parent)
	tree.AddVertex(parent, child1)
	tree.AddVertex(parent, child2)
	tree.AddVertex(child1, child3)
	if paths := tree.PathsFrom(parent); len(paths) != 2 {
		t.Errorf("Grafo:PathsFrom: path length mismatch: exp: %d got: %d\n", 2, len(paths))
	}
	if paths := tree.PathsFrom(child3); len(paths) != 0 {
		t.Errorf("Grafo:PathsFrom: path length mismatch: exp: %d got: %d\n", 0, len(paths))
	}
}

// TestGrafo_AddVtoV
func TestGrafo_AddVtoV(t *testing.T) {
	tree := grafo.NewGrafo("tree/1")
	parent := grafo.NewVertex("parent/1")
	child1 := grafo.NewVertex("child/1")
	child2 := grafo.NewVertex("child/2")
	tree.AddVertex(nil, parent)
	tree.AddVertex(parent, child1)
	tree.AddVertex(parent, child2)
	if err := tree.AddVtoV(grafo.StaticEdge(nil, parent)); err != nil {
		t.Errorf("Grafo:AddVtoV: return error code: %#v\n", err)
	}
}

// TestGrafo_SetAnchorTo
func TestGrafo_SetAnchorTo(t *testing.T) {
	tree := grafo.NewGrafo("tree/1")
	parent := grafo.NewVertex("parent/1")
	child1 := grafo.NewVertex("child/1")
	child2 := grafo.NewVertex("child/2")
	tree.AddVertex(nil, parent)
	tree.AddVertex(parent, child1)
	tree.AddVertex(parent, child2)
	vertex := tree.SetAnchorTo(parent)
	if vertex == nil {
		t.Errorf("Grafo:SetAnchorTo: anchor cannot be <nil>")
	}
	if vertex != parent {
		t.Errorf("Grafo:SetAnchorTo: anchor mismatch: exp: %#v got: %#v\n", parent, vertex)
	}
}
