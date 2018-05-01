package syntax_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jrecuero/go-cli/syntax"
	"github.com/jrecuero/go-cli/tools"
)

var commands []*syntax.Command

func quickCommandAllPrefixes(pattern string) *syntax.Command {
	cmd := &syntax.Command{
		Content: syntax.NewContent("", fmt.Sprintf("Help for %s", pattern), nil).(*syntax.Content),
		Syntax:  pattern,
	}
	seq := strings.Split(pattern, " ")
	for _, prefix := range seq[1:] {
		cmd.Prefixes = append(cmd.Prefixes, &syntax.Prefix{
			Content: syntax.NewContent(prefix, fmt.Sprintf("Help for %s", pattern), nil).(*syntax.Content),
			Type:    "string",
			Default: prefix,
		})
	}
	return cmd
}

func setupCommands() {
	commands = []*syntax.Command{
		quickCommandAllPrefixes("set speed device"),
		quickCommandAllPrefixes("set speed aux"),
		quickCommandAllPrefixes("set baudrate main"),
	}
	//commands = []*syntax.Command{
	//    {
	//        Content: syntax.NewContent("", "Set speed device command", nil).(*syntax.Content),
	//        Syntax:  "set speed device",
	//        Prefixes: []*syntax.Prefix{
	//            {
	//                Content: syntax.NewContent("speed", "Set the speed", nil).(*syntax.Content),
	//                Type:    "string",
	//                Default: "speed",
	//            },
	//            {
	//                Content: syntax.NewContent("device", "Set device", nil).(*syntax.Content),
	//                Type:    "string",
	//                Default: "device",
	//            },
	//        },
	//    },
	//    {
	//        Content: syntax.NewContent("", "Set speed aux command", nil).(*syntax.Content),
	//        Syntax:  "set speed aux",
	//        Prefixes: []*syntax.Prefix{
	//            {
	//                Content: syntax.NewContent("speed", "Set the speed", nil).(*syntax.Content),
	//                Type:    "string",
	//                Default: "speed",
	//            },
	//            {
	//                Content: syntax.NewContent("aux", "Set aux", nil).(*syntax.Content),
	//                Type:    "string",
	//                Default: "aux",
	//            },
	//        },
	//    },
	//    {
	//        Content: syntax.NewContent("", "Set main baudrate", nil).(*syntax.Content),
	//        Syntax:  "set baudrate main",
	//        Prefixes: []*syntax.Prefix{
	//            {
	//                Content: syntax.NewContent("baudrate", "Set the baudrate", nil).(*syntax.Content),
	//                Type:    "string",
	//                Default: "baudrate",
	//            },
	//            {
	//                Content: syntax.NewContent("main", "Set main", nil).(*syntax.Content),
	//                Type:    "string",
	//                Default: "main",
	//            },
	//        },
	//    },
	//}
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
