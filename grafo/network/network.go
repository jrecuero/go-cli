package network

import (
	"github.com/jrecuero/go-cli/grafo"
)

// Network represents ...
type Network struct {
	*grafo.Grafo
}

// AddNode is ...
func (net *Network) AddNode(parent *grafo.Vertex, child *grafo.Vertex, weight int) error {
	if parent == nil {
		parent = net.GetRoot()
	}
	var edge grafo.IEdge = NewWeight(parent, child, weight)
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
func (net *Network) pathsFromNodeToNode(anchor *Node, dest *Node, ids []grafo.Ider) []*grafo.Path {
	var paths []*grafo.Path
	ids = append(ids, anchor.GetID())
	//tools.ToDisplay("%#v\n", ids)
	for _, edge := range anchor.Edges {
		if found := findIDInArray(edge.GetChild().GetID(), ids); !found {
			if edge.GetChild() == NodeToVertex(dest) {
				p := grafo.NewPath("")
				p.Edges = append(p.Edges, edge)
				paths = append(paths, p)
			} else {
				if childPaths := net.pathsFromNodeToNode(ToNode(edge.GetChild()), dest, ids); childPaths != nil {
					for _, childPath := range childPaths {
						edgees := []grafo.IEdge{edge}
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
func (net *Network) PathsFromNodeToNode(anchor *Node, dest *Node) []*grafo.Path {
	return net.pathsFromNodeToNode(anchor, dest, nil)
}

// FindLoops is ...
func (net *Network) FindLoops(anchor *Node, ids []grafo.Ider) [][]*Node {
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
func (net *Network) TotalWeightInPath(path *grafo.Path) int {
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
func (net *Network) isBetterPath(best *grafo.Path, bestWeight int, path *grafo.Path) (*grafo.Path, int) {
	pathWeight := net.TotalWeightInPath(path)
	if best == nil || pathWeight < bestWeight {
		return path, pathWeight
	}
	return best, bestWeight
}

// BestPathFromNodeToNode is ...
func (net *Network) BestPathFromNodeToNode(anchor *Node, dest *Node) (*grafo.Path, int) {
	var bestPath *grafo.Path
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
		Grafo: grafo.NewGrafo(label),
	}
}

// findIDInArray is ...
func findIDInArray(id grafo.Ider, lista []grafo.Ider) bool {
	for _, val := range lista {
		if val == id {
			return true
		}
	}
	return false
}

// indexOf is ...
func indexOf(id grafo.Ider, lista []grafo.Ider) int {
	for index, val := range lista {
		if val == id {
			return index
		}
	}
	return -1
}
