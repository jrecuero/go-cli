package grafo

import (
	"bytes"
	"fmt"
	"strconv"
)

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
	id       Ider
	Label    string
	Branches []IBranch
}

// String is ...
func (path *Path) String() string {
	var buffer bytes.Buffer
	if len(path.Branches) != 0 {
		for _, b := range path.Branches {
			buffer.WriteString(fmt.Sprintf("%s : ", b.GetParent().Label))
		}
		buffer.WriteString(fmt.Sprintf("%s", path.Branches[len(path.Branches)-1].GetChild().Label))
	}
	return buffer.String()
}

// NewPath is ...
func NewPath(label string) *Path {
	id := nextIder()
	if label == "" {
		label = strconv.Itoa(int(id))
	}
	return &Path{
		id:    id,
		Label: label,
	}
}
