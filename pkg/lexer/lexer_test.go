package lexer

import (
	"testing"

	"github.com/uesteibar/lainoa/pkg/token"
	"gotest.tools/v3/assert"
)

func TestNextToken(t *testing.T) {
	input := `
		let ten = 10;
		let five = 5;

		fun add(x, y) {
			x + y;
		};

		let result = add(ten, five);
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.FUNCTION, "fun"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "ten"},
		{token.COMMA, ","},
		{token.IDENT, "five"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for _, et := range tests {
		tok := l.NextToken()

		assert.Equal(t, et.expectedType, tok.Type)
		assert.Equal(t, et.expectedLiteral, tok.Literal)
	}
}
