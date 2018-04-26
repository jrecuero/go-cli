package syntax_test

import (
	"testing"

	"github.com/jrecuero/go-cli/syntax"
)

// TestArgument_Argument ensures the argument structure works properly.
func TestArgument_Argument(t *testing.T) {
	a := syntax.Argument{
		Label:   "name",
		Type:    "string",
		Default: "",
		Help:    "Name information",
	}
	if a.GetLabel() != "name" {
		t.Errorf("GetLabel <Argument> failed")
	}
	if a.GetHelp() != "Name information" {
		t.Errorf("GetHelp <Argument> failed")
	}
}
