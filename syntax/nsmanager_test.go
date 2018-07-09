package syntax_test

import (
	"fmt"
	"testing"

	"github.com/jrecuero/go-cli/syntax"
	"github.com/jrecuero/go-cli/tools"
)

func createNameSpaceForTest() *syntax.NameSpace {
	//setCmd := syntax.NewCommand(nil, "set", "Set test help", nil, nil)
	//setCmd := &syntax.Command{
	//    Content: syntax.NewContent("set", "Set command", nil).(*syntax.Content),
	//    Syntax:  "set version",
	//    Arguments: []*syntax.Argument{
	//        syntax.NewArgument("version", "Version number", nil, "string", ""),
	//    },
	//}
	setCmd := syntax.NewCommand(nil, "set version", "Set test help",
		[]*syntax.Argument{
			syntax.NewArgument("version", "Version number", nil, "string", ""),
		}, nil)
	getCmd := syntax.NewCommand(nil, "get", "Get test help", nil, nil)
	setBoolCmd := syntax.NewCommand(setCmd, "bool", "Set Bool test help", nil, nil)
	setBaudrateCmd := syntax.NewCommand(setCmd, "baudrate [speed | parity]?", "Set baudrate help",
		[]*syntax.Argument{
			syntax.NewArgument("speed", "Baudrate speed", nil, "string", ""),
			syntax.NewArgument("parity", "Baudrate parity value", nil, "string", ""),
		}, nil)
	setSpeedCmd := syntax.NewCommand(setCmd, "speed", "Set Speed test help", nil, nil)
	setSpeedDeviceCmd := syntax.NewCommand(setSpeedCmd, "device name", "Set speed device help",
		[]*syntax.Argument{
			syntax.NewArgument("name", "Device name", nil, "string", ""),
		}, nil)
	getSpeedCmd := syntax.NewCommand(getCmd, "speed [device name | value]?", "Get speed help",
		[]*syntax.Argument{
			syntax.NewArgument("device", "Device", nil, "string", ""),
			syntax.NewArgument("name", "Device name", nil, "string", ""),
			syntax.NewArgument("value", "Speed value", nil, "string", ""),
		}, nil)
	setCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		fmt.Println("executint enter")
		return nil
	}
	commands := []*syntax.Command{
		setCmd,
		getCmd,
		syntax.NewCommand(nil, "config", "Config test help", nil, nil),
		setBaudrateCmd,
		setSpeedCmd,
		setBoolCmd,
		syntax.NewCommand(getCmd, "baudrate", "Get Baudrate test help", nil, nil),
		//syntax.NewCommand(getCmd, "speed", "Get Speed test help", nil, nil),
		getSpeedCmd,
		setSpeedDeviceCmd,
	}
	ns := syntax.NewNameSpace("test")
	for _, c := range commands {
		c.Setup()
		ns.Add(c)
	}
	return ns
}

// TestNSManager_NSManager ensures the namespace handler struct works properly.
//func TestNSManager_NSManager(t *testing.T) {
//    ns := createNameSpaceForTest()
//    nsm := syntax.NewNSManager(ns)
//    if nsm.GetNameSpace() == nil {
//        t.Errorf("create manager: namespace: nil")
//    }
//    if nsm.GetNameSpace().Name != "test" {
//        t.Errorf("create manager: namespace: name")
//    }
//    nsm.Setup()
//    if len(nsm.GetCommands()) != 1 {
//        t.Errorf("create manager: commands: len")
//    }
//    if len(tools.MapCast(nsm.GetCommands()["set"])) != 2 {
//        t.Errorf("create manager: commands: set command: len")
//    }
//    for k, v := range nsm.GetCommands() {
//        tools.Tester("%s  %#v\n\n", k, v)
//        for k2, v2 := range tools.MapCast(v) {
//            tools.Tester("%s  %#v\n\n", k2, v2)
//            for k3, v3 := range tools.MapCast(v2) {
//                tools.Tester("%s  %#v\n\n", k3, v3)
//            }
//        }
//    }
//}

