package syntax_test

import (
	"testing"

	"github.com/jrecuero/go-cli/syntax"
	"github.com/jrecuero/go-cli/tools"
)

// TestParseTree_NewParseTree ensures ParseTree works properly.
func TestParseTree_NewParseTree(t *testing.T) {
	pt := syntax.NewParseTree()
	if pt == nil {
		t.Errorf("new parse tree error: nil")
	}
}

// TestParseTree_AddCommand ensures ParseTree works properly.
func TestParseTree_AddCommand(t *testing.T) {
	tools.Log().Printf("TestParseTree:%s\n", "AddCommand")
	commands := []*syntax.Command{
		syntax.NewCommand(nil, "set", "Set test command", nil, nil).SetupGraph(false),
		syntax.NewCommand(nil, "get", "Get test command", nil, nil).SetupGraph(false),
	}
	pt := syntax.NewParseTree()
	for _, cmd := range commands {
		if err := pt.AddCommand(nil, cmd); err != nil {
			t.Errorf("add to command parse tree error: nil")
		}
	}
	for i, node := range pt.Root.Children {
		cn := syntax.NodeToContentNode(node)
		tools.Log().Printf("pt.Root.Children %d : %#v\n", i, cn)
		tools.Log().Printf("pt.Root.Children.Content %d : %#v\n", i, cn.GetContent())
		if len(cn.Children) != 0 {
			for _, child := range cn.Children {
				cnChild := syntax.NodeToContentNode(child)
				tools.Log().Printf("cn.Children %#v\n", cnChild)
				tools.Log().Printf("cn.Children.Content : %#v\n", cnChild.GetContent())
			}
		}
	}
}
