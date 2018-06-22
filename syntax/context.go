package syntax

// Token represents the structure that stores information with any token that
// has been matched.
type Token struct {
	Node  *ContentNode
	Value string
}

// NewToken creates a new Token instance.
func NewToken(cn *ContentNode, v string) *Token {
	return &Token{
		Node:  cn,
		Value: v,
	}
}

// Context represents the structure that stores information about any match.
type Context struct {
	Matched   []*Token
	Arguments interface{}
}

// NewContext creates a new Context instance.
func NewContext() *Context {
	ctx := &Context{}
	return ctx
}

// AddToken adds a matched token to the context.
func (c *Context) AddToken(n *ContentNode, v string) error {
	token := NewToken(n, v)
	c.Matched = append(c.Matched, token)
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
