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
		l.position = l.readPosition
		l.readPosition++
	}
}

var symbols = map[byte]token.TokenType{
	'=': token.ASSIGN,
	',': token.COMMA,
	';': token.SEMICOLON,
	'(': token.LPAREN,
	')': token.RPAREN,
	'{': token.LBRACE,
	'}': token.RBRACE,
	'+': token.PLUS,
}

func (l *Lexer) NextToken() (t token.Token) {
	l.skipWhitespace()

	if tokenType, exists := symbols[l.ch]; exists {
		t = newToken(tokenType, l.ch)

		l.readChar()
	} else if l.ch == 0 {
		t.Literal = ""
		t.Type = token.EOF

		l.readChar()
	} else if isLetter(l.ch) {
		t.Literal = l.readIdentifier()
		t.Type = token.LookupIdentType(t.Literal)
	} else if isDigit(l.ch) {
		t.Type = token.INT
		t.Literal = l.readNumber()
	} else {
		t = newToken(token.ILLEGAL, l.ch)

		l.readChar()
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
