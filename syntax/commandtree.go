package syntax

import (
	"reflect"

	"github.com/jrecuero/go-cli/graph"
)

// CommandTree represents a command tree.
type CommandTree struct {
	*graph.Graph
}

// searchDeep searches for a command searching children first.
func (ct *CommandTree) searchDeep(node *graph.Node, cmd *Command) *graph.Node {
	if cmd != nil {
		for _, child := range node.Children {
			childCmd := child.Content.(*Command)
			//tools.Tracer("\tchild: %v\n\tcommand: %#v\n", childCmd, cmd)
			if reflect.DeepEqual(childCmd, cmd) {
				return child
			}
			retChild := ct.searchDeep(child, cmd)
			if retChild != nil {
				return retChild
			}
		}
	}
	return nil
}

// searchFlat searches for a command searching siblings first.
func (ct *CommandTree) searchFlat(node *graph.Node, cmd *Command) *graph.Node {
	if cmd != nil {
		for _, child := range node.Children {
			childCmd := child.Content.(*Command)
			if reflect.DeepEqual(childCmd, cmd) {
				return child
			}
		}
		for _, child := range node.Children {
			retChild := ct.searchFlat(child, cmd)
			if retChild != nil {
				return retChild
			}
		}
	}
	return nil
}

// SearchDeep looks for the given command in the Command Tree.
func (ct *CommandTree) SearchDeep(cmd *Command) *graph.Node {
	return ct.searchDeep(ct.Root, cmd)
}

// SearchFlat looks for the given command in the Command Tree.
func (ct *CommandTree) SearchFlat(cmd *Command) *graph.Node {
	return ct.searchFlat(ct.Root, cmd)
}

// SearchFlatUnder looks for the given command in the Command Tree.
func (ct *CommandTree) SearchFlatUnder(root *graph.Node, cmd *Command) *graph.Node {
	return ct.searchFlat(root, cmd)
}

// SearchDeepToContentNode looks for the given command in the Command Tree.
func (ct *CommandTree) SearchDeepToContentNode(cmd *Command) *ContentNode {
	return NodeToContentNode(ct.SearchDeep(cmd))
}

// SearchFlatToContentNode looks for the given command in the Command Tree.
func (ct *CommandTree) SearchFlatToContentNode(cmd *Command) *ContentNode {
	return NodeToContentNode(ct.SearchFlat(cmd))
}

// AddTo adds a new command to the command tree.
func (ct *CommandTree) AddTo(parent *graph.Node, cmd *Command) *ContentNode {
	if parent == nil {
		parent = ct.Root
	}
	node := NewContentNode(cmd.GetLabel(), cmd)
	parent.AddChild(ContentNodeToNode(node))
	return node
}

// NewCommandTree creates a new CommandTree instance.
func NewCommandTree() *CommandTree {
	return &CommandTree{
		graph.NewGraph(nil),
	}
}
