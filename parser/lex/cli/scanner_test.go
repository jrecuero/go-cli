package cli_test

import (
	"strings"
	"testing"

	"github.com/jrecuero/go-cli/parser"
	lexcli "github.com/jrecuero/go-cli/parser/lex/cli"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok parser.Token
		lit string
	}{
		// Special tokens
		{s: "", tok: parser.EOF},
		{s: "#", tok: parser.ILLEGAL, lit: "#"},
		{s: " ", tok: parser.WS, lit: " "},
		{s: "\t", tok: parser.WS, lit: "\t"},
		{s: "\n", tok: parser.WS, lit: "\n"},

		// Misc characters
		{s: "[", tok: lexcli.OPENBRACKET, lit: "["},
		{s: "]", tok: lexcli.CLOSEBRACKET, lit: "]"},
		{s: "|", tok: lexcli.PIPE, lit: "|"},
		{s: "*", tok: lexcli.ASTERISK, lit: "*"},
		{s: "+", tok: lexcli.PLUS, lit: "+"},
		{s: "?", tok: lexcli.QUESTION, lit: "?"},
		{s: "!", tok: lexcli.ADMIRATION, lit: "!"},
		{s: "@", tok: lexcli.AT, lit: "@"},
		{s: "<", tok: lexcli.OPENMARK, lit: "<"},
		{s: ">", tok: lexcli.CLOSEMARK, lit: ">"},

		// Identifiers
		{s: "foo", tok: parser.IDENT, lit: "foo"},
	}

	for i, tt := range tests {
		s := parser.NewScanner(strings.NewReader(tt.s), lexcli.NewParser())
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got %q", i, tt.s, tt.lit, lit)
		}
	}
}
