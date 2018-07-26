package novel_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/jrecuero/go-cli/parser"
	lexnovel "github.com/jrecuero/go-cli/parser/lex/novel"
	"github.com/jrecuero/go-cli/tools"
)

// Ensure the parser can parse string into Syntax
func TestParser_ParseSyntax(t *testing.T) {
	var tests = []struct {
		s      string
		syntax *lexnovel.Syntax
		err    string
	}{
		{
			s: "ACTOR hit TARGET",
			syntax: &lexnovel.Syntax{
				Idents: []string{"ACTOR", "hit", "TARGET"},
				Tokens: []parser.Token{parser.IDENT, parser.IDENT, parser.IDENT},
			},
		},
		{
			s: "ACTOR hit <self>",
			syntax: &lexnovel.Syntax{
				Idents: []string{"ACTOR", "hit", "<self>"},
				Tokens: []parser.Token{parser.IDENT, parser.IDENT, parser.IDENT},
			},
		},
		{
			s: "ACTOR hit TARGET flags",
			syntax: &lexnovel.Syntax{
				Idents: []string{"ACTOR", "hit", "TARGET", "flags"},
				Tokens: []parser.Token{parser.IDENT, parser.IDENT, parser.IDENT, parser.IDENT},
			},
		},
		{
			s: "[A1 A2] hit TARGET",
			syntax: &lexnovel.Syntax{
				Idents: []string{"[", "A1", "A2", "]", "hit", "TARGET"},
				Tokens: []parser.Token{lexnovel.OPENBRACKET, parser.IDENT, parser.IDENT, lexnovel.CLOSEBRACKET, parser.IDENT, parser.IDENT},
			},
		},
		{
			s: "[A1 A2] hit <self>",
			syntax: &lexnovel.Syntax{
				Idents: []string{"[", "A1", "A2", "]", "hit", "<self>"},
				Tokens: []parser.Token{lexnovel.OPENBRACKET, parser.IDENT, parser.IDENT, lexnovel.CLOSEBRACKET, parser.IDENT, parser.IDENT},
			},
		},
		{
			s: "[A1 A2] hit <self> flags",
			syntax: &lexnovel.Syntax{
				Idents: []string{"[", "A1", "A2", "]", "hit", "<self>", "flags"},
				Tokens: []parser.Token{lexnovel.OPENBRACKET, parser.IDENT, parser.IDENT, lexnovel.CLOSEBRACKET, parser.IDENT, parser.IDENT, parser.IDENT},
			},
		},
		{
			s: "ACTOR hit [T1 T2]",
			syntax: &lexnovel.Syntax{
				Idents: []string{"ACTOR", "hit", "[", "T1", "T2", "]"},
				Tokens: []parser.Token{parser.IDENT, parser.IDENT, lexnovel.OPENBRACKET, parser.IDENT, parser.IDENT, lexnovel.CLOSEBRACKET},
			},
		},
		{
			s: "ACTOR hit [T1 T2] flags",
			syntax: &lexnovel.Syntax{
				Idents: []string{"ACTOR", "hit", "[", "T1", "T2", "]", "flags"},
				Tokens: []parser.Token{parser.IDENT, parser.IDENT, lexnovel.OPENBRACKET, parser.IDENT, parser.IDENT, lexnovel.CLOSEBRACKET, parser.IDENT},
			},
		},
		{
			s: "[A1 A2] hit [T1 T2]",
			syntax: &lexnovel.Syntax{
				Idents: []string{"[", "A1", "A2", "]", "hit", "[", "T1", "T2", "]"},
				Tokens: []parser.Token{lexnovel.OPENBRACKET, parser.IDENT, parser.IDENT, lexnovel.CLOSEBRACKET, parser.IDENT, lexnovel.OPENBRACKET, parser.IDENT, parser.IDENT, lexnovel.CLOSEBRACKET},
			},
		},
		{
			s: "[A1 A2] hit [T1 T2] flags",
			syntax: &lexnovel.Syntax{
				Idents: []string{"[", "A1", "A2", "]", "hit", "[", "T1", "T2", "]", "flags"},
				Tokens: []parser.Token{lexnovel.OPENBRACKET, parser.IDENT, parser.IDENT, lexnovel.CLOSEBRACKET, parser.IDENT, lexnovel.OPENBRACKET, parser.IDENT, parser.IDENT, lexnovel.CLOSEBRACKET, parser.IDENT},
			},
		},
		// Errors
		{s: "1one", err: `found "1", illegal token`, syntax: nil},
		{s: "SELECT 2", err: `found "2", illegal token`, syntax: nil},
	}

	for i, tt := range tests {
		syntax, err := parser.NewParser(strings.NewReader(tt.s), lexnovel.NewParser()).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q error mismatch:\n exp=%s\n got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.syntax, syntax) {
			t.Errorf("%d. %q\n\nsyntax mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.syntax, syntax)
		} else {
			if syntax != nil {
				tools.Log().Printf("idents: %s tokens: %d\n", tt.syntax.Idents, tt.syntax.Tokens)
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
