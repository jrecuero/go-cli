package syntax

import (
	"errors"
	"strings"
)

// Command represents any CLI command internally in the system.
type Command struct {
	*Callback                     // command callbacks (enter/exit/...)
	*Content                      // command content
	Syntax         string         // command syntas as a string.
	CmdSyntax      *CommandSyntax //command syntax instance.
	Arguments      []*Argument    // command arguments.
	NameSpaceNames []string       // command namespaces.
	Parent         *Command       //command parent
	HasChildren    bool           // has the command children Ccommands)?
	IsBuiltIn      bool           // is the command a built-on command?
	RunAsNoFinal   bool           // is the command exec as a parent command?
	Prompt         interface{}    // mode prompt (only for modes)
	ismode         bool           // is the command a mode?
}

var commandGraphPattern = "[%s]"
var modeGraphPattern = "{%s}"

// IsCommand returns if content is a command.
func (cmd *Command) IsCommand() bool {
	return cmd.ismode == false
}

// IsMode returns if content is a mode.
func (cmd *Command) IsMode() bool {
	return cmd.ismode
}

//// GetString returns the string for the command content.
//func (cmd *Command) GetString() string {
//    return cmd.GetLabel()
//}

// GetStrType returns the short string for the content type.
func (cmd *Command) GetStrType() string {
	if cmd.IsMode() {
		return "M"
	}
	return "C"
}

// GetGraphPattern returns the string with the graphical pattern.
func (cmd *Command) GetGraphPattern() *string {
	if cmd.ismode {
		return &modeGraphPattern
	}
	return &commandGraphPattern
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
	// Graph can not be setup at this point, it requires the command tree in
	// order to identify possible command children.
	// cmd.CmdSyntax.CreateGraph(cmd)
	// cmd.label = cmd.CmdSyntax.Parsed.Command
	if cmd.completer == nil {
		cmd.completer = NewCompleterCommand(cmd.GetLabel())
	}
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