// TestNSManager_Setup ensures the namespace handler struct works properly.
func TestNSManager_Setup(t *testing.T) {
	ns := createNameSpaceForTest()
	nsm := syntax.NewNSManager(ns)
	if err := nsm.Setup(); err == nil {
		t.Errorf("NSManager setup error: %v", err)
	}
	tools.Tester("Display Command Tree")
	for _, c := range nsm.GetCommandTree().Root.Children {
		tools.Tester("%#v\n", c)
	}
	//tools.Tester(nsm.GetCommandTree().ToMermaid())
	tools.Tester("Display Parse Tree")
	for _, c := range nsm.GetParseTree().Root.Children {
		tools.Tester("%#v\n", c)
	}
	//tools.Tester(nsm.GetParseTree().ToMermaid())

	ctx := syntax.NewContext()
	m := syntax.NewMatcher(ctx, nsm.GetParseTree().Graph)
	line := "set 1.0 speed device home"
	if _, ok := m.Match(line); !ok {
		t.Errorf("match return %#v for line: %s", ok, line)
	}
	//tools.Tester("%#v\n", ctx.GetLastCommand())
	//tools.Tester("-----------------------------------------")
	//for _, token := range ctx.Matched {
	//    c := token.Node.GetContent()
	//    tools.Tester("%v %v %#v %s\n", c.IsCommand(), c.IsArgument(), c.GetLabel(), token.Value)
	//}
	//tools.Tester("-----------------------------------------")
	//for i, cbox := range ctx.GetCommandBox() {
	//    tools.Tester("%d %#v\n", i, cbox)
	//}
	//tools.Tester("-----------------------------------------")
	//v, _ := ctx.GetArgValueForArgLabel(tools.PString("set"), "version")
	//tools.Tester("set version is %#v\n", v)
	//v, _ = ctx.GetArgValueForArgLabel(tools.PString("device"), "name")
	//tools.Tester("device name is %#v\n", v)
	//tools.Tester("-----------------------------------------")
}

// TestNSManager_Complete ensures the namespace handler struct works properly.
func TestNSManager_Complete(t *testing.T) {
	ns := createNameSpaceForTest()
	nsm := syntax.NewNSManager(ns)
	if err := nsm.Setup(); err == nil {
		t.Errorf("NSManager setup error: %v", err)
	}
	ctx := syntax.NewContext()
	m := syntax.NewMatcher(ctx, nsm.GetParseTree().Graph)
	//line := []string{"set", "1.0", "b"}
	line := "set 1.0 b"
	tools.Tester("line: %#v\n", line)
	result, _ := m.Complete(line)
	tools.Tester("%#v\n", result)
	if len(result.([]interface{})) != 2 || result.([]interface{})[0] != "baudrate" || result.([]interface{})[1] != "bool" {
		t.Errorf("complete didn't match: %#v line: %s", result, line)
	}
	//for i, token := range ctx.Matched {
	//    tools.Tester("%d %#v\n", i, token)
	//}
	//cn := ctx.Matched[1].Node
	//tools.Tester("%#v children: %#v\n", cn.GetContent().GetLabel(), cn.Children)
	line = "set 1.0 speed "
	tools.Tester("line: %#v\n", line)
	result, _ = m.Complete(line)
	tools.Tester("%#v\n", result)
	if len(result.([]interface{})) != 2 || result.([]interface{})[0] != "<<<_CR_>>>" || result.([]interface{})[1] != "device" {
		t.Errorf("complete didn't match: %#v line: %s", result, line)
	}
}

// TestNSManager_Help ensures the namespace handler struct works properly.
func TestNSManager_Help(t *testing.T) {
	ns := createNameSpaceForTest()
	nsm := syntax.NewNSManager(ns)
	if err := nsm.Setup(); err == nil {
		t.Errorf("NSManager setup error: %v", err)
	}
	ctx := syntax.NewContext()
	m := syntax.NewMatcher(ctx, nsm.GetParseTree().Graph)
	//line := []string{"set", "1.0", "b"}
	line := "set 1.0 baudrate "
	helps, _ := m.Help(line)
	tools.Tester("line: %#v\n", line)
	for _, h := range helps.([]interface{}) {
		tools.Tester("%#v\n", h.(string))
	}
	if len(helps.([]interface{})) != 3 {
		t.Errorf("help didn't match: %#v line: %s", helps, line)
	}
	line = "set "
	helps, _ = m.Help(line)
	tools.Tester("line: %#v\n", line)
	for _, h := range helps.([]interface{}) {
		tools.Tester("%#v\n", h.(string))
	}
	if len(helps.([]interface{})) != 1 {
		t.Errorf("help didn't match: %#v line: %s", helps, line)
	}
}

// TestNSManager_Execute_Enter ensures the namespace handler struct works properly.
func TestNSManager_Execute_Enter(t *testing.T) {
	ns := createNameSpaceForTest()
	nsm := syntax.NewNSManager(ns)
	if err := nsm.Setup(); err == nil {
		t.Errorf("NSManager setup error: %v", err)
	}
	tools.Tester("Display Command Tree")
	for _, c := range nsm.GetCommandTree().Root.Children {
		tools.Tester("%#v\n", c)
	}
	//tools.Tester(nsm.GetCommandTree().ToMermaid())
	tools.Tester("Display Parse Tree")
	for _, c := range nsm.GetParseTree().Root.Children {
		tools.Tester("%#v\n", c)
	}
	//tools.Tester(nsm.GetParseTree().ToMermaid())

	ctx := syntax.NewContext()
	m := syntax.NewMatcher(ctx, nsm.GetParseTree().Graph)
	line := "set 1.0 speed device home"
	if _, ok := m.Match(line); !ok {
		t.Errorf("match return %#v for line: %s", ok, line)
	} else {
		lastCommand := ctx.GetLastCommand()
		lastCommand.Enter(ctx, nil)
	}
}
