package syntax

import (
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
	tools.Log().Printf("ParseTree:AddCommand parent: %#v command: %#v\n", parent, cmd)
	var hook *graph.Node
	if parent == nil {
		hook = pt.Root
	} else {
		hook = parent.CmdSyntax.Graph.Next
	}
	cmdGraph := cmd.CmdSyntax.Graph
	hook.AddChild(cmdGraph.Root)
	return nil
}

// NewParseTree creates a new ParseTree instance.
func NewParseTree() *ParseTree {
	return &ParseTree{
		graph.NewGraph(nil),
	}
}
