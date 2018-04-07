package syntax

// Context represents the structure that stores information about any match.
type Context struct {
	Matched []*Node
}

// AddToken adds a matched token to the context.
func (c *Context) AddToken(n *Node) bool {
	c.Matched = append(c.Matched, n)
	return true
}
