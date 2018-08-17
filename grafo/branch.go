package grafo

import (
	"bytes"
	"fmt"
)

// Traverse represents ...
type Traverse struct {
	id     Ider
	Parent *Leaf
	Child  *Leaf
}

// NewTraverse is ...
func NewTraverse(parent *Leaf, child *Leaf) *Traverse {
	return &Traverse{
		id:     nextIder(),
		Parent: parent,
		Child:  child,
	}
}

// Branch is ...
type Branch struct {
	*Traverse
	clearance ClearanceCb
}

// GetParent is ...
func (branch *Branch) GetParent() *Leaf {
	return branch.Parent
}

// SetParent is ...
func (branch *Branch) SetParent(parent *Leaf) {
	branch.Parent = parent
}

// GetChild is ...
func (branch *Branch) GetChild() *Leaf {
	return branch.Child
}

// SetChild is ...
func (branch *Branch) SetChild(child *Leaf) {
	branch.Child = child
}

// GetTraverse is ...
func (branch *Branch) GetTraverse() *Traverse {
	return branch.Traverse
}

// Check is ...
func (branch *Branch) Check(params ...interface{}) (interface{}, bool) {
	return branch.clearance(branch.GetParent(), branch.GetChild(), params...)
}

// ToMermaid is ...
func (branch *Branch) ToMermaid() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s-->%s\n", branch.GetParent().Label, branch.GetChild().Label))
	return buffer.String()
}

// NewBranch is ...
func NewBranch(parent *Leaf, child *Leaf, clearance ClearanceCb) *Branch {
	return &Branch{
		Traverse:  NewTraverse(parent, child),
		clearance: clearance,
	}
}

// StaticBranch is ...
func StaticBranch(parent *Leaf, child *Leaf) *Branch {
	return NewBranch(parent, child, func(parent *Leaf, child *Leaf, params ...interface{}) (interface{}, bool) {
		return nil, true
	})
}
