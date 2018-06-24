package syntax_test

import (
	"reflect"
	"testing"

	"github.com/jrecuero/go-cli/syntax"
)

// TestCommandTree_NewCommandTree ensures CommandTree works properly.
func TestCommandTree_NewCommandTree(t *testing.T) {
	ct := syntax.NewCommandTree()
	if ct == nil {
		t.Errorf("new command tree error: nil")
	}
}

// TestCommandTree_AddTo ensures CommandTree works properly.
func TestCommandTree_AddTo(t *testing.T) {
	//var commands []*syntax.Command
	//commands := make([]*syntax.Command, 2)
	//commands[0] = syntax.NewCommand(nil, "set", "Set test command", nil)
	//commands[1] = syntax.NewCommand(nil, "get", "Get test command", nil)
	commands := []*syntax.Command{
		syntax.NewCommand(nil, "set", "Set test command", nil),
		syntax.NewCommand(nil, "get", "Get test command", nil),
		syntax.NewCommand(nil, "setup", "Setup test command", nil),
	}
	ct := syntax.NewCommandTree()
	var cn *syntax.ContentNode
	for _, c := range commands {
		cn = ct.AddTo(ct.Root, c)
		if cn == nil {
			t.Errorf("add to command tree error: nil")
		}
		got := cn.Content.(*syntax.Command)
		if !reflect.DeepEqual(c, got) {
			t.Errorf("add to command tree error:\n\texp: %#v\n\tgot: %#v\n", c, got)
		}
	}
}

// TestCommandTree_AddTo_Default ensures CommandTree works properly.
func TestCommandTree_AddTo_Default(t *testing.T) {
	commands := []*syntax.Command{
		syntax.NewCommand(nil, "set", "Set test command", nil),
		syntax.NewCommand(nil, "get", "Get test command", nil),
		syntax.NewCommand(nil, "setup", "Setup test command", nil),
	}
	ct := syntax.NewCommandTree()
	var cn *syntax.ContentNode
	for _, c := range commands {
		cn = ct.AddTo(nil, c)
		if cn == nil {
			t.Errorf("add to command tree error: nil")
		}
		got := cn.Content.(*syntax.Command)
		if !reflect.DeepEqual(c, got) {
			t.Errorf("add to command tree error:\n\texp: %#v\n\tgot: %#v\n", c, got)
		}
	}
}

// TestCommandTree_SearchDeep ensures CommandTree works properly.
func TestCommandTree_SearchDeep(t *testing.T) {
	commands := []*syntax.Command{
		syntax.NewCommand(nil, "set", "Set test command", nil),
		syntax.NewCommand(nil, "get", "Get test command", nil),
		syntax.NewCommand(nil, "setup", "Setup test command", nil),
	}
	ct := syntax.NewCommandTree()
	var cn *syntax.ContentNode
	for _, c := range commands {
		cn = ct.AddTo(ct.Root, c)
	}
	for _, c := range commands {
		cn = ct.SearchDeep(c)
		got := cn.Content.(*syntax.Command)
		if !reflect.DeepEqual(c, got) {
			t.Errorf("search deep error:\n\texp: %#v\n\tgot: %#v\n", c, got)
		}
	}

	cn = ct.SearchDeep(commands[0])
	deep := syntax.NewCommand(nil, "baudrate", "Baudrate test command", nil)
	deepNode := ct.AddTo(syntax.ContentNodeToNode(cn), deep)
	if deepNode == nil {
		t.Errorf("add to command tree error: nil")
	}
	cn = ct.SearchDeep(deep)
	got := cn.Content.(*syntax.Command)
	if !reflect.DeepEqual(deep, got) {
		t.Errorf("search deep error:\n\texp: %#v\n\tgot: %#v\n", deep, got)
	}
}

// TestCommandTree_SearchFlat ensures CommandTree works properly.
func TestCommandTree_SearchFlat(t *testing.T) {
	commands := []*syntax.Command{
		syntax.NewCommand(nil, "set", "Set test command", nil),
		syntax.NewCommand(nil, "get", "Get test command", nil),
		syntax.NewCommand(nil, "setup", "Setup test command", nil),
	}
	ct := syntax.NewCommandTree()
	var cn *syntax.ContentNode
	for _, c := range commands {
		cn = ct.AddTo(nil, c)
	}
	for _, c := range commands {
		cn = ct.SearchFlat(c)
		got := cn.Content.(*syntax.Command)
		if !reflect.DeepEqual(c, got) {
			t.Errorf("search deep error:\n\texp: %#v\n\tgot: %#v\n", c, got)
		}
	}

	deep := syntax.NewCommand(commands[0], "baudrate", "Baudrate test command", nil)
	cn = ct.SearchFlat(deep.Parent)
	deepNode := ct.AddTo(syntax.ContentNodeToNode(cn), deep)
	if deepNode == nil {
		t.Errorf("add to command tree error: nil")
	}
	cn = ct.SearchFlat(deep)
	//fmt.Printf("%#v\n", cn)
	//fmt.Printf("%#v\n", cn.Content)
	got := cn.Content.(*syntax.Command)
	if !reflect.DeepEqual(deep, got) {
		t.Errorf("search deep error:\n\texp: %#v\n\tgot: %#v\n", deep, got)
	}
}
