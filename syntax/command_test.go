package syntax_test

import (
	"testing"

	"github.com/jrecuero/go-cli/syntax"
)

// TestCommand_Command ensures the command structure works properly.
func TestCommand_Command(t *testing.T) {
	c := syntax.Command{
		Content: syntax.NewContent("", "Test command", nil).(*syntax.Content),
		Syntax:  "test name",
		Arguments: []*syntax.Argument{
			{
				Content: syntax.NewContent("name", "Name information", nil).(*syntax.Content),
				Type:    "string",
				Default: "",
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
