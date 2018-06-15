package syntax_test

import (
	"fmt"
	"testing"

	"github.com/jrecuero/go-cli/syntax"
)

// TestMatcher_NewMatcher ensures the matcher structure works properly.
func TestMatcher_NewMatcher(t *testing.T) {
	m := syntax.NewMatcher(nil, nil)
	if m.Ctx != nil {
		t.Errorf("Context <Matcher> failed")
	}
	if m.G != nil {
		t.Errorf("Graph <Matcher> failed")
	}
}

// TestMatcher_Matcher ensures the matcher structure works properly.
func TestMatcher_Matcher(t *testing.T) {
	cs := syntax.NewCommandSyntax("SELECT name age")
	cs.CreateGraph(&syntax.Command{})
	fmt.Printf("%s", cs.Graph.ToString())
	m := syntax.NewMatcher(syntax.NewContext(), cs.Graph)
	line := []string{"SELECT", "name", "age"}
	m.MatchCommandLine(line)
}
