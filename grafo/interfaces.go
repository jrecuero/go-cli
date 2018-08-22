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
	GetParent() *Vertex
	SetParent(*Vertex)
	GetChild() *Vertex
	SetChild(*Vertex)
	GetVtoV() *VtoV
	ToMermaid() string
	Check(params ...interface{}) (interface{}, bool)
}

// ClearanceCb represents ...
type ClearanceCb func(parent *Vertex, child *Vertex, params ...interface{}) (interface{}, bool)
