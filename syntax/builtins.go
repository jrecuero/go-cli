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

func newHelpCmd() *Command {
	helpCmd := NewCommand(nil, "help", "Display help", nil, nil)
	helpCmd.Callback.Enter = func(ctx *Context, arguments interface{}) error {
		if inst, err := ctx.Cache.Get("nsm"); err == nil {
			nsm := inst.(*NSManager)
			helps := nsm.GetParseTree().GetAllRootFromAnchor(nsm.GetMatcher().Rooter)
			for _, node := range helps {
				content := NodeToContentNode(node).GetContent()
				tools.ToDisplay("[%s] %-16s: %s\n",
					content.GetStrType(), content.GetLabel(), content.GetHelp())
			}
		}
		return nil
	}
	helpCmd.IsBuiltIn = true
	return helpCmd
}

func newSyntaxCmd() *Command {
	syntaxCmd := NewCommand(nil, "syntax", "Display syntax", nil, nil)
	syntaxCmd.Callback.Enter = func(ctx *Context, arguments interface{}) error {
		if inst, err := ctx.Cache.Get("nsm"); err == nil {
			nsm := inst.(*NSManager)
			helps := nsm.GetParseTree().GetAllRootFromAnchor(nsm.GetMatcher().Rooter)
			for _, node := range helps {
				content := NodeToContentNode(node).GetContent()
				tools.ToDisplay("[%s] %-16s: %s\n",
					content.GetStrType(), content.GetLabel(), content.(*Command).CmdSyntax.Syntax)
			}
		}
		return nil
	}
	syntaxCmd.IsBuiltIn = true
	return syntaxCmd
}

func newRecorderCmd() *Command {
	recorderCmd := NewCommand(nil, "record", "Access to the recorder", nil, nil)
	recorderCmd.IsBuiltIn = true
	return recorderCmd
}

func newRecorderStartCmd(parent *Command) *Command {
	recorderStartCmd := NewCommand(parent, "start", "Start recorder", nil, nil)
	recorderStartCmd.Callback.Enter = func(ctx *Context, arguments interface{}) error {
		//params := arguments.(map[string]interface{})
		if nsm, err := ctx.Cache.Get("nsm"); err == nil {
			if nsm.(*NSManager).Record.Start() == nil {
				tools.ToDisplay("Start recorder ...\n")
			}
		}
		return nil
	}
	recorderStartCmd.IsBuiltIn = true
	return recorderStartCmd
}

func newRecorderStopCmd(parent *Command) *Command {
	recorderStopCmd := NewCommand(parent, "stop", "Stop recorder", nil, nil)
	recorderStopCmd.Callback.Enter = func(ctx *Context, arguments interface{}) error {
		//params := arguments.(map[string]interface{})
		if nsm, err := ctx.Cache.Get("nsm"); err == nil {
			if nsm.(*NSManager).Record.Stop() == nil {
				tools.ToDisplay("... Stop recorder\n")
			}
		}
		return nil
	}
	recorderStopCmd.IsBuiltIn = true
	return recorderStopCmd
}

func newRecorderDisplayCmd(parent *Command) *Command {
	recorderDisplayCmd := NewCommand(parent, "display", "Display recorder", nil, nil)
	recorderDisplayCmd.Callback.Enter = func(ctx *Context, arguments interface{}) error {
		//params := arguments.(map[string]interface{})
		if nsm, err := ctx.Cache.Get("nsm"); err == nil {
			nsm.(*NSManager).Record.Display()
		}
		return nil
	}
	recorderDisplayCmd.IsBuiltIn = true
	return recorderDisplayCmd
}

func newRecorderCleanCmd(parent *Command) *Command {
	recorderCleanCmd := NewCommand(parent, "clean", "Clean recorder", nil, nil)
	recorderCleanCmd.Callback.Enter = func(ctx *Context, arguments interface{}) error {
		//params := arguments.(map[string]interface{})
		if nsm, err := ctx.Cache.Get("nsm"); err == nil {
			if nsm.(*NSManager).Record.Clean() == nil {
				tools.ToDisplay("Clean recorder\n")
			}
		}
		return nil
	}
	recorderCleanCmd.IsBuiltIn = true
	return recorderCleanCmd
}

func newRecorderPlayCmd(parent *Command) *Command {
	recorderPlayCmd := NewCommand(parent, "play", "Play recorder", nil, nil)
	recorderPlayCmd.Callback.Enter = func(ctx *Context, arguments interface{}) error {
		//params := arguments.(map[string]interface{})
		if nsm, err := ctx.Cache.Get("nsm"); err == nil {
			if nsm.(*NSManager).Record.Play(nsm.(*NSManager).GetMatcher()) == nil {
				tools.ToDisplay("Play recorder\n")
			}
		}
		return nil
	}
	recorderPlayCmd.IsBuiltIn = true
	return recorderPlayCmd
}

func newRecorderSaveCmd(parent *Command) *Command {
	recorderSaveCmd := NewCommand(parent, "save filename", "Save recorder",
		[]*Argument{
			NewArgument("filename", "Filename to save", nil, "string", "recorder.log", nil),
		}, nil)
	recorderSaveCmd.Callback.Enter = func(ctx *Context, arguments interface{}) error {
		//filename := arguments.(map[string]interface{})["filename"].(string)
		filename := tools.GetStringFromArgs(arguments, "filename")
		if nsm, err := ctx.Cache.Get("nsm"); err == nil {
			if nsm.(*NSManager).Record.Save(filename, false) == nil {
				tools.ToDisplay("Save recorder\n")
			}
		}
		return nil
	}
	recorderSaveCmd.IsBuiltIn = true
	return recorderSaveCmd
}

func newRecorderLoadCmd(parent *Command) *Command {
	recorderLoadCmd := NewCommand(parent, "load", "Load recorder",
		[]*Argument{
			NewArgument("filename", "Filename to load", nil, "string", "recorder.log", nil),
		}, nil)
	recorderLoadCmd.Callback.Enter = func(ctx *Context, arguments interface{}) error {
		//filename := arguments.(map[string]interface{})["filename"].(string)
		filename := tools.GetStringFromArgs(arguments, "filename")
		if nsm, err := ctx.Cache.Get("nsm"); err == nil {
			if nsm.(*NSManager).Record.Load(filename, false) == nil {
				tools.ToDisplay("Load recorder\n")
			}
		}
		return nil
	}
	recorderLoadCmd.IsBuiltIn = true
	return recorderLoadCmd
}

// NewBuiltins returns all builtin commands
func NewBuiltins() []*Command {
	debugCmd := newDebugCommand()
	recorderCmd := newRecorderCmd()
	return []*Command{
		newExitCommand(),
		debugCmd,
		newDebugParseCommand(debugCmd),
		newDebugCommandCmd(debugCmd),
		newDebugExploreCommandTreeCmd(debugCmd),
		newDebugExploreParseTreeCmd(debugCmd),
		newHelpCmd(),
		newSyntaxCmd(),
		recorderCmd,
		newRecorderStartCmd(recorderCmd),
		newRecorderStopCmd(recorderCmd),
		newRecorderDisplayCmd(recorderCmd),
		newRecorderCleanCmd(recorderCmd),
		newRecorderPlayCmd(recorderCmd),
		newRecorderSaveCmd(recorderCmd),
		newRecorderLoadCmd(recorderCmd),
	}
}
