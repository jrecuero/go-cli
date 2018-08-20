package grafo

import (
	"bytes"
	"fmt"

	"github.com/jrecuero/go-cli/tools"
)

// Tree represents ...
type Tree struct {
	id     Ider
	Label  string
	root   *Leaf
	anchor *Leaf
	path   *Path
	leafs  map[Ider]*Leaf
}

// GetRoot is ...
func (tree *Tree) GetRoot() *Leaf {
	return tree.root
}

// GetAnchor is ...
func (tree *Tree) GetAnchor() *Leaf {
	return tree.anchor
}

// GetLeafs is ...
func (tree *Tree) GetLeafs() map[Ider]*Leaf {
	return tree.leafs
}

// AddBranch adds the given branch to the given parent. If parent is nil, use
// the tree root leaf. Parent attribute in the Child leaf is set properly.
func (tree *Tree) AddBranch(parent *Leaf, branch IBranch) error {
	if parent == nil {
		parent = tree.GetRoot()
		branch.SetParent(parent)
	}
	if !parent.hooked {
		return tools.ERROR(nil, false, "Parent not found in tree: %#v\n", parent)
	}
	if err := parent.AddBranch(branch); err != nil {
		return err
	}
	child := branch.GetChild()
	if err := child.AddParent(parent); err != nil {
		return err
	}
	child.hooked = true
	tree.leafs[child.GetID()] = child
	return nil
}

// AddChild adds an static branch fromt the given parent to the given child.
func (tree *Tree) AddChild(parent *Leaf, child *Leaf) error {
	if parent == nil {
		parent = tree.GetRoot()
	}
	var branch IBranch = StaticBranch(parent, child)
	return tree.AddBranch(parent, branch)
}

// ExistPathTo checks if there is branch from the given anchor to the given
// child. If not anchor leaf is provided, the tree anchor is used instead of.
func (tree *Tree) ExistPathTo(anchor *Leaf, dest *Leaf) (IBranch, bool) {
	if anchor == nil {
		anchor = tree.anchor
	}
	for _, branch := range anchor.Branches {
		if branch.GetChild() == dest {
			return branch, true
		}
	}
	return nil, false
}

// IsPathTo check if there is a branch from the given anchor to the given child
// ana if the path is possible. If not anchor is leaf is provided, the tree
// anchor is used instead of.
func (tree *Tree) IsPathTo(anchor *Leaf, dest *Leaf, params ...interface{}) (IBranch, bool) {
	if anchor == nil {
		anchor = tree.anchor
	}
	if branch, ok := tree.ExistPathTo(anchor, dest); ok {
		if _, bok := branch.Check(params...); bok {
			return branch, true
		}
	}
	return nil, false
}

// PathsFrom returns all existance and possible branches from the given anchor.
func (tree *Tree) PathsFrom(anchor *Leaf, params ...interface{}) []*Leaf {
	var children []*Leaf
	if anchor == nil {
		anchor = tree.anchor
	}
	for _, branch := range anchor.Branches {
		if _, ok := branch.Check(params...); ok {
			children = append(children, branch.GetChild())
		}
	}
	return children
}

// setAnchor is ..
func (tree *Tree) setAnchor(anchor *Leaf) *Leaf {
	tree.anchor = anchor
	return tree.GetAnchor()
}

// AddTraverse adds a branch to the tree traverse.
func (tree *Tree) AddTraverse(branch IBranch) error {
	if branch.GetParent() == nil {
		branch.SetParent(tree.GetRoot())
	}
	if tree.GetAnchor() != branch.GetParent() {
		return tools.ERROR(nil, false, "parent is not the anchor: %#v\n", branch.GetParent())
	}
	tree.setAnchor(branch.GetChild())
	tree.path.Traversed = append(tree.path.Traversed, branch.GetTraverse())
	return nil
}

// SetAnchorTo moves the anchor to the destination leaf and adds the branch to
// the tree traverse.
func (tree *Tree) SetAnchorTo(dest *Leaf) *Leaf {
	for _, branch := range tree.anchor.Branches {
		if branch.GetChild() == dest {
			if err := tree.AddTraverse(branch); err != nil {
				return nil
			}
			return tree.GetAnchor()
		}
	}
	return nil
}

// ToMermaid is ...
func (tree *Tree) ToMermaid() string {
	var buffer bytes.Buffer
	buffer.WriteString("graph LR\n")
	for _, leaf := range tree.GetLeafs() {
		for _, branch := range leaf.Branches {
			buffer.WriteString(branch.ToMermaid())
		}
	}
	return buffer.String()
}

// NewTree is ...
func NewTree(label string) *Tree {
	root := NewLeaf("root/0")
	root.hooked = true
	tree := &Tree{
		id:    nextIder(),
		Label: label,
		root:  root,
		path:  NewPath(fmt.Sprintf("%s/path", label)),
		leafs: make(map[Ider]*Leaf),
	}
	tree.anchor = tree.GetRoot()
	return tree
}
