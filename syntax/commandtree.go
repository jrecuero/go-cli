package syntax

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/jrecuero/go-cli/graph"
	"github.com/jrecuero/go-cli/tools"
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

// Explore implements a mechanism to interactibily explore the graph.
func (ct *CommandTree) Explore() {
	reader := bufio.NewReader(os.Stdin)
	traverse := ct.Root
	parents := []*graph.Node{}
	var index int
	for {
		tools.ToDisplay(fmt.Sprintf("\n\nNode Information: %s\n", NodeToContentNode(traverse).ToContent()))
		tools.ToDisplay(fmt.Sprintf("Nbr of children: %d\n", len(traverse.Children)))
		if len(traverse.Children) > 0 {
			for i, child := range traverse.Children {
				//tools.ToDisplay("\t> %d %s\n", i, child.Label)
				tools.ToDisplay("\t> %d %s\n", i, NodeToContentNode(child).GetContent().GetLabel())
			}
			tools.ToDisplay("\n[0-%d] Select children", len(traverse.Children)-1)
		}
		tools.ToDisplay("\n[-] Select Parent\n[x] Exit\nSelect: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "x" {
			break
		} else if text == "-" {
			traverse = parents[len(parents)-1]
			parents = parents[:len(parents)-1]
			tools.ToDisplay("Parent selected %s\n", traverse.Label)
		} else {
			index, _ = strconv.Atoi(text)
			parents = append(parents, traverse)
			traverse = traverse.Children[index]
			tools.ToDisplay("Children selected %d - %s\n", index, traverse.Label)
		}
	}
}

// NewCommandTree creates a new CommandTree instance.
func NewCommandTree() *CommandTree {
	return &CommandTree{
		graph.NewGraph(nil),
	}
}
