package syntax_test

import (
	"reflect"
	"testing"

	"github.com/jrecuero/go-cli/syntax"
	"github.com/jrecuero/go-cli/tools"
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
	//commands[0] = syntax.NewCommand(nil, "set", "Set test command", nil, nil)
	//commands[1] = syntax.NewCommand(nil, "get", "Get test command", nil, nil)
	commands := []*syntax.Command{
		syntax.NewCommand(nil, "set", "Set test command", nil, nil).SetupGraph(false),
		syntax.NewCommand(nil, "get", "Get test command", nil, nil).SetupGraph(false),
		syntax.NewCommand(nil, "setup", "Setup test command", nil, nil).SetupGraph(false),
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
		syntax.NewCommand(nil, "set", "Set test command", nil, nil).SetupGraph(false),
		syntax.NewCommand(nil, "get", "Get test command", nil, nil).SetupGraph(false),
		syntax.NewCommand(nil, "setup", "Setup test command", nil, nil).SetupGraph(false),
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

// TestCommandTree_AddTo_WithChildren ensures CommandTree works properly.
func TestCommandTree_AddTo_WithChildren(t *testing.T) {
	getCmd := syntax.NewCommand(nil, "get", "Get test command", nil, nil).SetupGraph(false)
	setCmd := syntax.NewCommand(nil, "set", "Set test command", nil, nil).SetupGraph(false)
	//speedCmd := syntax.NewCommand(setCmd, "baudrate", "Set baudrate test command", nil, nil)
	speedCmd := syntax.NewCommand(setCmd,
		"baudrate",
		"Set baudrate test command",
		[]*syntax.Argument{
			syntax.NewArgument("speedv", "Speed value", nil, "int", 0),
		},
		nil).SetupGraph(false)
	setupCmd := syntax.NewCommand(nil, "setup", "Setup test command", nil, nil).SetupGraph(false)
	commands := []*syntax.Command{
		getCmd,
		setCmd,
		speedCmd,
		setupCmd,
	}
	ct := syntax.NewCommandTree()
	var cn *syntax.ContentNode
	for _, c := range commands {
		cn = ct.AddTo(nil, c)
	}
	for _, c := range commands {
		cn = ct.SearchFlatToContentNode(c)
		got := cn.Content.(*syntax.Command)

		tools.Tester("\n")
		tools.Tester("%#v\n", cn)
		tools.Tester("%#v\n", got)
		tools.Tester("%#v\n", got.Content)
		tools.Tester("%#v\n", got.CmdSyntax)

		if !reflect.DeepEqual(c, got) {
			t.Errorf("add to with children error:\n\texp: %#v\n\tgot: %#v\n", c, got)
		}
	}
}

//func SetEnter(ctx *syntax.Context, arguments interface{}) error {
//    tools.Tester(">>>>> Set Command Enter\n")
//    return nil
//}

// TestCommandTree_AddTo_WithCallback ensures CommandTree works properly.
func TestCommandTree_AddTo_WithCallback(t *testing.T) {
	enterCbs := []bool{false, false, false}
	expEnterCbs := []bool{false, true, false}
	exitCbs := []bool{false, false, false}
	expExitCbs := []bool{true, true, true}
	setCmd := syntax.NewCommand(nil, "set", "Set test command", nil, nil).SetupGraph(false)
	//setCmd.Callback.Enter = SetEnter
	setCmd.Callback.Exit = func(ctx *syntax.Context) error {
		exitCbs[0] = true
		return nil
	}
	getCmd := syntax.NewCommand(nil,
		"get",
		"Get test command",
		nil,
		syntax.NewDefaultCallback().SetEnter(func(ctx *syntax.Context, arguments interface{}) error {
			enterCbs[1] = true
			return nil
		}).SetExit(func(ctx *syntax.Context) error {
			exitCbs[1] = true
			return nil
		})).SetupGraph(false)
	setupCmd := syntax.NewCommand(nil,
		"setup",
		"Setup test command",
		nil,
		syntax.NewCallback(nil,
			nil,
			func(ctx *syntax.Context) error {
				exitCbs[2] = true
				return nil
			},
			nil)).SetupGraph(false)
	commands := []*syntax.Command{
		getCmd,
		setCmd,
		setupCmd,
	}
	ct := syntax.NewCommandTree()
	var cn *syntax.ContentNode
	for _, c := range commands {
		cn = ct.AddTo(nil, c)
	}
	for _, c := range commands {
		cn = ct.SearchFlatToContentNode(c)
		got := cn.Content.(*syntax.Command)
		tools.Tester("%#v\n", got)
		got.Enter(nil, nil)
		got.Exit(nil)
		if !reflect.DeepEqual(c, got) {
			t.Errorf("add to with callback error:\n\texp: %#v\n\tgot: %#v\n", c, got)
		}
	}
	if !reflect.DeepEqual(expEnterCbs, enterCbs) {
		t.Errorf("add to with callback error: enter callbacks:\n\texp: %#v\n\tgot: %#v\n", expEnterCbs, enterCbs)
	}
	if !reflect.DeepEqual(expExitCbs, exitCbs) {
		t.Errorf("add to with callback error: exit callbacks:\n\texp: %#v\n\tgot: %#v\n", expExitCbs, exitCbs)
	}
}

// TestCommandTree_SearchDeep ensures CommandTree works properly.
func TestCommandTree_SearchDeep(t *testing.T) {
	commands := []*syntax.Command{
		syntax.NewCommand(nil, "set", "Set test command", nil, nil).SetupGraph(false),
		syntax.NewCommand(nil, "get", "Get test command", nil, nil).SetupGraph(false),
		syntax.NewCommand(nil, "setup", "Setup test command", nil, nil).SetupGraph(false),
	}
	ct := syntax.NewCommandTree()
	var cn *syntax.ContentNode
	for _, c := range commands {
		cn = ct.AddTo(ct.Root, c)
	}
	for _, c := range commands {
		cn = ct.SearchDeepToContentNode(c)
		got := cn.Content.(*syntax.Command)
		if !reflect.DeepEqual(c, got) {
			t.Errorf("search deep error:\n\texp: %#v\n\tgot: %#v\n", c, got)
		}
	}

	cn = ct.SearchDeepToContentNode(commands[0])
	deep := syntax.NewCommand(nil, "baudrate", "Baudrate test command", nil, nil).SetupGraph(false)
	deepNode := ct.AddTo(syntax.ContentNodeToNode(cn), deep)
	if deepNode == nil {
		t.Errorf("add to command tree error: nil")
	}
	cn = ct.SearchDeepToContentNode(deep)
	got := cn.Content.(*syntax.Command)
	if !reflect.DeepEqual(deep, got) {
		t.Errorf("search deep error:\n\texp: %#v\n\tgot: %#v\n", deep, got)
	}
}

// TestCommandTree_SearchFlat ensures CommandTree works properly.
func TestCommandTree_SearchFlat(t *testing.T) {
	commands := []*syntax.Command{
		syntax.NewCommand(nil, "set", "Set test command", nil, nil).SetupGraph(false),
		syntax.NewCommand(nil, "get", "Get test command", nil, nil).SetupGraph(false),
		syntax.NewCommand(nil, "setup", "Setup test command", nil, nil).SetupGraph(false),
	}
	ct := syntax.NewCommandTree()
	var cn *syntax.ContentNode
	for _, c := range commands {
		cn = ct.AddTo(nil, c)
	}
	for _, c := range commands {
		cn = ct.SearchFlatToContentNode(c)
		got := cn.Content.(*syntax.Command)
		if !reflect.DeepEqual(c, got) {
			t.Errorf("search deep error:\n\texp: %#v\n\tgot: %#v\n", c, got)
		}
	}

	deep := syntax.NewCommand(commands[0], "baudrate", "Baudrate test command", nil, nil).SetupGraph(false)
	cn = ct.SearchFlatToContentNode(deep.Parent)
	deepNode := ct.AddTo(syntax.ContentNodeToNode(cn), deep)
	if deepNode == nil {
		t.Errorf("add to command tree error: nil")
	}
	cn = ct.SearchFlatToContentNode(deep)
	tools.Tester("%#v\n", cn)
	tools.Tester("%#v\n", cn.Content)
	got := cn.Content.(*syntax.Command)
	if !reflect.DeepEqual(deep, got) {
		t.Errorf("search deep error:\n\texp: %#v\n\tgot: %#v\n", deep, got)
	}
}
