package syntax

// Command represents any CLI command internally in the system.
type Command struct {
	Callback
	//cb          Callback
	Syntax      string
	CmdSyntax   *CommandSyntax
	Label       string
	Help        string
	Arguments   []Argument
	Completer   ICompleter
	Namespace   string
	ToNamespace string
}

//// Enter is the default Command, it executes with the running context.
//func (c *Command) Enter(ctx *Context, arguments interface{}) error {
//    return c.cb.Enter(ctx, arguments)
//}

//// PostEnter belongs only to mode commands and it executes with the new
//// namespace context.
//func (c *Command) PostEnter(ctx *Context, arguments interface{}) error {
//    return c.cb.PostEnter(ctx, arguments)
//}

//// Exit belongs only to mode commands and it executes with the running
//// namespace context.
//func (c *Command) Exit(ctx *Context) error {
//    return c.cb.Exit(ctx)
//}

//// PostExit belongs only to mode commands and it execute withe the parent
//// namespace context, which will become the actual context.
//func (c *Command) PostExit(ctx *Context) error {
//    return c.cb.PostExit(ctx)
//}

// GetLabel returns user command label.
func (c *Command) GetLabel() string {
	return c.Label
}

// GetType returns user command type.
func (c *Command) GetType() string {
	return ""
}

// GetDefault returns user command default value.
func (c *Command) GetDefault() interface{} {
	return c.Label
}

// GetHelp returns user command help.
func (c *Command) GetHelp() string {
	return c.Help
}

// GetCompleter returns user command completer.
func (c *Command) GetCompleter() ICompleter {
	return c.Completer
}

// Setup initializes all command fields.
func (c *Command) Setup() error {
	return nil
}
