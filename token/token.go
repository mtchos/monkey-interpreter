package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// IDENT INT Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// ASSIGN PLUS Operators
	ASSIGN = "="
	PLUS   = "+"

	// COMMA SEMICOLON Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	// LPAREN RPAREN LBRACE RBRACE Delimiters
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// FUNCTION LET Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
