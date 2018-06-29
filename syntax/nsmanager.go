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

// Setup initializes the namespace manager.
// It reads all commands for the NameSpace and will update the commands field
// with all of them.
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
	return nsm
}

// Search searches for the given pattern in the commands map.
func (nsm *NSManager) Search(pattern string) ([]*Command, error) {
	return nil, errors.New("error")
}

// CreateCommandTree creates the command tree with all commands already stored
// in the manager.
func (nsm *NSManager) CreateCommandTree() error {
	for _, cmd := range nsm.commands {
		// Look for the command parent in the command tree.
		parent := nsm.cmdTree.SearchFlat(cmd.Parent)
		nsm.cmdTree.AddTo(parent, cmd)
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
		//tools.Log().Println("SetupGraph")
		cmd.SetupGraph(len(node.Children) != 0)
		//tools.Log().Printf("parseTree.Root: %p\n", nsm.parseTree.Root)
		tools.Log().Printf("Add Command to Parse Tree:\n\tparent: %#v\n\tcmd: %#v\n", parentCmd, cmd)
		nsm.parseTree.AddCommand(parentCmd, cmd)
		if err := nsm.CreateParseTree(node); err != nil {
			return fmt.Errorf("traverse children error: %v", err)
		}
	}
	return nil
}
