package lexer

import (
	"testing"

	"github.com/uesteibar/lainoa/pkg/token"
)

func TestNextToken(t *testing.T) {
	input := `
		let ten = 10;
		let five = 5;

		fun add(x, y) {
			x + y;
		}

		let result = add(ten, five);

		!-/*5;
		5 < 10 > 5;

		if (5 < 10) {
			return true;
		} else {
			return false;
		}
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"}, {token.IDENT, "ten"}, {token.ASSIGN, "="}, {token.INT, "10"}, {token.SEMICOLON, ";"},

		{token.LET, "let"}, {token.IDENT, "five"}, {token.ASSIGN, "="}, {token.INT, "5"}, {token.SEMICOLON, ";"},

		{token.FUNCTION, "fun"}, {token.IDENT, "add"}, {token.LPAREN, "("}, {token.IDENT, "x"}, {token.COMMA, ","}, {token.IDENT, "y"}, {token.RPAREN, ")"}, {token.LBRACE, "{"},
		{token.IDENT, "x"}, {token.PLUS, "+"}, {token.IDENT, "y"}, {token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.LET, "let"}, {token.IDENT, "result"}, {token.ASSIGN, "="}, {token.IDENT, "add"}, {token.LPAREN, "("}, {token.IDENT, "ten"}, {token.COMMA, ","}, {token.IDENT, "five"}, {token.RPAREN, ")"}, {token.SEMICOLON, ";"},

		{token.BANG, "!"}, {token.MINUS, "-"}, {token.SLASH, "/"}, {token.ASTERISK, "*"}, {token.INT, "5"}, {token.SEMICOLON, ";"},

		{token.INT, "5"}, {token.LT, "<"}, {token.INT, "10"}, {token.BT, ">"}, {token.INT, "5"}, {token.SEMICOLON, ";"},

		{token.IF, "if"}, {token.LPAREN, "("}, {token.INT, "5"}, {token.LT, "<"}, {token.INT, "10"}, {token.RPAREN, ")"}, {token.LBRACE, "{"},
		{token.RETURN, "return"}, {token.TRUE, "true"}, {token.SEMICOLON, ";"},
		{token.RBRACE, "}"}, {token.ELSE, "else"}, {token.LBRACE, "{"},
		{token.RETURN, "return"}, {token.FALSE, "false"}, {token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
	}

	l := New(input)

	for i, et := range tests {
		tok := l.NextToken()

		if tok.Type != et.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, et.expectedType, tok.Type)
		}

		if tok.Literal != et.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, et.expectedLiteral, tok.Literal)
		}
	}
}
