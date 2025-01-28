package lexer

import (
	"monkey-interpreter/token"
)

type Lexer struct {
	input   string
	currPos int
	nextPos int
	ch      byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tk token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			str := string(ch) + string(l.ch)
			tk = newToken(token.EQ, str)
		} else {
			tk = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			str := string(ch) + string(l.ch)
			tk = newToken(token.NEQ, str)
		} else {
			tk = newToken(token.BANG, l.ch)
		}
	case '+':
		tk = newToken(token.PLUS, l.ch)
	case '-':
		tk = newToken(token.MINUS, l.ch)
	case '*':
		tk = newToken(token.ASTERISK, l.ch)
	case '/':
		tk = newToken(token.SLASH, l.ch)
	case '<':
		tk = newToken(token.LT, l.ch)
	case '>':
		tk = newToken(token.GT, l.ch)
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

	// Lexer ch pointer is moved forward
	l.readChar()

	return tk
}

func (l *Lexer) peekChar() byte {
	if l.nextPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPos]
	}
}

func (l *Lexer) readChar() {
	if l.nextPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPos]
	}

	l.currPos = l.nextPos
	l.nextPos += 1
}

func (l *Lexer) readIdentifier() string {
	pos := l.currPos

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.currPos]
}

func (l *Lexer) readInteger() string {
	pos := l.currPos

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.currPos]
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

func newToken[T byte | string](tkType token.TokenType, lit T) token.Token {
	return token.Token{Type: tkType, Literal: string(lit)}
}
