package novel

import (
	"github.com/jrecuero/go-cli/parser"
	"github.com/jrecuero/go-cli/tools"
)

const (
	// OPENBRACKET token [. #4
	OPENBRACKET parser.Token = iota + parser.CUSTOM
	// CLOSEBRACKET token ]. #5
	CLOSEBRACKET
	// PIPE token |. #6
)

// novelCharMap represents the mapping between runes and codes.
var novelCharMap = map[rune]parser.Token{
	'[': OPENBRACKET,
	']': CLOSEBRACKET,
}

// Syntax represents the Novel command syntax.
type Syntax struct {
	Idents []string
	Tokens []parser.Token
}

// Parser represents the Novel parser.
type Parser struct {
	syntax  *Syntax
	charMap map[rune]parser.Token
}

// Parse implements the Novel parse functionality.
func (p *Parser) Parse(index int, token parser.Token, lit string) error {
	p.syntax.Idents = append(p.syntax.Idents, lit)
	p.syntax.Tokens = append(p.syntax.Tokens, token)
	return nil
}

// Result returns the parse result.
func (p *Parser) Result() interface{} {
	return p.syntax
}

// getIdentRunes returns special runes to be scanned as part of idents.
func (p *Parser) getIdentRunes() []rune {
	return []rune{'_', '-', '<', '>'}
}

// IsIdentRune returns if the rune can be part of an ident.
func (p *Parser) IsIdentRune(ch rune) bool {
	return parser.IsLetter(ch) || parser.IsDigit(ch) || tools.SearchKeyInRuneTable(p.getIdentRunes(), ch) == nil
}

// IsIdentPrefixRune returns if the rune can be ident prefix.
func (p *Parser) IsIdentPrefixRune(ch rune) bool {
	return parser.IsLetter(ch) || ch == '<' || ch == '>'
}

// GetCharMap returns the mapping between runes to tokens.
func (p *Parser) GetCharMap() map[rune]parser.Token {
	return p.charMap
}

// Parser should implement ILexer interface.
var _ parser.ILexer = (*Parser)(nil)

// NewParser creates a new Novel Parser instance.
func NewParser() *Parser {
	return &Parser{
		syntax:  &Syntax{},
		charMap: novelCharMap,
	}
}
