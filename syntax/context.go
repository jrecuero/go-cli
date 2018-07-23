package syntax

import (
	"fmt"

	"github.com/jrecuero/go-cli/graph"
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

	// POPMODE identifies popup node process.
	POPMODE = "popmode"

	// RUNASNOFINAL identifes running a command as not the final one.
	RUNASNOFINAL = "run-as-no-final"

	// DEFAULTPROMPT identifies default prompt.
	DEFAULTPROMPT = ">>> "
)

// ContextProcess represents the process running in the context.
type ContextProcess struct {
	Process []string
}

// Set assigns a unique value to the context process.
func (cp *ContextProcess) Set(proc string) error {
	if proc == MATCH || proc == COMPLETE || proc == HELP ||
		proc == QUERY || proc == EXECUTE || proc == POPMODE || proc == RUNASNOFINAL {
		cp.Clean()
		cp.Process = append(cp.Process, proc)
		return nil
	}
	return tools.ERROR(nil, false, "Unknown process %#v\n", proc)
}

// Append adds an additional value to the context process.
func (cp *ContextProcess) Append(proc string) error {
	if proc == MATCH || proc == COMPLETE || proc == HELP ||
		proc == QUERY || proc == EXECUTE || proc == POPMODE || proc == RUNASNOFINAL {
		cp.Process = append(cp.Process, proc)
		return nil
	}
	return tools.ERROR(nil, false, "Unknown process %#v\n", proc)
}

// Clean clears all values in the context process.
func (cp *ContextProcess) Clean() error {
	cp.Process = []string{}
	return nil
}

// Check finds in the given value is in the context process.
func (cp *ContextProcess) Check(proc ...string) (bool, error) {
	for _, ppar := range proc {
		if ppar == MATCH || ppar == COMPLETE || ppar == HELP ||
			ppar == QUERY || ppar == EXECUTE || ppar == POPMODE || ppar == RUNASNOFINAL {
			for _, pin := range cp.Process {
				if pin == ppar {
					return true, nil
				}
			}
		} else {
			return false, tools.ERROR(nil, false, "Unknown process %#v\n", proc)
		}
	}
	return false, nil
}

// Remove removes the given value from the context process.
func (cp *ContextProcess) Remove(proc ...string) (bool, error) {
	processes := []string{}
	for _, ppar := range proc {
		if ppar == MATCH || ppar == COMPLETE || ppar == HELP ||
			ppar == QUERY || ppar == EXECUTE || ppar == POPMODE || ppar == RUNASNOFINAL {
			for _, pin := range cp.Process {
				if pin != ppar {
					processes = append(processes, pin)
				}
			}
		} else {
			return false, tools.ERROR(nil, false, "Unknown process %#v\n", proc)
		}
	}
	cp.Process = processes
	return true, nil
}

// NewContextProcess creates a new ContextProcess instance.
func NewContextProcess() *ContextProcess {
	return &ContextProcess{
		Process: []string{},
	}
}

// Cache represnets the context cache.
type Cache struct {
	data map[string]interface{} // data storage for the cache.
}

// Add enters a new data in the context cache.
func (cache *Cache) Add(key string, data interface{}) error {
	cache.data[key] = data
	return nil
}

// Get returns the data for an entry in the context cache.
func (cache *Cache) Get(key string) (interface{}, error) {
	if data, ok := cache.data[key]; ok {
		return data, nil
	}
	return nil, tools.ERROR(nil, false, "not found")
}

// GetAll returns all cache data.
func (cache *Cache) GetAll() (map[string]interface{}, error) {
	return cache.data, nil
}

// Clean removes all entries from the cache.
func (cache *Cache) Clean() error {
	cache.data = make(map[string]interface{})
	return nil
}

// NewCache creates a new Cache instance.
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

// Token represents the structure that stores information with any token that
// has been matched.
type Token struct {
	Node  *ContentNode // content node matched.
	Value interface{}  // value for the matched content node.
}

// ArgValue represents the structure used to store the argument values being
// marched.
type ArgValue struct {
	Arg   *Argument   // argument matched.
	Value interface{} // value for the matched argument.
}

// CommandBox represents the struture for any command with arguments matched.
type CommandBox struct {
	Cmd       *Command    // command matched.
	ArgValues []*ArgValue // arguments for the matched command.
}

// ModeBox represents the structure for any mode.
type ModeBox struct {
	Prompt string        // context mode prompt.
	Anchor *graph.Node   // last root used to anchor matcher functionality.
	Mode   *Command      // last mode being matched.
	CmdBox []*CommandBox // arrays with all previous commands matched.
}

// NewToken creates a new Token instance.
func NewToken(cn *ContentNode, value interface{}) *Token {
	return &Token{
		Node:  cn,    // Content node matched.
		Value: value, // Value entered in the comamnd line for the token.
	}
}

// Context represents the structure that stores information about any match.
type Context struct {
	Matched []*Token        // array with all tokens matched in command line.
	Modes   []*ModeBox      // array with all modes entered in command line.
	Cache   *Cache          // context cache to be used by command methods.
	prompt  string          // mode prompt for the running mode.
	lastcmd *Command        // last command found in the matched tokens.
	cmdbox  []*CommandBox   // array with all command matched.
	process *ContextProcess // status process running the context.
}

