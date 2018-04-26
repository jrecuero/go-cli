package syntax_test

import (
	"testing"

	"github.com/jrecuero/go-cli/syntax"
)

// TestCompleter_Completer ensures the completer structure works properly.
func TestCompleter_Completer(t *testing.T) {
	content := syntax.NewContent("test", "test help", nil)
	c := syntax.NewCompleter(content)

	if c.GetLabel() != "test" {
		t.Errorf("GetLabel <Completer> failed")
	}
	if c.GetContent() != content {
		t.Errorf("GetContent <Completer> failed")
	}
}

// TestCompleter_Command ensures the completer command structure works properly.
func TestCompleter_Command(t *testing.T) {
	content := syntax.NewContent("cmd", "test command help", nil)
	var c *syntax.CompleterCommand
	c = syntax.NewCompleterCommand(content)

	if c.GetLabel() != "cmd" {
		t.Errorf("GetLabel <Command> failed")
	}
}

// TestCompleter_Ident ensures the completer ident structure works properly.
func TestCompleter_Ident(t *testing.T) {
	content := syntax.NewContent("test ident", "test ident help", nil)
	c := syntax.NewCompleterIdent(content)

	if c.GetLabel() != "test ident" {
		t.Errorf("GetLabel <Ident> failed")
	}
	if c.GetContent() != content {
		t.Errorf("GetContent <Ident> failed")
	}
}
