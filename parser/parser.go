package parser

import (
	"fmt"
	"io"
)

// ILexer represents any lexer interface.
type ILexer interface {
	Parse(index int, token Token, char string) error
	Result() interface{}
	GetCharMap() map[rune]Token
	IsIdentRune(ch rune) bool
	IsIdentPrefixRune(ch rune) bool
}

// Parser represents a parser.
type Parser struct {
	lexer ILexer
	s     *Scanner
	buf   struct {
		// last read token
		tok Token
		// last read literal
		lit string
		// buffer size (max=1)
		n int
	}
}

// Parse parses a statement.
func (p *Parser) Parse() (interface{}, error) {
	index := 1
	for {
		// Read a field
		tok, lit := p.scanIgnoreWhitespace()

		//tools.Debug("tok:%d, lit: '%s'\n", tok, lit)
		if tok == ILLEGAL {
			return nil, fmt.Errorf("found %q, illegal token", lit)
		} else if tok == EOF {
			break
		}
		p.lexer.Parse(index, tok, lit)
		index++
	}
	return p.lexer.Result(), nil
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanner then read that instead
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// scanIgnoreWhitespace scan the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	if tok, lit = p.scan(); tok == WS {
		tok, lit = p.scan()
	}
	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() {
	p.buf.n = 1
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader, lexer ILexer) *Parser {
	return &Parser{
		s:     NewScanner(r, lexer),
		lexer: lexer,
	}
}
