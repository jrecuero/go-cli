package syntax_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jrecuero/go-cli/syntax"
	"github.com/jrecuero/go-cli/tools"
)

var commands []*syntax.Command

func buildQuickCommand(pattern string) *syntax.Command {
	cmd := &syntax.Command{
		Content: syntax.NewContent("", fmt.Sprintf("Help for %s", pattern), nil).(*syntax.Content),
		Syntax:  pattern,
	}
	seq := strings.Split(pattern, " ")
	return cmd
}

func setupCommands() {
	commands = []*syntax.Command{
		buildQuickCommand("set speed device"),
		buildQuickCommand("set speed aux"),
		buildQuickCommand("set baudrate main"),
	}
}

func createNameSpaceForTest() *syntax.NameSpace {
	setupCommands()
	ns := syntax.NewNameSpace("test")
	for _, c := range commands {
		c.Setup()
		ns.Add(c)
	}
	return ns
}

// TestNSManager_NSManager ensures the namespace handler struct works properly.
func TestNSManager_NSManager(t *testing.T) {
	ns := createNameSpaceForTest()
	m := syntax.NewNSManager(ns)
	if m.GetNameSpace() == nil {
		t.Errorf("create manager: namespace: nil")
	}
	if m.GetNameSpace().Name != "test" {
		t.Errorf("create manager: namespace: name")
	}
	m.Setup()
	if len(m.GetCommands()) != 1 {
		t.Errorf("create manager: commands: len")
	}
	if len(tools.MapCast(m.GetCommands()["set"])) != 2 {
		t.Errorf("create manager: commands: set command: len")
	}
	for k, v := range m.GetCommands() {
		fmt.Printf("%s  %#v\n\n", k, v)
		for k2, v2 := range tools.MapCast(v) {
			fmt.Printf("%s  %#v\n\n", k2, v2)
			for k3, v3 := range tools.MapCast(v2) {
				fmt.Printf("%s  %#v\n\n", k3, v3)
			}
		}
	}
}
