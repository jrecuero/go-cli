package grafo

import (
	"bytes"
	"fmt"
)

// NodeContent represents ...
type NodeContent struct {
	label  string
	weight int
}

// GetLabel is ...
func (nc *NodeContent) GetLabel() string {
	return nc.label
}

// GetWeight is ...
func (nc *NodeContent) GetWeight() int {
	return nc.weight
}

// NewNodeContent is ...
func NewNodeContent(label string, weight int) *NodeContent {
	return &NodeContent{
		label:  label,
		weight: weight,
	}
}

// Node represents ...
type Node struct {
	*Vertex
}

// String is ...
func (node *Node) String() string {
	return node.Label
}

// NewNode is ...
func NewNode(label string, nc *NodeContent) *Node {
	node := &Node{
		Vertex: NewVertex(label),
	}
	node.Content = nc
	return node
}

// Weight represents ...
type Weight struct {
	*Edge
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
func NewWeight(parent *Vertex, child *Vertex, w int) *Weight {
	return &Weight{
		Edge: NewEdge(parent,
			child,
			func(parent *Vertex, child *Vertex, params ...interface{}) (interface{}, bool) {
				return w, true
			}),
		weight: w,
	}
}

// Network represents ...
type Network struct {
	*Grafo
}

// AddNode is ...
func (net *Network) AddNode(parent *Vertex, child *Vertex, weight int) error {
	if parent == nil {
		parent = net.GetRoot()
	}
	var edge IEdge = NewWeight(parent, child, weight)
	return net.AddEdge(parent, edge)
}

// CostToNode returns how much weight is required to move fromt the network
// anchor to the destination node.
func (net *Network) CostToNode(dest *Node) (int, bool) {
	if edge, ok := net.ExistPathTo(nil, dest.Vertex); ok {
		if w, bok := edge.Check(); bok {
			weight := w.(int) + dest.Content.(*NodeContent).GetWeight()
			return weight, true
		}
	}
	return 0, false
}

// pathsFromNodeToNode is ...
func (net *Network) pathsFromNodeToNode(anchor *Node, dest *Node, ids []Ider) []*Path {
	var paths []*Path
	ids = append(ids, anchor.GetID())
	//tools.ToDisplay("%#v\n", ids)
	for _, edge := range anchor.Edges {
		if found := findIDInArray(edge.GetChild().GetID(), ids); !found {
			if edge.GetChild() == NodeToVertex(dest) {
				p := NewPath("")
				p.Edges = append(p.Edges, edge)
				paths = append(paths, p)
			} else {
				if childPaths := net.pathsFromNodeToNode(ToNode(edge.GetChild()), dest, ids); childPaths != nil {
					for _, childPath := range childPaths {
						edgees := []IEdge{edge}
						childPath.Edges = append(edgees, childPath.Edges...)
						paths = append(paths, childPath)
					}
				}
			}
		}
	}
	return paths
}

// PathsFromNodeToNode is ...
func (net *Network) PathsFromNodeToNode(anchor *Node, dest *Node) []*Path {
	return net.pathsFromNodeToNode(anchor, dest, nil)
}

// FindLoops is ...
func (net *Network) FindLoops(anchor *Node, ids []Ider) [][]*Node {
	if index := indexOf(anchor.GetID(), ids); index != -1 {
		//tools.ToDisplay("Loop at: %#v : %d\n", ids, index)
		var nodes []*Node
		for _, id := range ids[index:len(ids)] {
			nodes = append(nodes, ToNode(net.GetVertices()[id]))
		}
		return [][]*Node{nodes}
	}
	ids = append(ids, anchor.GetID())
	var loops [][]*Node
	for _, edge := range anchor.Edges {
		if loop := net.FindLoops(ToNode(edge.GetChild()), ids); loop != nil {
			for _, l := range loop {
				loops = append(loops, l)
			}
		}
	}
	return loops
}

// TotalWeightInPath is ...
func (net *Network) TotalWeightInPath(path *Path) int {
	var weight int
	for _, edge := range path.Edges {
		if w, ok := edge.Check(); ok {
			weight += w.(int)
		}
		weight += edge.GetChild().Content.(*NodeContent).GetWeight()
	}
	return weight
}

// isBetterPath is ...
func (net *Network) isBetterPath(best *Path, bestWeight int, path *Path) (*Path, int) {
	pathWeight := net.TotalWeightInPath(path)
	if best == nil || pathWeight < bestWeight {
		return path, pathWeight
	}
	return best, bestWeight
}

// BestPathFromNodeToNode is ...
func (net *Network) BestPathFromNodeToNode(anchor *Node, dest *Node) (*Path, int) {
	var bestPath *Path
	var bestWeight int
	paths := net.PathsFromNodeToNode(anchor, dest)
	for _, path := range paths {
		bestPath, bestWeight = net.isBetterPath(bestPath, bestWeight, path)
	}
	return bestPath, bestWeight
}

//NewNetwork is ...
func NewNetwork(label string) *Network {
	return &Network{
		Grafo: NewGrafo(label),
	}
}

// NodeToVertex is ...
func NodeToVertex(node *Node) *Vertex {
	return node.Vertex
}

// ToNode is ...
func ToNode(vertex *Vertex) *Node {
	return &Node{
		vertex,
	}
}

// findIDInArray is ...
func findIDInArray(id Ider, lista []Ider) bool {
	for _, val := range lista {
		if val == id {
			return true
		}
	}
	return false
}

// indexOf is ...
func indexOf(id Ider, lista []Ider) int {
	for index, val := range lista {
		if val == id {
			return index
		}
	}
	return -1
}
