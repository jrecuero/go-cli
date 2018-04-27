package syntax

import "errors"

//// Command represents any CLI command internally in the system.
//type Command struct {
//    Callback
//    Syntax      string
//    CmdSyntax   *CommandSyntax
//    Label       string
//    Help        string
//    Arguments   []*Argument
//    Prefixes    []*Prefix
//    Completer   ICompleter
//    FullCmd     string
//    Namespace   string
//    ToNamespace string
//}

// Command represents any CLI command internally in the system.
type Command struct {
	Callback
	*Content
	Syntax      string
	CmdSyntax   *CommandSyntax
	Arguments   []*Argument
	Prefixes    []*Prefix
	FullCmd     string
	Namespace   string
	ToNamespace string
}

//// GetLabel returns user command label.
//func (c *Command) GetLabel() string {
//    return c.Label
//}

//// GetType returns user command type.
//func (c *Command) GetType() string {
//    return "string"
//}

//// GetDefault returns user command default value.
//func (c *Command) GetDefault() interface{} {
//    return c.Label
//}

//// GetHelp returns user command help.
//func (c *Command) GetHelp() string {
//    return c.Help
//}

//// GetCompleter returns user command completer.
//func (c *Command) GetCompleter() ICompleter {
//    return c.Completer
//}

//// ToString returns the string with the content information.
//func (c *Command) ToString() string {
//    return c.Label
//}

// IsCommand returns if content is a command.
func (c *Command) IsCommand() bool {
	return true
}

// LookForPrefix searches for an prefix in a Command with the given label.
func (c *Command) LookForPrefix(label string) (*Prefix, error) {
	for _, prefix := range c.Prefixes {
		if prefix.GetLabel() == label {
			return prefix, nil
		}
	}
	return nil, errors.New("not found")
}

// LookForArgument searches for an argument in a Command with the given label.
func (c *Command) LookForArgument(label string) (*Argument, error) {
	for _, argo := range c.Arguments {
		if argo.GetLabel() == label {
			return argo, nil
		}
	}
	return nil, errors.New("not found")
}

// Setup initializes all command fields.
func (c *Command) Setup() error {
	c.CmdSyntax = NewCommandSyntax(c.Syntax)
	c.CmdSyntax.CreateGraph(c)
	c.label = c.CmdSyntax.Parsed.Command
	if c.completer == nil {
		c.completer = NewCompleterCommand(c)
	} else {
		c.completer.Setup(c)
	}
	c.FullCmd = c.GetLabel()
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
