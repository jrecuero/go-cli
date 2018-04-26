package syntax

// Command represents any CLI command internally in the system.
type Command struct {
	Callback
	Syntax      string
	CmdSyntax   *CommandSyntax
	Label       string
	Help        string
	Arguments   []*Argument
	Prefixes    []*Prefix
	Completer   ICompleter
	FullCmd     string
	Namespace   string
	ToNamespace string
}

// GetLabel returns user command label.
func (c *Command) GetLabel() string {
	return c.Label
}

// GetType returns user command type.
func (c *Command) GetType() string {
	return "string"
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
	return c.Label
}

// LookForPrefix searches for an prefix in a Command with the given label.
func (c *Command) LookForPrefix(label string) *Prefix {
	for _, prefix := range c.Prefixes {
		if prefix.GetLabel() == label {
			return prefix
		}
	}
	return nil
}

// LookForArgument searches for an argument in a Command with the given label.
func (c *Command) LookForArgument(label string) *Argument {
	for _, argo := range c.Arguments {
		if argo.GetLabel() == label {
			return argo
		}
	}
	return nil
}

// Setup initializes all command fields.
func (c *Command) Setup() error {
	c.CmdSyntax = NewCommandSyntax(c.Syntax)
	c.CmdSyntax.CreateGraph(c)
	c.Label = c.CmdSyntax.Parsed.Command
	if c.Completer == nil {
		c.Completer = NewCompleterCommand(c)
	} else {
		c.Completer.Setup(c)
	}
	c.FullCmd = c.Label
	for _, prefix := range c.Prefixes {
		c.FullCmd += " " + prefix.GetLabel()
		prefix.Setup()
	}
	for _, argument := range c.Arguments {
		argument.Setup()
	}
	return nil
}

var _ IContent = (*Command)(nil)
var _ ICallback = (*Command)(nil)
