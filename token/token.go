package token

type TokenType string

type Token struct {
	Literal string
	Type    TokenType
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

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
