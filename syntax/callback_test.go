package syntax_test

import (
	"testing"

	"github.com/jrecuero/go-cli/syntax"
)

// TestCallback_Callback ensures the argument structure works properly.
func TestCallback_Callback(t *testing.T) {
	c := &syntax.Callback{}

	if c == nil {
		t.Errorf("Callback <Callback> failed")
	}
}
