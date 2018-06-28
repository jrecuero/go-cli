package syntax_test

import (
	"testing"

	"github.com/jrecuero/go-cli/syntax"
	"github.com/jrecuero/go-cli/tools"
)

func createNameSpaceForTest() *syntax.NameSpace {
	commands := []*syntax.Command{
		syntax.NewCommand(nil, "set", "Set test help", nil, nil).SetupGraph(false),
		syntax.NewCommand(nil, "get", "Get test help", nil, nil).SetupGraph(false),
		syntax.NewCommand(nil, "config", "Config test help", nil, nil).SetupGraph(false),
	}
	ns := syntax.NewNameSpace("test")
	for _, c := range commands {
		c.Setup()
		ns.Add(c)
	}
	return ns
}

// TestNSManager_NSManager ensures the namespace handler struct works properly.
func TestNSManager_NSManager(t *testing.T) {
	//ns := createNameSpaceForTest()
	//nsm := syntax.NewNSManager(ns)
	//if nsm.GetNameSpace() == nil {
	//    t.Errorf("create manager: namespace: nil")
	//}
	//if nsm.GetNameSpace().Name != "test" {
	//    t.Errorf("create manager: namespace: name")
	//}
	//nsm.Setup()
	//if len(nsm.GetCommands()) != 1 {
	//    t.Errorf("create manager: commands: len")
	//}
	//if len(tools.MapCast(nsm.GetCommands()["set"])) != 2 {
	//    t.Errorf("create manager: commands: set command: len")
	//}
	//for k, v := range nsm.GetCommands() {
	//    tools.Log().Printf("%s  %#v\n\n", k, v)
	//    for k2, v2 := range tools.MapCast(v) {
	//        tools.Log().Printf("%s  %#v\n\n", k2, v2)
	//        for k3, v3 := range tools.MapCast(v2) {
	//            tools.Log().Printf("%s  %#v\n\n", k3, v3)
	//        }
	//    }
	//}
}

// TestNSManager_CreateCommandTree ensures the namespace handler struct works properly.
func TestNSManager_CreateCommandTree(t *testing.T) {
	ns := createNameSpaceForTest()
	nsm := syntax.NewNSManager(ns)
	err := nsm.Setup()
	if err == nil {
		t.Errorf("NSManager setup error: %v", err)
	}
	for _, c := range nsm.GetCommandTree().Root.Children {
		tools.Log().Printf("%#v\n", c)
	}
}
