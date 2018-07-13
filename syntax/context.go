package syntax

import (
	"fmt"

	"github.com/jrecuero/go-cli/tools"
)

const (
	// MATCH identifies match process.
	MATCH = "match"

	// COMPLETE identifies complete process.
	COMPLETE = "complete"

	// HELP identifies help process.
	HELP = "help"

	// QUERY identifies query process.
	QUERY = "query"

	// EXECUTE identifies execute process.
	EXECUTE = "execute"
)

// Token represents the structure that stores information with any token that
// has been matched.
type Token struct {
	Node  *ContentNode
	Value interface{}
}

// ArgValue represents the structure used to store the argument values being
// marched.
type ArgValue struct {
	Arg   *Argument
	Value interface{}
}

// CommandBox represents the struture for any command with arguments matched.
type CommandBox struct {
	Cmd       *Command
	ArgValues []*ArgValue
}

// NewToken creates a new Token instance.
func NewToken(cn *ContentNode, value interface{}) *Token {
	return &Token{
		Node:  cn,
		Value: value,
	}
}

// Context represents the structure that stores information about any match.
type Context struct {
	Matched []*Token
	lastcmd *Command
	cmdbox  []*CommandBox
	process *string
}

// SetProcess sets the context process running.
func (ctx *Context) SetProcess(process *string) bool {
	proc := tools.String(process)
	if proc == MATCH || proc == COMPLETE || proc == HELP || proc == QUERY || proc == EXECUTE {
		ctx.process = process
		return true
	}
	return false
}

// GetProcess retrieves the context process runnning.
func (ctx *Context) GetProcess() string {
	return tools.String(ctx.process)
}

// GetLastCommand retrieves the lastcmd field.
func (ctx *Context) GetLastCommand() *Command {
	return ctx.lastcmd
}

// SetLastCommand sets the lastcmd field.
func (ctx *Context) SetLastCommand(cmd *Command) {
	ctx.cmdbox = append(ctx.cmdbox, &CommandBox{Cmd: cmd})
	ctx.lastcmd = cmd
}

// SetLastArgument sets the last argument.
func (ctx *Context) SetLastArgument(arg *Argument, value interface{}) {
	index := len(ctx.cmdbox) - 1
	ctx.cmdbox[index].ArgValues = append(ctx.cmdbox[index].ArgValues, &ArgValue{Arg: arg, Value: value})
}

// GetCommandBox rettrieves the cmdbox field.
func (ctx *Context) GetCommandBox() []*CommandBox {
	return ctx.cmdbox
}

// GetCmdBoxIndexForCommandLabel retrieves the index in the cmdbox field for the
// given command label.
func (ctx *Context) GetCmdBoxIndexForCommandLabel(label *string) (int, error) {
	if label == nil {
		return len(ctx.cmdbox) - 1, nil
	}
	for i, cbox := range ctx.cmdbox {
		if cbox.Cmd.GetLabel() == tools.String(label) {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Command with label %s not found", tools.String(label))
}

// GetArgValueForArgLabel retrieves the argument value for the given argument
// label.
func (ctx *Context) GetArgValueForArgLabel(cmdlabel *string, arglabel string) (interface{}, error) {
	if icmd, err := ctx.GetCmdBoxIndexForCommandLabel(cmdlabel); err == nil {
		for _, argval := range ctx.cmdbox[icmd].ArgValues {
			if argval.Arg.GetLabel() == arglabel {
				// TODO: Any argument type processing should be done at this
				// point.
				//return argval.Value, nil
				return argval.Arg.Cast(argval.Value.(string))
			}
		}
	}
	return nil, fmt.Errorf("Argument %s not found for Command %s", arglabel, tools.String(cmdlabel))
}

// GetArgValuesForCommandLabel retrieves all arguments for the given command
// label.
func (ctx *Context) GetArgValuesForCommandLabel(cmdlabel *string) (interface{}, error) {
	result := make(map[string]interface{})
	if icmd, err := ctx.GetCmdBoxIndexForCommandLabel(cmdlabel); err == nil {
		for _, argval := range ctx.cmdbox[icmd].ArgValues {
			// TODO: Any argument type processing should be done at this
			// point.
			//result[argval.Arg.GetLabel()] = argval.Value
			if r, err := argval.Arg.Cast(argval.Value.(string)); err == nil {
				result[argval.Arg.GetLabel()] = r
			} else {
				return nil, fmt.Errorf("Argument casting failed %#v", err)
			}
		}
		return result, nil
	}
	return nil, fmt.Errorf("Arguments not found for Command %s", tools.String(cmdlabel))
}

// AddToken adds a matched token to the context.
func (ctx *Context) AddToken(cn *ContentNode, value interface{}) error {
	token := NewToken(cn, value)
	ctx.Matched = append(ctx.Matched, token)
	if cn.GetContent().IsCommand() {
		ctx.SetLastCommand(cn.GetContent().(*Command))
	} else if cn.GetContent().IsArgument() {
		ctx.SetLastArgument(cn.GetContent().(*Argument), value)
	}
	return nil
}

// Clean cleans context content.
func (ctx *Context) Clean() error {
	ctx.Matched = nil
	ctx.lastcmd = nil
	ctx.cmdbox = nil
	return nil
}

// NewContext creates a new Context instance.
func NewContext() *Context {
	ctx := &Context{}
	return ctx
}
