package parser

import (
	"bufio"
	"bytes"
	"io"
)

// Scanner represents a lexical scanner.
type Scanner struct {
	r       *bufio.Reader
	charmap map[rune]Token
	lexer   ILexer
}

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	// If we see a digit then consume as a number
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if IsLetter(ch) {
		s.unread()
		return s.scanIdent()
	} else if ch == eof {
		return EOF, ""
	}

	if code, ok := s.charmap[ch]; ok {
		return code, string(ch)
	}

	return ILLEGAL, string(ch)
}

// scanWhitespace consumes the current rune an all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
			//} else if !IsLetter(ch) && !IsDigit(ch) && ch != '_' && ch != '-' {
		} else if !s.lexer.IsIdentRune(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return IDENT, buf.String()
}

// read reads the next rune from the buffered reader.
// Returns the rune(0) if an error occurs.
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

//unread places the previously read rune on the reader.
func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

// isWhitespace returns true if the rune is a space, tab or newline.
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

// IsLetter returns true if the rune is a letter.
func IsLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// IsDigit returns true if the rune is a digit.
func IsDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

// eof represents a marker rune for the end of the reader.
var eof = rune(0)

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader, lexer ILexer) *Scanner {
	scan := &Scanner{
		r:       bufio.NewReader(r),
		lexer:   lexer,
		charmap: lexer.GetCharMap(),
	}
	return scan
}
