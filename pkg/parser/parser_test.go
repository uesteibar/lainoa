package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/lexer"
)

func assertLetStatement(t *testing.T, s ast.Statement, name string) {
	assert.Equal(t, "let", s.TokenLiteral())

	letStmt, ok := s.(*ast.LetStatement)
	assert.True(t, ok)

	assert.Equal(t, name, letStmt.Name.Value)
}

func TestParseLetStatements(t *testing.T) {
	l := lexer.New(`
		let x = 5;
		let y = 10;
	`)
	p := New(l)
	program := p.ParseProgram()

	assert.Len(t, program.Statements, 2)

	assertLetStatement(t, program.Statements[0], "x")
	assertLetStatement(t, program.Statements[1], "y")
}

func TestParseLetStatementError(t *testing.T) {
	l := lexer.New(`
		let 5;
		let x = 10;
		let x + 1;
	`)
	p := New(l)
	program := p.ParseProgram()

	assert.NotNil(t, program)

	errors := p.Errors()
	assert.Len(t, errors, 2)
	assert.Equal(t, "expected next token to be IDENT, got INT instead", errors[0])
	assert.Equal(t, "expected next token to be =, got + instead", errors[1])
}

func assertReturnStatement(t *testing.T, s ast.Statement) {
	assert.Equal(t, "return", s.TokenLiteral())
	_, ok := s.(*ast.ReturnStatement)
	assert.True(t, ok)
}

func TestParseReturnStatements(t *testing.T) {
	l := lexer.New(`
		return true;
		return 1;
	`)
	p := New(l)
	program := p.ParseProgram()

	assert.Len(t, program.Statements, 2)

	assertReturnStatement(t, program.Statements[0])
	assertReturnStatement(t, program.Statements[1])
}
