package lexer

import (
	"github.com/uesteibar/lainoa/pkg/token"
)

type Lexer struct {
	input        string
	filename     string
	position     int
	readPosition int
	curLine      int
	ch           byte
}

func New(input string, filename string) *Lexer {
	l := &Lexer{input: input, filename: filename}
	l.readChar()
	l.curLine = 1

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

	if l.isLineBreak() {
		l.curLine++
	}
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
			t.Metadata = l.metadata()
		} else {
			t = l.newToken(token.ASSIGN, l.ch)
		}

		l.readChar()
	case '!':
		if l.peekNextChar() == '=' {
			l.readChar()
			t.Literal = "!="
			t.Type = token.NOT_EQ
			t.Metadata = l.metadata()
		} else {
			t = l.newToken(token.BANG, l.ch)
		}
		l.readChar()
	case ',':
		t = l.newToken(token.COMMA, l.ch)
		l.readChar()
	case ';':
		t = l.newToken(token.SEMICOLON, l.ch)
		l.readChar()
	case '(':
		t = l.newToken(token.LPAREN, l.ch)
		l.readChar()
	case ')':
		t = l.newToken(token.RPAREN, l.ch)
		l.readChar()
	case '{':
		t = l.newToken(token.LBRACE, l.ch)
		l.readChar()
	case '}':
		t = l.newToken(token.RBRACE, l.ch)
		l.readChar()
	case '[':
		t = l.newToken(token.LBRACKET, l.ch)
		l.readChar()
	case ']':
		t = l.newToken(token.RBRACKET, l.ch)
		l.readChar()
	case '+':
		t = l.newToken(token.PLUS, l.ch)
		l.readChar()
	case '-':
		t = l.newToken(token.MINUS, l.ch)
		l.readChar()
	case '*':
		t = l.newToken(token.ASTERISK, l.ch)
		l.readChar()
	case '/':
		t = l.newToken(token.SLASH, l.ch)
		l.readChar()
	case '<':
		t = l.newToken(token.LT, l.ch)
		l.readChar()
	case '>':
		t = l.newToken(token.GT, l.ch)
		l.readChar()
	case '"':
		t.Metadata = l.metadata()
		t.Literal = l.readString()
		t.Type = token.STRING
		l.readChar()
	case '#':
		t.Metadata = l.metadata()
		t.Literal = l.readComment()
		t.Type = token.COMMENT
		l.readChar()
	case 0:
		t.Metadata = l.metadata()
		t.Literal = ""
		t.Type = token.EOF

		l.readChar()
	default:
		t.Metadata = l.metadata()
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdentType(t.Literal)
		} else if isDigit(l.ch) {
			t.Type = token.INT
			t.Literal = l.readNumber()
		} else {
			t = l.newToken(token.ILLEGAL, l.ch)

			l.readChar()
		}
	}

	return t
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.isLineBreak() || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) isLineBreak() bool {
	return l.ch == '\n'
}

func (l *Lexer) readString() string {
	l.readChar()
	initialPosition := l.position
	for l.ch != '"' {
		l.readChar()
	}

	return l.input[initialPosition:l.position]
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

func (l *Lexer) readComment() string {
	l.readChar()
	l.skipWhitespace()
	initialPosition := l.position
	for l.ch != '\n' && l.ch != '\r' && l.ch != 0 {
		l.readChar()
	}

	return l.input[initialPosition:l.position]
}

func (l *Lexer) newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), Metadata: l.metadata()}
}

func (l *Lexer) metadata() token.Metadata {
	return token.Metadata{Line: l.curLine, File: l.filename}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
