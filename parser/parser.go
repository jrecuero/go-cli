package parser

import (
	"fmt"
	"io"
)

// Syntax represent a command syntax.
type Syntax struct {
	Command   string
	Arguments []string
	Tokens    []Token
}

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		// last read token
		tok Token
		// last read literal
		lit string
		// buffer size (max=1)
		n int
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// Parse parses a statement.
func (p *Parser) Parse() (*Syntax, error) {
	syntax := &Syntax{}

	tok, lit := p.scanIgnoreWhitespace()
	//fmt.Printf("1. tok is '%s'\n", lit)
	if tok != IDENT {
		return nil, fmt.Errorf("found %q, expected command", lit)
	}

	syntax.Command = lit

	// Next we should loop over all arguments
	for {
		// Read a field
		tok, lit = p.scanIgnoreWhitespace()
		//fmt.Printf("2. tok:%d, lit: '%s'\n", tok, lit)
		if tok == EOF {
			break
		}
		syntax.Arguments = append(syntax.Arguments, lit)
		syntax.Tokens = append(syntax.Tokens, tok)
	}
	return syntax, nil
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
