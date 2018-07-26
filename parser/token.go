package parser

// Token represents a lexical token
type Token int

const (
	// ILLEGAL token. #0
	ILLEGAL Token = iota
	// EOF token. #1
	EOF
	// WS token. #2
	WS

	// IDENT represents command and argument tokens. #3
	IDENT

	// CUSTOM represents the last token for the parse framwork. #4
	CUSTOM
)
