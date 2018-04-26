package syntax

// Command represents any CLI command internally in the system.
type Command struct {
	Callback
	Syntax      string
	CmdSyntax   *CommandSyntax
	Label       string
	Help        string
	Arguments   []Argument
	Completer   ICompleter
	Namespace   string
	ToNamespace string
}

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

// ToString returns the string with the content information.
func (c *Command) ToString() string {
	return c.label
}

// Setup initializes all command fields.
func (c *Command) Setup() error {
	c.CmdSyntax = NewCommandSyntax(c.Syntax)
	c.CmdSyntax.CreateGraph()
	return nil
}
