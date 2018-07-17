package syntax_test

import (
	"testing"

	"github.com/jrecuero/go-cli/graph"
	"github.com/jrecuero/go-cli/syntax"
	"github.com/jrecuero/go-cli/tools"
)

// TestMatcher_NewMatcher ensures the matcher structure works properly.
func TestMatcher_NewMatcher(t *testing.T) {
	m := syntax.NewMatcher(nil, graph.NewGraph(nil))
	if m.Ctx != nil {
		t.Errorf("Context <Matcher> failed")
	}
	if m.Grapher == nil {
		t.Errorf("Graph <Matcher> failed")
	}
}

// TestMatcher_Matcher ensures the matcher structure works properly.
func TestMatcher_Matcher(t *testing.T) {
	command := &syntax.Command{
		Content: syntax.NewContent("", "User command", nil).(*syntax.Content),
		Syntax:  "SELECT name age",
		Arguments: []*syntax.Argument{
			syntax.NewArgument("name", "", nil, "string", "", nil),
			syntax.NewArgument("age", "Age information", nil, "int", 0, nil),
		},
	}
	command.Setup()
	cs := command.CmdSyntax
	tools.Tester("%s", cs.Graph.ToString())
	m := syntax.NewMatcher(syntax.NewContext(), cs.Graph)
	//line := []string{"SELECT", "name", "age"}
	line := "SELECT name age"
	m.Match(line)
}