// GetProcess retrieves the context process runnning.
func (ctx *Context) GetProcess() *ContextProcess {
	return ctx.process
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
	indexCmd := len(ctx.cmdbox) - 1
	ctx.cmdbox[indexCmd].ArgValues = append(ctx.cmdbox[indexCmd].ArgValues, &ArgValue{Arg: arg, Value: value})
}

// UpdateMatched updates matched attribute with number of tokens entered.
func (ctx *Context) UpdateMatched(threshold int) {
	//tools.Debug("threshold: %d matched: %#v\n", threshold, ctx.Matched)
	if threshold <= len(ctx.Matched) {
		if threshold == 0 {
			ctx.Matched = nil
		} else {
			ctx.Matched = ctx.Matched[0 : threshold-1]
		}
	}
	//tools.Debug("matched: %#v\n", ctx.Matched)
}

//UpdateCommandBox updates cmdbox attribute with values from Matched.
func (ctx *Context) UpdateCommandBox() {
	for _, tok := range ctx.Matched {
		cn := tok.Node
		if cn.GetContent().IsCommand() || cn.GetContent().IsMode() {
			ctx.SetLastCommand(cn.GetContent().(*Command))
		} else if cn.GetContent().IsArgument() {
			ctx.SetLastArgument(cn.GetContent().(*Argument), tok.Value)
		}
	}
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
				v, err := argval.Arg.Cast(argval.Value.(string))
				if err != nil {
					return nil, tools.ERROR(err, false, "%#v\n", err)
				}
				return v, nil
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
			if r, err := argval.Arg.Cast(argval.Value.(string)); err == nil {
				result[argval.Arg.GetLabel()] = argval.Arg.Assign(result[argval.Arg.GetLabel()], r)
				//tools.Debug("set arg %#v : %#v result: %#v\n", argval.Arg.GetLabel(), r, result)
			} else {
				return nil, tools.ERROR(err, false, "Argument casting failed %#v", err)
			}
		}
		for _, argo := range ctx.cmdbox[icmd].Cmd.Arguments {
			//tools.Debug("argo %#v : %#v\n", argo.GetLabel(), argo.Default)
			if result[argo.GetLabel()] == nil {
				result[argo.GetLabel()] = argo.Default
			}
		}
		return result, nil
	}
	return nil, tools.ERROR(nil, false, "Arguments not found for Command %s", tools.String(cmdlabel))
}

// AddToken adds a matched token to the context.
func (ctx *Context) AddToken(index int, cn *ContentNode, value interface{}) error {
	token := NewToken(cn, value)
	if len(ctx.Matched) < index+1 {
		ctx.Matched = append(ctx.Matched, token)
	} else {
		ctx.Matched[index] = token
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

// CleanAll cleans context content.
func (ctx *Context) CleanAll() error {
	ctx.Matched = nil
	ctx.lastcmd = nil
	ctx.cmdbox = nil
	ctx.Modes = nil
	ctx.Cache.Clean()
	ctx.SetPrompt(nil)
	return nil
}

// GetPrompt returns context prompt
func (ctx *Context) GetPrompt() string {
	return ctx.prompt
}

// SetPrompt sets the value for the context prompt.
func (ctx *Context) SetPrompt(newPrompt *string) {
	if newPrompt == nil {
		ctx.prompt = DEFAULTPROMPT
	} else {
		ctx.prompt = tools.String(newPrompt)
	}
}

// PushMode adds a new mode.
func (ctx *Context) PushMode(anchor *graph.Node) error {
	modeBox := &ModeBox{
		Prompt: ctx.GetPrompt(),
		Anchor: anchor,
		Mode:   ctx.lastcmd,
		CmdBox: ctx.cmdbox,
	}
	ctx.Modes = append(ctx.Modes, modeBox)
	ctx.SetPrompt(tools.PString(ctx.lastcmd.Prompt.(string)))
	return nil
}

// PopMode returns the last mode
func (ctx *Context) PopMode() *ModeBox {
	boxLen := len(ctx.Modes)
	if boxLen == 0 {
		return nil
	}
	result := ctx.Modes[boxLen-1]
	ctx.Modes = ctx.Modes[0 : boxLen-1]
	ctx.SetPrompt(tools.PString(result.Prompt))
	return result
}

// GetLastAnchor returns the last token matched before sink node.
func (ctx *Context) GetLastAnchor() *graph.Node {
	index := len(ctx.Matched) - 2
	tools.Debug("last token: %#v\n", ctx.Matched[index])
	cn := ctx.Matched[index]
	for _, child := range cn.Node.Children {
		if child.IsNext {
			tools.Debug("child anchor: %#v\n", child)
			return child
		}
	}
	return nil
}

// NewContext creates a new Context instance.
func NewContext(prefix *string) *Context {
	ctx := &Context{
		Cache:   NewCache(),
		process: NewContextProcess(),
	}
	ctx.SetPrompt(prefix)
	return ctx
}
