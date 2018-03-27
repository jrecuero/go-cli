package parser_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/jrecuero/go-cli/parser"
)

// Ensure the parser can parse string into Syntax
func TestParser_ParseSyntax(t *testing.T) {
	var tests = []struct {
		s      string
		syntax *parser.Syntax
		err    string
	}{
		{
			s: "SELECT name",
			syntax: &parser.Syntax{
				Command:   "SELECT",
				Arguments: []string{"name"},
				Tokens:    []parser.Token{parser.IDENT},
			},
		},
		{
			s: "SELECT fname lname",
			syntax: &parser.Syntax{
				Command:   "SELECT",
				Arguments: []string{"fname", "lname"},
				Tokens:    []parser.Token{parser.IDENT, parser.IDENT},
			},
		},
		{
			s: "SELECT name [ age ]",
			syntax: &parser.Syntax{
				Command:   "SELECT",
				Arguments: []string{"name", "[", "age", "]"},
				Tokens: []parser.Token{parser.IDENT, parser.OPENBRACKET,
					parser.IDENT, parser.CLOSEBRACKET},
			},
		},
		{
			s: "SELECT name age [ id | passport ]",
			syntax: &parser.Syntax{
				Command:   "SELECT",
				Arguments: []string{"name", "age", "[", "id", "|", "passport", "]"},
				Tokens: []parser.Token{parser.IDENT, parser.IDENT,
					parser.OPENBRACKET, parser.IDENT,
					parser.PIPE, parser.IDENT, parser.CLOSEBRACKET},
			},
		},
		{
			s: "SELECT name [ age ]?",
			syntax: &parser.Syntax{
				Command:   "SELECT",
				Arguments: []string{"name", "[", "age", "]", "?"},
				Tokens: []parser.Token{parser.IDENT, parser.OPENBRACKET, parser.IDENT,
					parser.CLOSEBRACKET, parser.QUESTION},
			},
		},
	}

	for i, tt := range tests {
		syntax, err := parser.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q error mismatch:\n exp=%s\n got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.syntax, syntax) {
			t.Errorf("%d. %q\n\nsyntax mismatch:\n\nexp%#v\n\ngot=%#v\n\n", i, tt.s, tt.syntax, syntax)
		}
		fmt.Printf("command: %s arguments: %s tokens: %d\n",
			tt.syntax.Command, tt.syntax.Arguments, tt.syntax.Tokens)
	}
}

// errstring returns the string representation of an error
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
