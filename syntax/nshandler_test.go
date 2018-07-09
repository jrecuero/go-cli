package syntax_test

import (
	"testing"

	"github.com/jrecuero/go-cli/syntax"
	"github.com/jrecuero/go-cli/tools"
)

func setup(nsname string) (*syntax.NSHandler, *syntax.NSManager, *syntax.NameSpace) {
	setCmd := syntax.NewCommand(nil, "set version", "Set test help",
		[]*syntax.Argument{
			syntax.NewArgument("version", "Version number", nil, "string", ""),
		}, nil)
	setCmd.Callback.Enter = func(ctx *syntax.Context, arguments interface{}) error {
		version, _ := ctx.GetArgValueForArgLabel(nil, "version")
		tools.Tester("executing enter with version:", version)
		args, _ := ctx.GetArgValuesForCommandLabel(nil)
		tools.Tester("argumets:", args)
		params := arguments.(map[string]interface{})
		tools.Tester("version:", params["version"])
		return nil
	}
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
	if nsh, err := syntax.CreateNSHandler(nsname, commands); err == nil {
		return nsh, nsh.GetActive().NSMgr, nsh.GetActive().NS
	}
	return nil, nil, nil
}

// TestNSHandler_NSHandler ensures the namespace handler struct works properly.
func TestNSHandler_NSHandler(t *testing.T) {
	nsh := syntax.NewNSHandler()

	if len(nsh.GetNameSpaces()) != 0 {
		t.Errorf("namespaces: len")
	}

	_, ok := nsh.FindNameSpace("test")
	if ok == nil {
		t.Errorf("found namespace")
	}

	var ns *syntax.NameSpace
	ns, ok = nsh.CreateNameSpace("test")
	if ok != nil {
		t.Errorf("create namespace: return")
	}
	if ns == nil {
		t.Errorf("create namespace: ns")
	}
	newNS, ok := nsh.FindNameSpace("test")
	if ok != nil {
		t.Errorf("find amespace: not found")
	}
	if newNS != ns {
		t.Errorf("find namespace: different")
	}

	delNS, ok := nsh.DeleteNameSpace("test")
	if ok != nil {
		t.Errorf("delete amespace: not found")
	}
	if delNS != ns {
		t.Errorf("delete namespace: different")
	}

	c := &syntax.Command{}
	_, ok = nsh.CreateNameSpace("test1")
	ok = nsh.RegisterCommandToNameSpace("test1", c)
	if ok != nil {
		t.Errorf("register command: failed")
	}
	ok = nsh.UnregisterCommandFromNameSpace("test1", c)
	if ok != nil {
		t.Errorf("unregister command: failed")
	}
}

// TestNSHandler_Activate ensures activating namespaces works fine.
func TestNSHandler_Activate(t *testing.T) {
	nsh := syntax.NewNSHandler()
	ns1, ok := nsh.CreateNameSpace("test1")

	ans, ok := nsh.ActivateNameSpace("test1")
	if ok != nil {
		t.Errorf("activate namespace: failed")
	}
	if ans != ns1 {
		t.Errorf("activate namespace: different")
	}
	if nsh.GetActive() == nil {
		t.Errorf("activate namespace: active nil")
	}
	if nsh.GetActive().Name != "test1" {
		t.Errorf("activate namespace: active name")
	}
	if nsh.GetActive().NS != ns1 {
		t.Errorf("activate namespace: active ns")
	}

	dns, ok := nsh.DeactivateNameSpace("test1")
	if ok != nil {
		t.Errorf("deactivate namespace: failed")
	}
	if dns != ns1 {
		t.Errorf("deactivate namespace: different")
	}
	if nsh.GetActive() != nil {
		t.Errorf("deactivate namespace: active nil")
	}
}

// TestNSHandler_Switch ensures switching namespaces works fine.
func TestNSHandler_Switch(t *testing.T) {
	nsh := syntax.NewNSHandler()
	ns1, ok := nsh.CreateNameSpace("test1")
	ns2, ok := nsh.CreateNameSpace("test2")
	_, ok = nsh.ActivateNameSpace("test1")
	sns1, sns2, ok := nsh.SwitchToNameSpace("test2")
	if ok != nil {
		t.Errorf("switch namespace: failed")
	}
	if sns1 != ns1 {
		t.Errorf("switch namespace: old different")
	}
	if sns2 != ns2 {
		t.Errorf("switch namespace: new different")
	}
	if nsh.GetActive() == nil {
		t.Errorf("switch namespace: active nil")
	}
	if nsh.GetActive().Name != "test2" {
		t.Errorf("switch namespace: active name")
	}
	if nsh.GetActive().NS != ns2 {
		t.Errorf("switch namespace: active ns")
	}
	if len(nsh.GetStack()) != 1 {
		t.Errorf("switch namespace: stack size")
	}

	bns2, bns1, ok := nsh.SwitchBackToNameSpace()
	if ok != nil {
		t.Errorf("switch-back namespace: failed")
	}
	if bns1 != ns1 {
		t.Errorf("switch-back namespace: new different")
	}
	if bns2 != ns2 {
		t.Errorf("switch-back namespace: old different")
	}
	if nsh.GetActive() == nil {
		t.Errorf("switch-back namespace: active nil")
	}
	if nsh.GetActive().Name != "test1" {
		t.Errorf("switch-back namespace: active name: %s <> %s", nsh.GetActive().Name, "test1")
	}
	if nsh.GetActive().NS != ns1 {
		t.Errorf("switch-back namespace: active ns")
	}
	if len(nsh.GetStack()) != 0 {
		t.Errorf("switch-back namespace: stack size")
	}
}

// TestNSHandler_Setup ensures switching namespaces works fine.
func TestNSHandler_Setup(t *testing.T) {
	nsh, nsm, ns := setup("test")
	tools.Tester("Handler  : %#v\n", nsh)
	tools.Tester("Manager  : %#v\n", nsm)
	tools.Tester("NameSpace: %#v\n", ns)
	ctx := syntax.NewContext()
	m := syntax.NewMatcher(ctx, nsm.GetParseTree().Graph)
	line := "set 1.0 speed device home"
	if _, ok := m.Match(line); !ok {
		t.Errorf("match return %#v for line: %s", ok, line)
	}
}

// TestNSHandler_Execute_Enter ensures switching namespaces works fine.
func TestNSHandler_Execute_Enter(t *testing.T) {
	nsh, nsm, ns := setup("test")
	tools.Tester("Handler  : %#v\n", nsh)
	tools.Tester("Manager  : %#v\n", nsm)
	tools.Tester("NameSpace: %#v\n", ns)
	ctx := syntax.NewContext()
	m := syntax.NewMatcher(ctx, nsm.GetParseTree().Graph)
	line := "set 1.0 speed device home"
	if _, ok := m.Match(line); !ok {
		t.Errorf("match return %#v for line: %s", ok, line)
	}
	//line = "set 1.0"
	//if _, ok := m.Match(line); !ok {
	//    t.Errorf("match return %#v for line: %s", ok, line)
	//} else {
	//    lastCommand := ctx.GetLastCommand()
	//    lastCommand.Enter(ctx, nil)
	//}
	m.Execute("set 1.0")
}
