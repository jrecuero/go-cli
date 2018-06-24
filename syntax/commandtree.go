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
func (ct *CommandTree) searchDeep(node *graph.Node, c *Command) *graph.Node {
	for _, child := range node.Children {
		childCmd := child.Content.(*Command)
		//fmt.Printf("\tchild: %v\n\tcommand: %#v\n", childCmd, c)
		if reflect.DeepEqual(childCmd, c) {
			return child
		}
		retChild := ct.searchDeep(child, c)
		if retChild != nil {
			return retChild
		}
	}
	return nil
}

// searchFlaat searches for a command searching siblings first.
func (ct *CommandTree) searchFlat(node *graph.Node, c *Command) *graph.Node {
	for _, child := range node.Children {
		childCmd := child.Content.(*Command)
		if reflect.DeepEqual(childCmd, c) {
			return child
		}
	}
	for _, child := range node.Children {
		retChild := ct.searchFlat(child, c)
		if retChild != nil {
			return retChild
		}
	}
	return nil
}

// SearchDeep looks for the given command in the Command Tree.
func (ct *CommandTree) SearchDeep(c *Command) *ContentNode {
	return NodeToContentNode(ct.searchDeep(ct.Root, c))
}

// SearchFlat looks for the given command in the Command Tree.
func (ct *CommandTree) SearchFlat(c *Command) *ContentNode {
	return NodeToContentNode(ct.searchFlat(ct.Root, c))
}

// AddTo adds a new command to the command tree.
func (ct *CommandTree) AddTo(parent *graph.Node, c *Command) *ContentNode {
	if parent == nil {
		parent = ct.Root
	}
	node := NewContentNode(c.GetLabel(), c)
	parent.AddChild(ContentNodeToNode(node))
	return node
}

// NewCommandTree creates a new CommandTree instance.
func NewCommandTree() *CommandTree {
	return &CommandTree{
		graph.NewGraph(nil),
	}
}
