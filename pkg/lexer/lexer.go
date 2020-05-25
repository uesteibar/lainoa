package lexer

import (
	"github.com/uesteibar/lainoa/pkg/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekNextChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() (t token.Token) {
	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekNextChar() == '=' {
			l.readChar()
			t.Literal = "=="
			t.Type = token.EQ
		} else {
			t = newToken(token.ASSIGN, l.ch)
		}

		l.readChar()
	case '!':
		if l.peekNextChar() == '=' {
			l.readChar()
			t.Literal = "!="
			t.Type = token.NOT_EQ
		} else {
			t = newToken(token.BANG, l.ch)
		}
		l.readChar()
	case ',':
		t = newToken(token.COMMA, l.ch)
		l.readChar()
	case ';':
		t = newToken(token.SEMICOLON, l.ch)
		l.readChar()
	case '(':
		t = newToken(token.LPAREN, l.ch)
		l.readChar()
	case ')':
		t = newToken(token.RPAREN, l.ch)
		l.readChar()
	case '{':
		t = newToken(token.LBRACE, l.ch)
		l.readChar()
	case '}':
		t = newToken(token.RBRACE, l.ch)
		l.readChar()
	case '+':
		t = newToken(token.PLUS, l.ch)
		l.readChar()
	case '-':
		t = newToken(token.MINUS, l.ch)
		l.readChar()
	case '*':
		t = newToken(token.ASTERISK, l.ch)
		l.readChar()
	case '/':
		t = newToken(token.SLASH, l.ch)
		l.readChar()
	case '<':
		t = newToken(token.LT, l.ch)
		l.readChar()
	case '>':
		t = newToken(token.GT, l.ch)
		l.readChar()
	case 0:
		t.Literal = ""
		t.Type = token.EOF

		l.readChar()
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdentType(t.Literal)
		} else if isDigit(l.ch) {
			t.Type = token.INT
			t.Literal = l.readNumber()
		} else {
			t = newToken(token.ILLEGAL, l.ch)

			l.readChar()
		}
	}

	return t
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	initialPosition := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[initialPosition:l.position]
}

func (l *Lexer) readNumber() string {
	initialPosition := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[initialPosition:l.position]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
