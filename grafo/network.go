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
				p := NewPath(fmt.Sprintf("%s : %s", branch.GetParent().Label, branch.GetChild().Label))
				p.Traversed = append(p.Traversed, branch.GetTraverse())
				paths = append(paths, p)
			} else {
				if childPaths := net.pathsFromNodeToNode(ToNode(branch.GetChild()), dest, ids); childPaths != nil {
					for _, childPath := range childPaths {
						traversed := []*Traverse{branch.GetTraverse()}
						childPath.Traversed = append(traversed, childPath.Traversed...)
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
