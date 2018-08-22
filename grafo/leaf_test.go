package grafo_test

import (
	"testing"

	"github.com/jrecuero/go-cli/grafo"
)

// TestVertex_NewVertex is ...
func TestVertex_NewVertex(t *testing.T) {
	if vertex := grafo.NewVertex("vertex/0"); vertex == nil {
		t.Errorf("NewVertex: vertex can not be <nil>")
	}
}

// TestVertex_AddEdge is ...
func TestVertex_AddEdge(t *testing.T) {
	parentVertex := grafo.NewVertex("root/1")
	childVertex := grafo.NewVertex("child/1")
	edge := grafo.StaticEdge(parentVertex, childVertex)
	if err := parentVertex.AddEdge(edge); err != nil {
		t.Errorf("Vertex:AddEdge: return code error: %#v\n", err)
	}
	if len(parentVertex.Edges) != 1 {
		t.Errorf("Vertex:AddEdge: edges length mismatch: exp %d got: %d\n", 1, len(parentVertex.Edges))
	}
}
