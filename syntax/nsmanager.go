package syntax

import (
	"errors"
)

// NSManager represents the namespace manager.
// It is in charge to manage the active namespace, it will provides operation
// to settle all namespace commands properly and to find commands based on the
// command line input.
type NSManager struct {
	nsname    string       // NameSpace Manager name.
	ns        *NameSpace   // NameSpace instance.
	ctx       *Context     // Context instance.
	isup      bool         // Is NSManager up or down?
	cmdTree   *CommandTree // CommandTree instance
	parseTree *ParseTree   // ParseTree instance
	commands  []*Command   // Contains all commands that can be used for the NameSpace Manager.
}

// NewNSManager creates a new NSManager instance.
func NewNSManager(namespace *NameSpace) *NSManager {
	if namespace == nil {
		return &NSManager{}
	}
	return &NSManager{
		nsname: namespace.Name,
		ns:     namespace,
	}
}

// GetName returns the NameSpace Manager name.
func (nsm *NSManager) GetName() string {
	return nsm.nsname
}

// GetNameSpace returns the namespace related with the namager.
func (nsm *NSManager) GetNameSpace() *NameSpace {
	return nsm.ns
}

// GetCommands returns the manager commands.
func (nsm *NSManager) GetCommands() []*Command {
	return nsm.commands
}

// GetContext returns the manager context.
func (nsm *NSManager) GetContext() *Context {
	return nsm.ctx
}

// GetCommandTree returns the manager command tree instance.
func (nsm *NSManager) GetCommandTree() *CommandTree {
	return nsm.cmdTree
}

// GetParseTree returns the manager parse tree instance.
func (nsm *NSManager) GetParseTree() *ParseTree {
	return nsm.parseTree
}

// add inserts a new command in the internal command map.
func (nsm *NSManager) add(table map[string]interface{}, name []string, value interface{}) error {
	label := name[0]
	if table[label] == nil {
		if len(name) == 1 {
			table[label] = value
		} else {
			table[label] = make(map[string]interface{})
			nsm.add(table[label].(map[string]interface{}), name[1:], value)
		}
	} else {
		switch v := table[label].(type) {
		case map[string]interface{}:
			if len(name) == 1 {
				return errors.New("too short")
			}
			nsm.add(v, name[1:], value)
		default:
			return errors.New("error")
		}
	}
	return nil
}

// Setup initializes the namespace manager.
// It reads all commands for the NameSpace and will update the commands field
// with all of them.
func (nsm *NSManager) Setup() *NSManager {
	if nsm.ns == nil {
		return nil
	}
	//nsm.commands = make(map[string]interface{})
	for _, cmd := range nsm.ns.GetCommands() {
		//cmdSeq := strings.Split(cmd.FullCmd, " ")
		//err := nsm.add(nsm.commands, cmdSeq, cmd)
		nsm.commands = append(nsm.commands, cmd)
		//if err != nil {
		//    return err
		//}
	}
	nsm.CreateCommandTree()
	return nsm
}

// Search searches for the given pattern in the commands map.
func (nsm *NSManager) Search(pattern string) ([]*Command, error) {
	////sequence := strings.Split(pattern, " ")
	////return searchPatternInTable(nsm.commands, sequence)
	//locals, ok := tools.SearchPatternInMap(nsm.commands, pattern)
	//var results []*Command
	//for _, v := range locals {
	//    results = append(results, v.(*Command))
	//}
	//return results, ok
	return nil, errors.New("error")
}

// CreateCommandTree creates the command tree with all commands already stored
// in the manager.
func (nsm *NSManager) CreateCommandTree() error {
	nsm.cmdTree = NewCommandTree()
	for _, cmd := range nsm.commands {
		nsm.cmdTree.AddTo(nil, cmd)
	}
	return nil
}

// CreateParseTree creates the parse tree with all commands already stored in
// the manager command tree.
func (nsm *NSManager) CreateParseTree() error {
	return nil
}
