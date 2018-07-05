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
	lastcmd   *Command
	cmdbox    []*Command
}

// GetLastCommand retrieves the lastcmd field.
func (ctx *Context) GetLastCommand() *Command {
	return ctx.lastcmd
}

// SetLastCommand sets the lastcmd field.
func (ctx *Context) SetLastCommand(cmd *Command) {
	ctx.cmdbox = append(ctx.cmdbox, cmd)
	ctx.lastcmd = cmd
}

// GetCommandBox rettrieves the cmdbox field.
func (ctx *Context) GetCommandBox() []*Command {
	return ctx.cmdbox
}

// AddToken adds a matched token to the context.
func (ctx *Context) AddToken(n *ContentNode, v string) error {
	token := NewToken(n, v)
	ctx.Matched = append(ctx.Matched, token)
	return nil
}

// GetValueForArgument returns the value for the given field in arguments
// passed to.
func (ctx *Context) GetValueForArgument(field string) interface{} {
	argos := ctx.Arguments.(map[string]interface{})
	return argos[field]
}

// Clean cleans context content.
func (ctx *Context) Clean() error {
	ctx.Matched = nil
	ctx.Arguments = nil
	ctx.lastcmd = nil
	ctx.cmdbox = nil
	return nil
}

// NewContext creates a new Context instance.
func NewContext() *Context {
	ctx := &Context{}
	return ctx
}
