package lexer

import (
	"monkey-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
		let x = 1;
		let y = 243;

		let add = fn(a, b) {
			a + b;
		}

		let result = add(x + y);
		`

	tts := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "y"},
		{token.ASSIGN, "="},
		{token.INT, "243"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "a"},
		{token.COMMA, ","},
		{token.IDENT, "b"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "a"},
		{token.PLUS, "+"},
		{token.IDENT, "b"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tts {
		tk := l.NextToken()

		if tk.Type != tt.expectedType {
			t.Fatalf("tts[%d] - tokentype wrong. expected=%q, got=%q \n\n %q",
				i, tt.expectedType, tk.Type, l.nextPos)
		}

		if tk.Literal != tt.expectedLiteral {
			t.Fatalf("tts[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tk.Literal)
		}
	}
}
