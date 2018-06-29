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
	//setBaudrateCmd :=syntax.NewCommand(setCmd, "baudrate", "Set Baudrate test help", nil, nil),
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
	getSpeedCmd := syntax.NewCommand(getCmd, "speed [device]?", "Get speed help",
		[]*syntax.Argument{
			syntax.NewArgument("device", "Device", nil, "string", ""),
		}, nil)
	commands := []*syntax.Command{
		setCmd,
		getCmd,
		syntax.NewCommand(nil, "config", "Config test help", nil, nil),
		setBaudrateCmd,
		setSpeedCmd,
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
//        tools.Log().Printf("%s  %#v\n\n", k, v)
//        for k2, v2 := range tools.MapCast(v) {
//            tools.Log().Printf("%s  %#v\n\n", k2, v2)
//            for k3, v3 := range tools.MapCast(v2) {
//                tools.Log().Printf("%s  %#v\n\n", k3, v3)
//            }
//        }
//    }
//}

// TestNSManager_Setup ensures the namespace handler struct works properly.
func TestNSManager_Setup(t *testing.T) {
	ns := createNameSpaceForTest()
	nsm := syntax.NewNSManager(ns)
	err := nsm.Setup()
	if err == nil {
		t.Errorf("NSManager setup error: %v", err)
	}
	tools.Log().Println("Display Command Tree")
	for _, c := range nsm.GetCommandTree().Root.Children {
		tools.Log().Printf("%#v\n", c)
	}
	fmt.Println(nsm.GetCommandTree().ToMermaid())
	tools.Log().Println("Display Parse Tree")
	for _, c := range nsm.GetParseTree().Root.Children {
		tools.Log().Printf("%#v\n", c)
	}
	fmt.Println(nsm.GetParseTree().ToMermaid())
}
