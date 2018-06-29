package syntax

import (
	"errors"
)

// Command represents any CLI command internally in the system.
type Command struct {
	*Callback
	*Content
	Syntax         string
	CmdSyntax      *CommandSyntax
	Arguments      []*Argument
	FullCmd        string
	NameSpaceNames []string
	ToNameSpace    string
	Parent         *Command
	HasChildren    bool
}

// IsCommand returns if content is a command.
func (c *Command) IsCommand() bool {
	return true
}

// GetStrType returns the short string for the content type.
func (c *Command) GetStrType() string {
	return "C"
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

// AddNameSpaceName adds a new namespace.
func (c *Command) AddNameSpaceName(nsName string) error {
	c.NameSpaceNames = append(c.NameSpaceNames, nsName)
	return nil
}

// DeleteNameSpaceName deletes a namespace.
func (c *Command) DeleteNameSpaceName(nsName string) error {
	for i, nameSpaceNames := range c.NameSpaceNames {
		if nameSpaceNames == nsName {
			c.NameSpaceNames = append(c.NameSpaceNames[:i], c.NameSpaceNames[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

// Setup initializes all command fields.
func (c *Command) Setup() *Command {
	c.CmdSyntax = NewCommandSyntax(c.Syntax)
	//c.CmdSyntax.CreateGraph(c)
	//c.label = c.CmdSyntax.Parsed.Command
	if c.completer == nil {
		c.completer = NewCompleterCommand(c.GetLabel())
	}
	c.FullCmd = c.GetLabel()
	for _, argument := range c.Arguments {
		argument.Setup()
	}
	return c
}

// SetupGraph creates the command syntax graph.
func (c *Command) SetupGraph(children bool) *Command {
	c.HasChildren = children
	c.CmdSyntax.CreateGraph(c)
	c.label = c.CmdSyntax.Parsed.Command
	return c
}

var _ IContent = (*Command)(nil)

// NewCommand creates a new command instance.
func NewCommand(parent *Command, syntax string, help string, arguments []*Argument, callbacks *Callback) *Command {
	if callbacks == nil {
		callbacks = NewCallback(nil, nil, nil, nil)
	}
	command := &Command{
		Callback:  callbacks,
		Content:   NewContent(syntax, help, nil).(*Content),
		Syntax:    syntax,
		Arguments: arguments,
		Parent:    parent,
	}
	command.Setup()
	return command
}
