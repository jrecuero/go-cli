package novel

import (
	"github.com/jrecuero/go-cli/syntax"
	"github.com/jrecuero/go-cli/tools"
)

// SetupCommands configures all command to run.
func SetupCommands() []*syntax.Command {

	displayCommand := syntax.NewCommand(
		nil,
		"display",
		"Display ...",
		nil,
		nil)
	displayCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		//params := arguments.(map[string]interface{})
		for _, actor := range TheNovel.Actors {
			tools.ToDisplay("Actors:\t%#v\n", actor)
		}
		return nil
	}

	actionCommand := syntax.NewCommand(
		nil,
		"action action-str",
		"Action ...",
		[]*syntax.Argument{
			syntax.NewArgument(
				"action-str",
				"Action to execute",
				nil,
				"string",
				"none",
				nil),
		},
		nil)
	actionCommand.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		params := arguments.(map[string]interface{})
		actionStr := params["action-str"].(string)
		//tools.ToDisplay("%#v\n", actionStr)
		//ae := n.Compile(actionStr)
		TheNovel.Execute(actionStr)
		//tools.ToDisplay("%#v\n", ae)
		return nil
	}

	commands := []*syntax.Command{
		actionCommand,
		displayCommand,
	}
	return commands
}
