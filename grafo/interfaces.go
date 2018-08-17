package grafo

// Ider represents ...
type Ider uint64

var _ider Ider

// IContent represents the interface for any leaf content.
type IContent interface {
	GetLabel() string
}

// nextIder is ...
func nextIder() Ider {
	_ider++
	return _ider
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
