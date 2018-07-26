package cli_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/jrecuero/go-cli/parser"
	lexcli "github.com/jrecuero/go-cli/parser/lex/cli"
	"github.com/jrecuero/go-cli/tools"
)

// Ensure the parser can parse string into Syntax
func TestParser_ParseSyntax(t *testing.T) {
	var tests = []struct {
		s      string
		syntax *lexcli.Syntax
		err    string
	}{
		{
			s: "SELECT name",
			syntax: &lexcli.Syntax{
				Command:   "SELECT",
				Arguments: []string{"name"},
				Tokens:    []parser.Token{parser.IDENT},
			},
		},
		{
			s: "SELECT fname lname",
			syntax: &lexcli.Syntax{
				Command:   "SELECT",
				Arguments: []string{"fname", "lname"},
				Tokens:    []parser.Token{parser.IDENT, parser.IDENT},
			},
		},
		{
			s: "SELECT name [ age ]",
			syntax: &lexcli.Syntax{
				Command:   "SELECT",
				Arguments: []string{"name", "[", "age", "]"},
				Tokens: []parser.Token{parser.IDENT, lexcli.OPENBRACKET,
					parser.IDENT, lexcli.CLOSEBRACKET},
			},
		},
		{
			s: "SELECT name age [ id | passport ]",
			syntax: &lexcli.Syntax{
				Command:   "SELECT",
				Arguments: []string{"name", "age", "[", "id", "|", "passport", "]"},
				Tokens: []parser.Token{parser.IDENT, parser.IDENT,
					lexcli.OPENBRACKET, parser.IDENT,
					lexcli.PIPE, parser.IDENT, lexcli.CLOSEBRACKET},
			},
		},
		{
			s: "SELECT name [ age ]?",
			syntax: &lexcli.Syntax{
				Command:   "SELECT",
				Arguments: []string{"name", "[", "age", "]", "?"},
				Tokens: []parser.Token{parser.IDENT, lexcli.OPENBRACKET, parser.IDENT,
					lexcli.CLOSEBRACKET, lexcli.QUESTION},
			},
		},
		{
			s: "SELECT <name>",
			syntax: &lexcli.Syntax{
				Command:   "SELECT",
				Arguments: []string{"<", "name", ">"},
				Tokens:    []parser.Token{lexcli.OPENMARK, parser.IDENT, lexcli.CLOSEMARK},
			},
		},

		// Errors
		{s: "1one", err: `found "1", illegal token`, syntax: nil},
		{s: "SELECT 2", err: `found "2", illegal token`, syntax: nil},
	}

	for i, tt := range tests {
		syntax, err := parser.NewParser(strings.NewReader(tt.s), lexcli.NewParser()).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q error mismatch:\n exp=%s\n got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.syntax, syntax) {
			t.Errorf("%d. %q\n\nsyntax mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.syntax, syntax)
		} else {
			if syntax != nil {
				tools.Log().Printf("command: %s arguments: %s tokens: %d\n",
					tt.syntax.Command, tt.syntax.Arguments, tt.syntax.Tokens)
			}
		}
	}
}

// errstring returns the string representation of an error
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
