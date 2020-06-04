package lexer

import (
	"testing"

	"github.com/uesteibar/lainoa/pkg/token"
	"gotest.tools/assert"
)

func TestNextToken(t *testing.T) {
	input := `let ten = 10;
		let five = 6;
		five = 5;

		let add = fun (x, y) {
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

		10 == 10;
		9 != 10;

		-5;
		a + bb

		nil

		# this is a comment
		"test-string" # inline comment
		let name = "unai esteibar";

		let array = [1, 2]
		let new_array = push(array, 0)
		[1, 2]`

	tests := [][]struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{{token.LET, "let"}, {token.IDENT, "ten"}, {token.ASSIGN, "="}, {token.INT, "10"}, {token.SEMICOLON, ";"}},
		{{token.LET, "let"}, {token.IDENT, "five"}, {token.ASSIGN, "="}, {token.INT, "6"}, {token.SEMICOLON, ";"}},
		{{token.IDENT, "five"}, {token.ASSIGN, "="}, {token.INT, "5"}, {token.SEMICOLON, ";"}},
		{},
		{{token.LET, "let"}, {token.IDENT, "add"}, {token.ASSIGN, "="}, {token.FUNCTION, "fun"}, {token.LPAREN, "("}, {token.IDENT, "x"}, {token.COMMA, ","}, {token.IDENT, "y"}, {token.RPAREN, ")"}, {token.LBRACE, "{"}},
		{{token.IDENT, "x"}, {token.PLUS, "+"}, {token.IDENT, "y"}, {token.SEMICOLON, ";"}},
		{{token.RBRACE, "}"}},
		{},
		{{token.LET, "let"}, {token.IDENT, "result"}, {token.ASSIGN, "="}, {token.IDENT, "add"}, {token.LPAREN, "("}, {token.IDENT, "ten"}, {token.COMMA, ","}, {token.IDENT, "five"}, {token.RPAREN, ")"}, {token.SEMICOLON, ";"}},
		{},
		{{token.BANG, "!"}, {token.MINUS, "-"}, {token.SLASH, "/"}, {token.ASTERISK, "*"}, {token.INT, "5"}, {token.SEMICOLON, ";"}},
		{{token.INT, "5"}, {token.LT, "<"}, {token.INT, "10"}, {token.GT, ">"}, {token.INT, "5"}, {token.SEMICOLON, ";"}},
		{},
		{{token.IF, "if"}, {token.LPAREN, "("}, {token.INT, "5"}, {token.LT, "<"}, {token.INT, "10"}, {token.RPAREN, ")"}, {token.LBRACE, "{"}},
		{{token.RETURN, "return"}, {token.TRUE, "true"}, {token.SEMICOLON, ";"}},
		{{token.RBRACE, "}"}, {token.ELSE, "else"}, {token.LBRACE, "{"}},
		{{token.RETURN, "return"}, {token.FALSE, "false"}, {token.SEMICOLON, ";"}},
		{{token.RBRACE, "}"}},
		{},
		{{token.INT, "10"}, {token.EQ, "=="}, {token.INT, "10"}, {token.SEMICOLON, ";"}},
		{{token.INT, "9"}, {token.NOT_EQ, "!="}, {token.INT, "10"}, {token.SEMICOLON, ";"}},
		{},
		{{token.MINUS, "-"}, {token.INT, "5"}, {token.SEMICOLON, ";"}},
		{{token.IDENT, "a"}, {token.PLUS, "+"}, {token.IDENT, "bb"}},
		{},
		{{token.NIL, "nil"}},
		{},
		{{token.COMMENT, "this is a comment"}},
		{{token.STRING, "test-string"}, {token.COMMENT, "inline comment"}},
		{{token.LET, "let"}, {token.IDENT, "name"}, {token.ASSIGN, "="}, {token.STRING, "unai esteibar"}, {token.SEMICOLON, ";"}},
		{},
		{{token.LET, "let"}, {token.IDENT, "array"}, {token.ASSIGN, "="}, {token.LBRACKET, "["}, {token.INT, "1"}, {token.COMMA, ","}, {token.INT, "2"}, {token.RBRACKET, "]"}},
		{{token.LET, "let"}, {token.IDENT, "new_array"}, {token.ASSIGN, "="}, {token.IDENT, "push"}, {token.LPAREN, "("}, {token.IDENT, "array"}, {token.COMMA, ","}, {token.INT, "0"}, {token.RPAREN, ")"}},
		{{token.LBRACKET, "["}, {token.INT, "1"}, {token.COMMA, ","}, {token.INT, "2"}, {token.RBRACKET, "]"}},
	}

	l := New(input, "/path/to/file")

	for i, line := range tests {
		expectedLine := i + 1
		for _, et := range line {
			tok := l.NextToken()

			assert.Equal(t, et.expectedType, tok.Type)
			assert.Equal(t, et.expectedLiteral, tok.Literal)
			assert.Equal(t, expectedLine, tok.Metadata.Line)
			assert.Equal(t, "/path/to/file", tok.Metadata.File)
		}
	}
}
