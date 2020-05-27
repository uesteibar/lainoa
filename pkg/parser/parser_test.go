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

func assertLetStatement(t *testing.T, s ast.Statement, name string, value interface{}) {
	assert.Equal(t, "let", s.TokenLiteral())

	letStmt, ok := s.(*ast.LetStatement)
	assert.True(t, ok)

	assert.Equal(t, name, letStmt.Name.Value)
	assert.EqualValues(t, value, letStmt.Value.String())
}

func assertIntegerLiteral(t *testing.T, il ast.Expression, expectedValue int64) {
	integer, ok := il.(*ast.IntegerLiteral)
	assert.True(t, ok)

	assert.EqualValues(t, expectedValue, integer.Value)
	assert.Equal(t, strconv.Itoa(int(expectedValue)), integer.Token.Literal)
	assert.EqualValues(t, token.INT, integer.Token.Type)
}

func assertIdentifier(t *testing.T, exp ast.Expression, expectedValue string) {
	ident, ok := exp.(*ast.Identifier)
	assert.True(t, ok)

	assert.Equal(t, expectedValue, ident.Value)
	assert.Equal(t, expectedValue, ident.TokenLiteral())
}

func assertBoolean(t *testing.T, exp ast.Expression, expectedValue bool) {
	ident, ok := exp.(*ast.Boolean)
	assert.True(t, ok)

	assert.Equal(t, expectedValue, ident.Value)
}

func assertLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) {
	switch v := expected.(type) {
	case int:
		assertIntegerLiteral(t, exp, int64(v))
	case int64:
		assertIntegerLiteral(t, exp, v)
	case string:
		assertIdentifier(t, exp, v)
	case bool:
		assertBoolean(t, exp, v)
	default:
		t.Errorf("type of exp not handled. got=%T", exp)
	}
}

func assertInfixExpression(
	t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{},
) {
	opExp, ok := exp.(*ast.InfixExpression)
	assert.True(t, ok)

	assertLiteralExpression(t, opExp.Left, left)
	assert.Equal(t, operator, opExp.Operator)
	assertLiteralExpression(t, opExp.Right, right)

}

func TestParseLetStatements(t *testing.T) {
	l := lexer.New(`
		let x = 5;
		let y = 10;
		let z = -10;
		let w = true;
	`)
	p := New(l)
	program := p.ParseProgram()

	assertNoErrors(t, p)
	assert.Len(t, program.Statements, 4)

	assertLetStatement(t, program.Statements[0], "x", "5")
	assertLetStatement(t, program.Statements[1], "y", "10")
	assertLetStatement(t, program.Statements[2], "z", "(-10)")
	assertLetStatement(t, program.Statements[3], "w", "true")
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

func TestParseReturnStatements(t *testing.T) {
	l := lexer.New(`
		return true;
		return 1;
	`)
	p := New(l)
	program := p.ParseProgram()

	assertNoErrors(t, p)
	assert.Len(t, program.Statements, 2)

	stmt, ok := program.Statements[0].(*ast.ReturnStatement)
	assert.True(t, ok)
	assert.Equal(t, "return", stmt.TokenLiteral())
	assertBoolean(t, stmt.Value, true)

	stmt, ok = program.Statements[1].(*ast.ReturnStatement)
	assert.True(t, ok)
	assert.Equal(t, "return", stmt.TokenLiteral())
	assertIntegerLiteral(t, stmt.Value, 1)

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

	assertLiteralExpression(t, stmt.Expression, "foo")
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

	assertLiteralExpression(t, stmt.Expression, 550)
}

func TestPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    interface{}
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
		{"!true;", "!", true},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		assertNoErrors(t, p)

		assert.Len(t, program.Statements, 1)
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.True(t, ok)
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		assert.True(t, ok)

		assert.Equal(t, tt.operator, exp.Operator)
		assertLiteralExpression(t, exp.Right, tt.value)
	}
}

func TestInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"5 != true;", 5, "!=", true},
		{"false != true;", false, "!=", true},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()

		assertNoErrors(t, p)

		assert.Len(t, program.Statements, 1)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.True(t, ok)

		assertInfixExpression(t, stmt.Expression, tt.leftValue, tt.operator, tt.rightValue)
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a;",
			"(!(-a))",
		},
		{
			"a + b + c;",
			"((a + b) + c)",
		},
		{
			"a + b - c;",
			"((a + b) - c)",
		},
		{
			"a * b * c;",
			"((a * b) * c)",
		},
		{
			"a * b / c;",
			"((a * b) / c)",
		},
		{
			"a + b / c;",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f;",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"3 + 4 - 5 * 5",
			"((3 + 4) - (5 * 5))",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4;",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5;",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"3 > 9 == false",
			"((3 > 9) == false)",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		assertNoErrors(t, p)

		actual := program.String()
		assert.Equal(t, tt.expected, actual)
	}
}

func TestGroupedExpressionsParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		assertNoErrors(t, p)

		actual := program.String()
		assert.Equal(t, tt.expected, actual)
	}
}

func TestBooleanExpression(t *testing.T) {
	l := lexer.New(`
		return true;
		false;
	`)
	p := New(l)
	program := p.ParseProgram()
	assertNoErrors(t, p)

	assert.Len(t, program.Statements, 2)

	ret, ok := program.Statements[0].(*ast.ReturnStatement)
	assert.True(t, ok)
	boolean, ok := ret.Value.(*ast.Boolean)
	assert.True(t, ok)

	assert.EqualValues(t, token.TRUE, boolean.Token.Type)
	assert.Equal(t, true, boolean.Value)

	exp, ok := program.Statements[1].(*ast.ExpressionStatement)
	assert.True(t, ok)
	boolean, ok = exp.Expression.(*ast.Boolean)
	assert.True(t, ok)

	assert.EqualValues(t, token.FALSE, boolean.Token.Type)
	assert.Equal(t, false, boolean.Value)
}

func TestIfExpression(t *testing.T) {
	l := lexer.New(`if (x < y) { x }`)
	p := New(l)
	program := p.ParseProgram()
	assertNoErrors(t, p)

	assert.Len(t, program.Statements, 1)

	exp, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	ifexp, ok := exp.Expression.(*ast.IfExpression)
	assert.True(t, ok)
	assert.EqualValues(t, token.IF, ifexp.Token.Type)

	cond, ok := ifexp.Condition.(*ast.InfixExpression)
	assert.True(t, ok)
	assertInfixExpression(t, cond, "x", token.LT, "y")

	consBlock := ifexp.Consequence
	assert.Len(t, consBlock.Statements, 1)

	ident, ok := consBlock.Statements[0].(*ast.ExpressionStatement)
	assertIdentifier(t, ident.Expression, "x")

	assert.Equal(t, "if(x < y) x", ifexp.String())
}

func TestIfElseExpression(t *testing.T) {
	l := lexer.New(`if (x < y) { x } else { y }`)
	p := New(l)
	program := p.ParseProgram()
	assertNoErrors(t, p)

	assert.Len(t, program.Statements, 1)

	exp, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	ifexp, ok := exp.Expression.(*ast.IfExpression)
	assert.True(t, ok)
	assert.EqualValues(t, token.IF, ifexp.Token.Type)

	cond, ok := ifexp.Condition.(*ast.InfixExpression)
	assert.True(t, ok)
	assertInfixExpression(t, cond, "x", token.LT, "y")

	consBlock := ifexp.Consequence
	assert.Len(t, consBlock.Statements, 1)

	ident, ok := consBlock.Statements[0].(*ast.ExpressionStatement)
	assertIdentifier(t, ident.Expression, "x")

	altBlock := ifexp.Consequence
	assert.Len(t, altBlock.Statements, 1)

	altIdent, ok := altBlock.Statements[0].(*ast.ExpressionStatement)
	assertIdentifier(t, altIdent.Expression, "x")

	assert.Equal(t, "if(x < y) x else y", ifexp.String())
}
