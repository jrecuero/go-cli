package grafo

import (
	"bytes"
	"fmt"
)

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

//NewNetwork is ...
func NewNetwork(label string) *Network {
	return &Network{
		NewTree(label),
	}
}
