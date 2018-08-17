package grafo

import (
	"bytes"
	"fmt"

	"github.com/jrecuero/go-cli/tools"
)

// IContent represents the interface for any leaf content.
type IContent interface {
	GetLabel() string
}

// IBranch represents ...
type IBranch interface {
	GetParent() *Leaf
	SetParent(*Leaf)
	GetChild() *Leaf
	SetChild(*Leaf)
	GetTraverse() *Traverse
	ToMermaid() string
	Check(params ...interface{}) (interface{}, bool)
}

// ClearanceCb represents ...
type ClearanceCb func(parent *Leaf, child *Leaf, params ...interface{}) (interface{}, bool)

// Ider represents ...
type Ider uint64

var _ider Ider

// nextIder is ...
func nextIder() Ider {
	_ider++
	return _ider
}

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

// Leaf represents ...
type Leaf struct {
	id        Ider
	Label     string
	Parents   []*Leaf
	Branches  []IBranch
	Traversed []*Traverse
	Content   IContent
	hooked    bool
}

// GetID is ...
func (leaf *Leaf) GetID() Ider {
	return leaf.id
}

// AddParent is ...
func (leaf *Leaf) AddParent(parent *Leaf) error {
	leaf.Parents = append(leaf.Parents, parent)
	return nil
}

// AddBranch is ...
func (leaf *Leaf) AddBranch(branch IBranch) error {
	leaf.Branches = append(leaf.Branches, branch)
	return nil
}

// NewLeaf is ...
func NewLeaf(label string) *Leaf {
	return &Leaf{
		id:    nextIder(),
		Label: label,
	}
}

// Path represents ...
type Path struct {
	id        Ider
	Label     string
	Traversed []*Traverse
}

// NewPath is ...
func NewPath(label string) *Path {
	return &Path{
		id:    nextIder(),
		Label: label,
	}
}

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

// AddBranch is ...
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

// AddChild is ...
func (tree *Tree) AddChild(parent *Leaf, child *Leaf) error {
	if parent == nil {
		parent = tree.GetRoot()
	}
	var branch IBranch = StaticBranch(parent, child)
	return tree.AddBranch(parent, branch)
}

// PathsFrom is ...
func (tree *Tree) PathsFrom(anchor *Leaf, params ...interface{}) []*Leaf {
	var children []*Leaf
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

// AddTraverse is ...
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

// SetAnchorTo is ..
func (tree *Tree) SetAnchorTo(anchor *Leaf) *Leaf {
	for _, branch := range tree.anchor.Branches {
		if branch.GetChild() == anchor {
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
