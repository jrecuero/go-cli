package syntax

// Context represents the structure that stores information about any match.
type Context struct {
	Matched   []*ContentNode
	Arguments interface{}
}

// NewContext creates a new Context instance.
func NewContext() *Context {
	ctx := &Context{}
	return ctx
}

// AddToken adds a matched token to the context.
func (c *Context) AddToken(n *ContentNode) error {
	c.Matched = append(c.Matched, n)
	return nil
}

// GetValueForArgument returns the value for the given field in arguments
// passed to.
func (c *Context) GetValueForArgument(field string) interface{} {
	argos := c.Arguments.(map[string]interface{})
	return argos[field]
}

// Clean cleans context content.
func (c *Context) Clean() error {
	c.Matched = nil
	c.Arguments = nil
	return nil
}
