package parser

// Token represents a lexical token
type Token int

const (
	// ILLEGAL token
	ILLEGAL Token = iota
	// EOF token
	EOF
	// WS token
	WS

	// IDENT represents command and argument tokens.
	IDENT

	// OPENBRACKET token [
	OPENBRACKET
	// CLOSEBRACKET token ]
	CLOSEBRACKET
	// PIPE token |
	PIPE
	// ASTERISK token *
	ASTERISK
	// PLUS token +
	PLUS
	// QUESTION mark token ?
	QUESTION
	// ADMIRATION mark token !
	ADMIRATION
	// AT token @
	AT
	// OPENMARK token <
	OPENMARK
	// CLOSEMARK token >
	CLOSEMARK
)
