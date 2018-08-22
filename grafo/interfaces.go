package grafo

// Ider represents ...
type Ider uint64

var _ider Ider

// IContent represents the interface for any vertex content.
type IContent interface {
	GetLabel() string
}

// nextIder is ...
func nextIder() Ider {
	_ider++
	return _ider
}

// IEdge represents ...
type IEdge interface {
	GetParent() IVertex
	SetParent(IVertex)
	GetChild() IVertex
	SetChild(IVertex)
	GetVtoV() *VtoV
	ToMermaid() string
	Check(params ...interface{}) (interface{}, bool)
}

// ClearanceCb represents ...
type ClearanceCb func(parent IVertex, child IVertex, params ...interface{}) (interface{}, bool)
