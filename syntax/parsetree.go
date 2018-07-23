package syntax

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jrecuero/go-cli/graph"
	"github.com/jrecuero/go-cli/tools"
)

// ParseTree represents the command parsing tree.
// The command parsing tree contains all commands splitted in tokens, where
// every token is an input for the command in the command line.
type ParseTree struct {
	*graph.Graph
}

// AddCommand adds a new command to the command parsing tree.
func (pt *ParseTree) AddCommand(parent *Command, cmd *Command) error {
	tools.Tracer("parent: %v | command: %v\n", parent, cmd)
	var hook *graph.Node
	if parent == nil {
		hook = pt.Root
	} else {
		hook = parent.CmdSyntax.Graph.Next
	}
	cmdGraph := cmd.CmdSyntax.Graph
	//tools.Debug("\tparent: %#v\n\tcommand: %#v\n\tcmdGraph: %#v\n", parent, cmd, cmdGraph)
	hook.AddChild(cmdGraph.Root)
	//tools.Debug("Hook: %#p\n", hook)
	//tools.Debug("pt.Root: %#p\n", pt.Root)
	return nil
}

// Explore implements a mechanism to interactibily explore the graph.
func (pt *ParseTree) Explore() {
	reader := bufio.NewReader(os.Stdin)
	traverse := pt.Root
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

// NewParseTree creates a new ParseTree instance.
func NewParseTree() *ParseTree {
	setupG := &graph.SetupGraph{
		RootContent: NewContentJoint("Root", "Root content", NewCompleterJoint("root")),
	}
	return &ParseTree{
		graph.NewGraph(setupG),
	}
}
