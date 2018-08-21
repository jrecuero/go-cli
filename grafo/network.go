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
	*Leaf
}

// NewNode is ...
func NewNode(label string, nc *NodeContent) *Node {
	node := &Node{
		Leaf: NewLeaf(label),
	}
	node.Content = nc
	return node
}

// Weight represents ...
type Weight struct {
	*Branch
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
func NewWeight(parent *Leaf, child *Leaf, w int) *Weight {
	return &Weight{
		Branch: NewBranch(parent,
			child,
			func(parent *Leaf, child *Leaf, params ...interface{}) (interface{}, bool) {
				return w, true
			}),
		weight: w,
	}
}

// Network represents ...
type Network struct {
	*Tree
}

// AddNode is ...
func (net *Network) AddNode(parent *Leaf, child *Leaf, weight int) error {
	if parent == nil {
		parent = net.GetRoot()
	}
	var branch IBranch = NewWeight(parent, child, weight)
	return net.AddBranch(parent, branch)
}

// CostToNode returns how much weight is required to move fromt the network
// anchor to the destination node.
func (net *Network) CostToNode(dest *Node) (int, bool) {
	if branch, ok := net.ExistPathTo(nil, dest.Leaf); ok {
		if w, bok := branch.Check(); bok {
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
	for _, branch := range anchor.Branches {
		if found := findIDInArray(branch.GetChild().GetID(), ids); !found {
			if branch.GetChild() == ToLeaf(dest) {
				p := NewPath("")
				p.Branches = append(p.Branches, branch)
				paths = append(paths, p)
			} else {
				if childPaths := net.pathsFromNodeToNode(ToNode(branch.GetChild()), dest, ids); childPaths != nil {
					for _, childPath := range childPaths {
						branches := []IBranch{branch}
						childPath.Branches = append(branches, childPath.Branches...)
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

// TotalWeightInPath is ...
func (net *Network) TotalWeightInPath(path *Path) int {
	var weight int
	for _, branch := range path.Branches {
		if w, ok := branch.Check(); ok {
			weight += w.(int)
		}
		weight += branch.GetChild().Content.(*NodeContent).GetWeight()
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
		Tree: NewTree(label),
	}
}

// ToLeaf is ...
func ToLeaf(node *Node) *Leaf {
	return node.Leaf
}

// ToNode is ...
func ToNode(leaf *Leaf) *Node {
	return &Node{
		leaf,
	}
}

// findIDInArray is ...
func findIDInArray(id Ider, lista []Ider) bool {
	for _, i := range lista {
		if i == id {
			return true
		}
	}
	return false
}
