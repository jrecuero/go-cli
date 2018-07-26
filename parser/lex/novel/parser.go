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
	PIPE
	// ASTERISK token *. #7
	ASTERISK
	// PLUS token +. #8
	PLUS
	// QUESTION mark token ?. #9
	QUESTION
	// ADMIRATION mark token !. #10
	ADMIRATION
	// AT token @. #11
	AT
	// OPENMARK token <. #12
	OPENMARK
	// CLOSEMARK token >. #13
	CLOSEMARK
)

// cliCharMap represents the mapping between runes and codes.
var cliCharMap = map[rune]parser.Token{
	'[': OPENBRACKET,
	']': CLOSEBRACKET,
	'|': PIPE,
	'*': ASTERISK,
	'+': PLUS,
	'?': QUESTION,
	'!': ADMIRATION,
	'@': AT,
	'<': OPENMARK,
	'>': CLOSEMARK,
}

// Syntax represents the Novel command syntax.
type Syntax struct {
	Arguments []string
	Tokens    []parser.Token
}

// Parser represents the Novel parser.
type Parser struct {
	syntax  *Syntax
	charMap map[rune]parser.Token
}

// Parse implements the Novel parse functionality.
func (p *Parser) Parse(index int, token parser.Token, lit string) error {
	p.syntax.Arguments = append(p.syntax.Arguments, lit)
	p.syntax.Tokens = append(p.syntax.Tokens, token)
	return nil
}

// Result returns the parse result.
func (p *Parser) Result() interface{} {
	return p.syntax
}

// getIdentRunes returns special runes to be scanned as part of idents.
func (p *Parser) getIdentRunes() []rune {
	return []rune{'_', '-'}
}

// IsIdentRune returns if the rune can be part of an ident.
func (p *Parser) IsIdentRune(ch rune) bool {
	return parser.IsLetter(ch) || parser.IsDigit(ch) || tools.SearchKeyInRuneTable(p.getIdentRunes(), ch) == nil
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
		charMap: cliCharMap,
	}
}
