package lexer

import (
	"monkey-interpreter/token"
)

type Lexer struct {
	input   string
	pos     int
	nextPos int
	ch      byte
}

func New(i string) *Lexer {
	l := &Lexer{input: i}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tk token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tk = newToken(token.ASSIGN, l.ch)
	case '+':
		tk = newToken(token.PLUS, l.ch)
	case ',':
		tk = newToken(token.COMMA, l.ch)
	case ';':
		tk = newToken(token.SEMICOLON, l.ch)
	case '(':
		tk = newToken(token.LPAREN, l.ch)
	case ')':
		tk = newToken(token.RPAREN, l.ch)
	case '{':
		tk = newToken(token.LBRACE, l.ch)
	case '}':
		tk = newToken(token.RBRACE, l.ch)
	case 0:
		tk.Literal = ""
		tk.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tk.Literal = l.readIdentifier()
			tk.Type = token.LookUpIdent(tk.Literal)
			return tk
		} else if isDigit(l.ch) {
			tk.Literal = l.readInteger()
			tk.Type = token.INT
			return tk
		} else {
			tk = newToken(token.ILLEGAL, l.ch)
		}
	}

	// Lexer ch pointer is advanced
	l.readChar()

	return tk
}

func (l *Lexer) readChar() {
	if l.nextPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPos]
	}
	l.pos = l.nextPos
	l.nextPos += 1
}

func (l *Lexer) readIdentifier() string {
	pos := l.pos
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) readInteger() string {
	pos := l.pos
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}
