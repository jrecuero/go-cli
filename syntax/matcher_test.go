package syntax_test

import (
	"testing"

	"github.com/jrecuero/go-cli/syntax"
)

// TestMatcher_Matcher ensures the matcher structure works properly.
func TestMatcher_Matcher(t *testing.T) {
	m := syntax.NewMatcher(nil, nil)

	if m.Ctx != nil {
		t.Errorf("Context <Matcher> failed")
	}
	if m.G != nil {
		t.Errorf("Graph <Matcher> failed")
	}
}
