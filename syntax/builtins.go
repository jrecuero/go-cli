package syntax

import (
	"os"

	"github.com/jrecuero/go-cli/tools"
)

func init() {
}

// NewExitCommand generates a new exit command.
func NewExitCommand() *Command {
	exitCmd := NewCommand(nil, "exit", "Exit application", nil, nil)
	exitCmd.Callback.Enter = func(ctx *Context, arguments interface{}) error {
		if len(ctx.Modes) == 0 {
			os.Exit(0)
			return nil
		}
		ctx.SetProcess(tools.PString(POPMODE))
		return nil
	}
	exitCmd.IsBuiltIn = true
	return exitCmd
}
