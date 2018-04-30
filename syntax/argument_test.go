package syntax_test

import (
	"testing"

	"github.com/jrecuero/go-cli/syntax"
)

// TestArgument_Argument ensures the argument structure works properly.
func TestArgument_Argument(t *testing.T) {
	a := syntax.Argument{
		Content: syntax.NewContent("name", "Name information", nil).(*syntax.Content),
		Type:    "string",
		Default: "",
	}
	if a.GetLabel() != "name" {
		t.Errorf("GetLabel <Argument> failed")
	}
	if a.GetHelp() != "Name information" {
		t.Errorf("GetHelp <Argument> failed")
	}
}
