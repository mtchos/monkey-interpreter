package token

type TokenType string

type Token struct {
	Literal string
	Type    TokenType
}

const (
	// ILLEGAL EOF Special Types
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// IDENT INT Identifiers + Literals
	IDENT = "IDENT"
	INT   = "INT"

	// ASSIGN Assignment Operators
	ASSIGN = "="

	// BANG EQ NEQ Logical Operators
	BANG = "!"
	EQ   = "=="
	NEQ  = "!="

	// PLUS MINUS ASTERISK SLASH LT GT Mathematical Operators
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	// COMMA SEMICOLON Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	// LPAREN RPAREN LBRACE RBRACE Delimiters
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// LET TRUE FALSE IF ELSE FUNCTION RETURN Keywords
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	FUNCTION = "FUNCTION"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"fn":     FUNCTION,
	"return": RETURN,
}

func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
