package syntax_test

import (
	"testing"

	"github.com/jrecuero/go-cli/syntax"
)

// TestContext_Context ensures the context structure works properly.
func TestContext_Context(t *testing.T) {
	c := syntax.NewContext(nil)

	if c.Matched != nil {
		t.Errorf("Matched <Context> failed")
	}
}
