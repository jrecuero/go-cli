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
	//tools.Tracer("line: %#v | index: %d | label: %#v\n", line, index, content.GetLabel())
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

// newExitCommand generates a new exit command.
func newExitCommand() *Command {
	exitCmd := NewCommand(nil, "exit", "Exit application|Exit mode", nil, nil)
	exitCmd.Callback.Enter = func(ctx *Context, arguments interface{}) error {
		if len(ctx.Modes) == 0 {
			os.Exit(0)
			return nil
		}
		ctx.GetProcess().Append(POPMODE)
		return nil
	}
	exitCmd.IsBuiltIn = true
	exitCmd.SetCompleter(&completerExit{NewCompleterCommand(exitCmd.GetLabel())})
	return exitCmd
}

// newDebugCommand generates a new debug mode
func newDebugCommand() *Command {
	debugCmd := NewMode(nil, "debug", "Debug mode", nil, nil)
	debugCmd.Callback.Enter = func(ctx *Context, arguments interface{}) error {
		return nil
	}
	debugCmd.Prompt = "debug>>> "
	debugCmd.IsBuiltIn = true
	return debugCmd
}

func newDebugParseCommand(parent *Command) *Command {
	debugParseCmd := NewCommand(parent, "parse", "Display parse tree", nil, nil)
	debugParseCmd.Callback.Enter = func(ctx *Context, arguments interface{}) error {
		if nsm, err := ctx.Cache.Get("nsm"); err == nil {
			tools.ToDisplay(nsm.(*NSManager).GetParseTree().ToMermaid())
		}
		return nil
	}
	debugParseCmd.IsBuiltIn = true
	return debugParseCmd
}

func newDebugCommandCmd(parent *Command) *Command {
	debugCommandCmd := NewCommand(parent, "command", "Display command tree", nil, nil)
	debugCommandCmd.Callback.Enter = func(ctx *Context, arguments interface{}) error {
		if nsm, err := ctx.Cache.Get("nsm"); err == nil {
			tools.ToDisplay(nsm.(*NSManager).GetCommandTree().ToMermaid())
		}
		return nil
	}
	debugCommandCmd.IsBuiltIn = true
	return debugCommandCmd
}

func newDebugExploreCommandTreeCmd(parent *Command) *Command {
	debugExploreCommandTreeCmd := NewCommand(parent, "explore-command", "Explore command tree", nil, nil)
	debugExploreCommandTreeCmd.Callback.Enter = func(ctx *Context, arguments interface{}) error {
		if nsm, err := ctx.Cache.Get("nsm"); err == nil {
			nsm.(*NSManager).GetCommandTree().Explore()
		}
		return nil
	}
	debugExploreCommandTreeCmd.IsBuiltIn = true
	return debugExploreCommandTreeCmd
}

func newDebugExploreParseTreeCmd(parent *Command) *Command {
	debugExploreParseTreeCmd := NewCommand(parent, "explore-parse", "Explore parse tree", nil, nil)
	debugExploreParseTreeCmd.Callback.Enter = func(ctx *Context, arguments interface{}) error {
		if nsm, err := ctx.Cache.Get("nsm"); err == nil {
			nsm.(*NSManager).GetParseTree().Explore()
		}
		return nil
	}
	debugExploreParseTreeCmd.IsBuiltIn = true
	return debugExploreParseTreeCmd
}

// NewBuiltins returns all builtin commands
func NewBuiltins() []*Command {
	debugCmd := newDebugCommand()
	return []*Command{
		newExitCommand(),
		debugCmd,
		newDebugParseCommand(debugCmd),
		newDebugCommandCmd(debugCmd),
		newDebugExploreCommandTreeCmd(debugCmd),
		newDebugExploreParseTreeCmd(debugCmd),
	}
}
