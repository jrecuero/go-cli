package grafo

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
