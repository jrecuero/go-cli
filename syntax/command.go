package syntax

import "errors"

// Command represents any CLI command internally in the system.
type Command struct {
	*Callback
	*Content
	Syntax      string
	CmdSyntax   *CommandSyntax
	Arguments   []*Argument
	Prefixes    []*Prefix
	FullCmd     string
	NameSpace   []string
	ToNameSpace string
}

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

// AddNameSpace adds a new namespace.
func (c *Command) AddNameSpace(ns string) error {
	c.NameSpace = append(c.NameSpace, ns)
	return nil
}

// DeleteNameSpace deletes a namespace.
func (c *Command) DeleteNameSpace(ns string) error {
	for i, namespace := range c.NameSpace {
		if namespace == ns {
			c.NameSpace = append(c.NameSpace[:i], c.NameSpace[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
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
