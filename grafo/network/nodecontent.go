package network

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
