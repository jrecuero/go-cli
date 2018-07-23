package syntax

import (
	"errors"
	"fmt"

	"github.com/jrecuero/go-cli/graph"
	"github.com/jrecuero/go-cli/tools"
)

// NSManager represents the namespace manager.
// It is in charge to manage the active namespace, it will provides operation
// to settle all namespace commands properly and to find commands based on the
// command line input.
// It creates and handle the Command Tree and the Parsing Tree.
// The command tree contains all commands in a hierarchical way.
// The parsing tree expand the command tree with all command arguments. This is
// the tree used for parsing the command line.
type NSManager struct {
	nsname    string       // NameSpace Manager name.
	ns        *NameSpace   // NameSpace instance.
	ctx       *Context     // Context instance.
	matcher   *Matcher     // Matcher instance
	isup      bool         // Is NSManager up or down?
	cmdTree   *CommandTree // CommandTree instance
	parseTree *ParseTree   // ParseTree instance
	commands  []*Command   // Contains all commands that can be used for the NameSpace Manager.
}

// GetName returns the NameSpace Manager name.
func (nsm *NSManager) GetName() string {
	return nsm.nsname
}

// GetNameSpace returns the namespace related with the namager.
func (nsm *NSManager) GetNameSpace() *NameSpace {
	return nsm.ns
}

// GetCommands returns all commands registered to the NS manager.
func (nsm *NSManager) GetCommands() []*Command {
	return nsm.commands
}

// GetContext returns the manager context.
func (nsm *NSManager) GetContext() *Context {
	return nsm.ctx
}

// GetMatcher returns the manager matcher.
func (nsm *NSManager) GetMatcher() *Matcher {
	return nsm.matcher
}

// GetCommandTree returns the manager command tree instance.
func (nsm *NSManager) GetCommandTree() *CommandTree {
	return nsm.cmdTree
}

// GetParseTree returns the manager parse tree instance.
func (nsm *NSManager) GetParseTree() *ParseTree {
	return nsm.parseTree
}

// Setup initializes the namespace manager.
// It reads all commands for the NameSpace and will update the commands field
// with all of them. Command Tree and Parsing Tree are created at this time.
// Context and Marcher are initialized here.
func (nsm *NSManager) Setup() *NSManager {
	if nsm.ns == nil {
		return nil
	}
	for _, cmd := range nsm.ns.GetCommands() {
		nsm.commands = append(nsm.commands, cmd)
	}
	nsm.cmdTree = NewCommandTree()
	if err := nsm.CreateCommandTree(); err != nil {
		return nil
	}
	nsm.parseTree = NewParseTree()
	if err := nsm.CreateParseTree(nil); err != nil {
		return nil
	}
	nsm.ctx = NewContext(nil)
	nsm.matcher = NewMatcher(nsm.ctx, nsm.parseTree.Graph)
	return nsm
}

// Search searches for the given pattern in the commands map.
func (nsm *NSManager) Search(pattern string) ([]*Command, error) {
	return nil, errors.New("error")
}

// createBuiltinTree creates the builtin tree inside any mode.
func (nsm *NSManager) createBuiltinTree(cmd *Command, cmdNode *ContentNode) {
	for _, cmdBuiltin := range NewBuiltins() {
		if cmdBuiltin.IsMode() {
			continue
		}
		if cmd.GetLabel() != cmdBuiltin.GetLabel() {
			newHookNode := ContentNodeToNode(cmdNode)
			if cmdBuiltin.Parent == nil {
				nsm.cmdTree.AddTo(newHookNode, cmdBuiltin)
			} else {
				parent := nsm.cmdTree.SearchFlatUnder(newHookNode, cmdBuiltin.Parent)
				if parent != nil {
					nsm.cmdTree.AddTo(parent, cmdBuiltin)
				}
			}
		}
	}

}

// CreateCommandTree creates the command tree with all commands already stored
// in the manager.
func (nsm *NSManager) CreateCommandTree() error {
	for _, cmd := range nsm.commands {
		// Look for the command parent in the command tree.
		parent := nsm.cmdTree.SearchFlat(cmd.Parent)
		cmdNode := nsm.cmdTree.AddTo(parent, cmd)
		if cmd.IsMode() {
			nsm.createBuiltinTree(cmd, cmdNode)
		}
	}
	return nil
}

// CreateParseTree creates the parse tree with all commands already stored in
// the manager command tree.
func (nsm *NSManager) CreateParseTree(root *graph.Node) error {
	var parentCmd *Command
	var traverse *graph.Node
	if root == nil {
		traverse = nsm.cmdTree.Root
		parentCmd = nil
	} else {
		traverse = root
		parentCmd = root.Content.(*Command)
	}
	for _, node := range traverse.Children {
		cmd := node.Content.(*Command)
		// At this point there is enought information to identify if command
		// has subcommands (children), so the graphical tree can be properly
		// built for the command.
		//tools.Debug("SetupGraph")
		cmd.SetupGraph(len(node.Children) != 0)
		//tools.Debug("parseTree.Root: %p\n", nsm.parseTree.Root)
		//tools.Debug("Add Command to Parse Tree:\n\tparent: %#v\n\tcmd: %#v\n", parentCmd, cmd)
		nsm.parseTree.AddCommand(parentCmd, cmd)
		if err := nsm.CreateParseTree(node); err != nil {
			return fmt.Errorf("traverse children error: %v", err)
		}
	}
	return nil
}

// Execute executes the command for the given command line.
func (nsm *NSManager) Execute(line interface{}) (interface{}, bool) {
	return nsm.matcher.Execute(line)
}

// Complete returns possible complete string for command line being entered.
func (nsm *NSManager) Complete(line interface{}) (interface{}, bool) {
	return nsm.matcher.Complete(line)
}

// Help returns the help for a node if it is matched.
func (nsm *NSManager) Help(line interface{}) (interface{}, bool) {
	return nsm.matcher.Help(line)
}

// CompleteAndHelp returns possible complete string for command line being entered.
func (nsm *NSManager) CompleteAndHelp(line interface{}) (interface{}, bool) {
	return nsm.matcher.CompleteAndHelp(line)
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

// CreateNSManager creates and setups a new NSManager.
func CreateNSManager(namespace *NameSpace) (*NSManager, error) {
	nsm := NewNSManager(namespace)
	if nsm.Setup() == nil {
		return nil, tools.ERROR(nil, true, "Error Create NSManager")
	}
	return nsm, nil
}
