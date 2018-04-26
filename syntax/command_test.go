package syntax_test

import (
	"testing"

	"github.com/jrecuero/go-cli/syntax"
)

// TestCommand_Command ensures the command structure works properly.
func TestCommand_Command(t *testing.T) {
	c := syntax.Command{
		Syntax: "test name",
		Help:   "Test command",
		Arguments: []syntax.Argument{
			{
				Label:   "name",
				Type:    "string",
				Default: "",
				Help:    "Name information",
			},
		},
	}

	if c.GetLabel() != "" {
		t.Errorf("GetLabel <Command> failed")
	}
	if c.GetHelp() != "Test command" {
		t.Errorf("GetHelp <Command> failed")
	}
}
