package parser

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/lexer"
	"github.com/uesteibar/lainoa/pkg/token"
)

func assertNoErrors(t *testing.T, p *Parser) {
	assert.Len(t, p.Errors(), 0)
}

func assertLetStatement(t *testing.T, s ast.Statement, name string, value string) {
	assert.Equal(t, "let", s.TokenLiteral())

	letStmt, ok := s.(*ast.LetStatement)
	assert.True(t, ok)

	assert.Equal(t, name, letStmt.Name.Value)
	assert.EqualValues(t, value, letStmt.Value.String())
}

func TestParseLetStatements(t *testing.T) {
	l := lexer.New(`
		let x = 5;
		let y = 10;
		let z = -10;
	`)
	p := New(l)
	program := p.ParseProgram()

	assertNoErrors(t, p)
	assert.Len(t, program.Statements, 3)

	assertLetStatement(t, program.Statements[0], "x", "5")
	assertLetStatement(t, program.Statements[1], "y", "10")
	assertLetStatement(t, program.Statements[2], "z", "(-10)")
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

	assertNoErrors(t, p)
	assert.Len(t, program.Statements, 2)

	assertReturnStatement(t, program.Statements[0])
	assertReturnStatement(t, program.Statements[1])
}

func TestIdentifierExpression(t *testing.T) {
	l := lexer.New(`
		foo;
	`)

	p := New(l)
	program := p.ParseProgram()

	assertNoErrors(t, p)
	assert.Len(t, program.Statements, 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	ident, ok := stmt.Expression.(*ast.Identifier)
	assert.True(t, ok)

	assert.Equal(t, "foo", ident.Value)
	assert.Equal(t, "foo", ident.Token.Literal)
	assert.EqualValues(t, token.IDENT, ident.Token.Type)
}

func assertIntegerLiteral(t *testing.T, il ast.Expression, expectedValue int64) {
	integer, ok := il.(*ast.IntegerLiteral)
	assert.True(t, ok)

	assert.EqualValues(t, expectedValue, integer.Value)
	assert.Equal(t, strconv.Itoa(int(expectedValue)), integer.Token.Literal)
	assert.EqualValues(t, token.INT, integer.Token.Type)

}

func TestIntegerExpression(t *testing.T) {
	l := lexer.New(`
		550;
	`)

	p := New(l)
	program := p.ParseProgram()

	assertNoErrors(t, p)
	assert.Len(t, program.Statements, 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	assertIntegerLiteral(t, stmt.Expression, 550)
}

func TestPrefixOperator(t *testing.T) {
	l := lexer.New(`
		-550;
		!15;
		return -5;
	`)

	p := New(l)
	program := p.ParseProgram()

	assertNoErrors(t, p)
	assert.Len(t, program.Statements, 3)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	exp, ok := stmt.Expression.(*ast.PrefixExpression)
	assert.True(t, ok)

	assert.EqualValues(t, token.MINUS, exp.Token.Type)
	assert.Equal(t, "-", exp.Token.Literal)
	assert.EqualValues(t, "-", exp.Operator)

	assertIntegerLiteral(t, exp.Right, 550)

	stmt, ok = program.Statements[1].(*ast.ExpressionStatement)
	assert.True(t, ok)

	exp, ok = stmt.Expression.(*ast.PrefixExpression)
	assert.True(t, ok)

	assert.EqualValues(t, token.BANG, exp.Token.Type)
	assert.Equal(t, "!", exp.Token.Literal)
	assert.EqualValues(t, "!", exp.Operator)

	assertIntegerLiteral(t, exp.Right, 15)

	ret, ok := program.Statements[2].(*ast.ReturnStatement)
	assert.True(t, ok)

	assert.EqualValues(t, token.RETURN, ret.Token.Type)

	_, ok = ret.Value.(*ast.PrefixExpression)
	assert.True(t, ok)
}
