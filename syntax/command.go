package syntax

import (
	"errors"
	"strings"
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
	ismode         bool
	IsBuiltIn      bool
	Prompt         interface{}
}

// IsCommand returns if content is a command.
func (cmd *Command) IsCommand() bool {
	return cmd.ismode == false
}

// IsMode returns if content is a mode.
func (cmd *Command) IsMode() bool {
	return cmd.ismode
}

// GetStrType returns the short string for the content type.
func (cmd *Command) GetStrType() string {
	if cmd.IsMode() {
		return "M"
	}
	return "C"
}

// LookForArgument searches for an argument in a Command with the given label.
func (cmd *Command) LookForArgument(label string) (*Argument, error) {
	for _, argo := range cmd.Arguments {
		if argo.GetLabel() == label {
			return argo, nil
		}
	}
	return nil, errors.New("not found")
}

// AddNameSpaceName adds a new namespace.
func (cmd *Command) AddNameSpaceName(nsName string) error {
	cmd.NameSpaceNames = append(cmd.NameSpaceNames, nsName)
	return nil
}

// DeleteNameSpaceName deletes a namespace.
func (cmd *Command) DeleteNameSpaceName(nsName string) error {
	for i, nameSpaceNames := range cmd.NameSpaceNames {
		if nameSpaceNames == nsName {
			cmd.NameSpaceNames = append(cmd.NameSpaceNames[:i], cmd.NameSpaceNames[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

// Setup initializes all command fields.
func (cmd *Command) Setup() *Command {
	cmd.CmdSyntax = NewCommandSyntax(cmd.Syntax)
	//cmd.CmdSyntax.CreateGraph(cmd)
	//cmd.label = cmd.CmdSyntax.Parsed.Command
	if cmd.completer == nil {
		cmd.completer = NewCompleterCommand(cmd.GetLabel())
	}
	cmd.FullCmd = cmd.GetLabel()
	for _, argument := range cmd.Arguments {
		argument.Setup()
	}
	return cmd
}

// SetupGraph creates the command syntax graph.
func (cmd *Command) SetupGraph(children bool) *Command {
	//tools.Debug("setup graph %#v\n", cmd.GetLabel())
	cmd.HasChildren = children
	cmd.CmdSyntax.CreateGraph(cmd)
	cmd.label = cmd.CmdSyntax.Parsed.Command
	return cmd
}

var _ IContent = (*Command)(nil)

// NewCommand creates a new command instance.
func NewCommand(parent *Command, syntax string, help string, arguments []*Argument, callbacks *Callback) *Command {
	if callbacks == nil {
		callbacks = NewCallback(nil, nil, nil, nil)
	}
	command := &Command{
		Callback:  callbacks,
		Content:   NewContent(strings.Split(syntax, " ")[0], help, nil).(*Content),
		Syntax:    syntax,
		Arguments: arguments,
		Parent:    parent,
	}
	command.Setup()
	return command
}

// NewMode creates a new command instance.
func NewMode(parent *Command, syntax string, help string, arguments []*Argument, callbacks *Callback) *Command {
	if callbacks == nil {
		callbacks = NewCallback(nil, nil, nil, nil)
	}
	command := &Command{
		Callback:  callbacks,
		Content:   NewContent(strings.Split(syntax, " ")[0], help, nil).(*Content),
		Syntax:    syntax,
		Arguments: arguments,
		Parent:    parent,
		ismode:    true,
	}
	command.Setup()
	return command
}
