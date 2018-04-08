package syntax

import "github.com/jrecuero/go-cli/graph"

// Context represents the structure that stores information about any match.
type Context struct {
	Matched []*graph.Node
}

// NewContext creates a new Context instance.
func NewContext() *Context {
	ctx := &Context{}
	return ctx
}

// AddToken adds a matched token to the context.
func (c *Context) AddToken(n *graph.Node) bool {
	c.Matched = append(c.Matched, n)
	return true
}
