package commands

import (
	"github.com/jrecuero/go-cli/syntax"
	"github.com/jrecuero/go-cli/tools"
)

// SetupCommands is ...
func SetupCommands() []*syntax.Command {
	displayTechsCommand := syntax.NewCommand(
		nil,
		"display-techs",
		"Display all available techniques.",
		nil,
		nil)
	displayTechsCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		tools.ToDisplay("Display all techs...\n")
		return nil
	}

	cmds := []*syntax.Command{
		displayTechsCommand,
	}
	return cmds
}
