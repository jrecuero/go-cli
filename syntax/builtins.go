package syntax

import (
	"os"
	"strings"

	"github.com/jrecuero/go-cli/tools"
)

func init() {
}

type completerExit struct {
	*CompleterCommand
}

// Help returns the help for any node completer.
func (ce *completerExit) Help(ctx *Context, content IContent, line interface{}, index int) (interface{}, bool) {
	tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
	helpSlice := strings.Split(content.GetHelp(), "|")
	var i int
	if len(ctx.Modes) == 0 {
		i = 0
	} else {
		i = 1
	}
	tokens := line.([]string)
	ilast := len(tokens) - 1
	if tokens[ilast] == "" || strings.HasPrefix(content.GetLabel(), tokens[ilast]) {
		return helpSlice[i], true
	}
	return nil, false
}

// NewExitCommand generates a new exit command.
func NewExitCommand() *Command {
	exitCmd := NewCommand(nil, "exit", "Exit application|Exit mode", nil, nil)
	exitCmd.Callback.Enter = func(ctx *Context, arguments interface{}) error {
		if len(ctx.Modes) == 0 {
			os.Exit(0)
			return nil
		}
		ctx.SetProcess(tools.PString(POPMODE))
		return nil
	}
	exitCmd.IsBuiltIn = true
	exitCmd.SetCompleter(&completerExit{NewCompleterCommand(exitCmd.GetLabel())})
	return exitCmd
}
