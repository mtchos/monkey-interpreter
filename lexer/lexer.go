package lexer

import "monkey-interpreter/token"

type Lexer struct {
	input string
	curr  int
	next  int
	char  byte
}

func New(i string) *Lexer {
	l := &Lexer{input: i}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.char {
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '=':
		tok = newToken(token.ASSIGN, l.char)
	case '+':
		tok = newToken(token.PLUS, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func (l *Lexer) readChar() {
	if l.next >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.next]
	}
	l.curr = l.next
	l.next += 1
}
