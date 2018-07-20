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

	// OPENBRACKET token [. #4
	OPENBRACKET
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
