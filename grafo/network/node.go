package network

import "github.com/jrecuero/go-cli/grafo"

// Node represents ...
type Node struct {
	*grafo.Vertex
}

// String is ...
func (node *Node) String() string {
	return node.Label
}

// NewNode is ...
func NewNode(label string, nc *NodeContent) *Node {
	node := &Node{
		Vertex: grafo.NewVertex(label),
	}
	node.Content = nc
	return node
}

// NodeToVertex is ...
func NodeToVertex(node *Node) *grafo.Vertex {
	return node.Vertex
}

// ToNode is ...
func ToNode(vertex *grafo.Vertex) *Node {
	return &Node{
		vertex,
	}
}
