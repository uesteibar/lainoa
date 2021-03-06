package parser

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/lexer"
	"github.com/uesteibar/lainoa/pkg/token"
)

func lex(input string) *lexer.Lexer {
	return lexer.New(input, "/path/to/file")
}

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
	l := lex(`
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
	l := lex(`
		let 5;
		let x = 10;
		let x + 1;
	`)
	p := New(l)
	program := p.ParseProgram()

	assert.NotNil(t, program)

	errors := p.Errors()
	assert.Len(t, errors, 2)
	assert.Equal(t, "/path/to/file:1 expected next token to be IDENT, got INT instead", errors[0].String())
	assert.Equal(t, "/path/to/file:3 expected next token to be =, got + instead", errors[1].String())
}

func TestParseReturnStatements(t *testing.T) {
	l := lex(`
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
	l := lex(`
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
	l := lex(`
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

func TestStringExpression(t *testing.T) {
	l := lex(`"unai"`)

	p := New(l)
	program := p.ParseProgram()

	assertNoErrors(t, p)
	assert.Len(t, program.Statements, 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	str, ok := stmt.Expression.(*ast.StringLiteral)
	assert.True(t, ok)
	assert.Equal(t, "unai", str.Value)
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
		l := lex(tt.input)
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

func TestPrefixExpressionsErrors(t *testing.T) {
	l := lex(`
		!5;
		&10;
	`)
	p := New(l)
	program := p.ParseProgram()

	assert.NotNil(t, program)

	errors := p.Errors()
	assert.Len(t, errors, 1)
	assert.Equal(t, "/path/to/file:2 prefix operation & not recognized", errors[0].String())
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
		l := lex(tt.input)
		p := New(l)
		program := p.ParseProgram()

		assertNoErrors(t, p)

		assert.Len(t, program.Statements, 1)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.True(t, ok)

		assertInfixExpression(t, stmt.Expression, tt.leftValue, tt.operator, tt.rightValue)
	}
}

func TestInfixExpressionsErrors(t *testing.T) {
	l := lex(`
		5 *;
		5 $ a;
		== 10;
	`)
	p := New(l)
	program := p.ParseProgram()

	assert.NotNil(t, program)

	errors := p.Errors()
	assert.Len(t, errors, 3)
	assert.Equal(t, "/path/to/file:1 prefix operation ; not recognized", errors[0].String())
	assert.Equal(t, "/path/to/file:2 prefix operation $ not recognized", errors[1].String())
	assert.Equal(t, "/path/to/file:3 prefix operation == not recognized", errors[2].String())
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
		{
			"a + add(b * c) + d",
			"((a + add((b * c))) + d)",
		},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g))",
		},
	}

	for _, tt := range tests {
		l := lex(tt.input)
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
		l := lex(tt.input)
		p := New(l)
		program := p.ParseProgram()
		assertNoErrors(t, p)

		actual := program.String()
		assert.Equal(t, tt.expected, actual)
	}
}

func TestBooleanExpression(t *testing.T) {
	l := lex(`
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
	l := lex(`if (x < y) { x }`)
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

	assert.Equal(t, "if (x < y) x", ifexp.String())
}

func TestIfElseExpression(t *testing.T) {
	l := lex(`if (x < y) { x } else { y }`)
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

	assert.Equal(t, "if (x < y) x else y", ifexp.String())
}

func TestIfElseExpressionsErrors(t *testing.T) {
	l := lex(`
		if (x < y) {
			x
		}
		if (x < y) x;
		if (x < y) { x
	`)
	p := New(l)
	program := p.ParseProgram()

	assert.NotNil(t, program)

	errors := p.Errors()
	assert.Len(t, errors, 2)
	assert.Equal(t, "/path/to/file:4 expected next token to be {, got IDENT instead", errors[0].String())
	assert.Equal(t, "/path/to/file:6 expected } at the end of the block, got EOF instead", errors[1].String())
}

func TestFunctionExpression(t *testing.T) {
	l := lex(`
		fun(a, b) {
			a + b;
		}
	`)
	p := New(l)
	program := p.ParseProgram()
	assertNoErrors(t, p)

	assert.Len(t, program.Statements, 1)

	exp, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	fun, ok := exp.Expression.(*ast.FunctionLiteral)
	assert.True(t, ok)

	assert.Len(t, fun.Parameters, 2)
	assertIdentifier(t, fun.Parameters[0], "a")
	assertIdentifier(t, fun.Parameters[1], "b")

	assert.Len(t, fun.Body.Statements, 1)

	exp, ok = fun.Body.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	assertInfixExpression(t, exp.Expression, "a", "+", "b")
}

func TestFunctionErrors(t *testing.T) {
	l := lex(`fun(1, b) { a + b; }`)
	p := New(l)
	p.ParseProgram()

	errors := p.Errors()
	assert.Len(t, errors, 1)
	assert.Equal(t,
		"/path/to/file:1 Function parameters can only be identifiers, found '1' instead",
		errors[0].String(),
	)
}

func TestFunctionParameterParsing(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
	}{
		{input: "fun() {};", expectedParams: []string{}},
		{input: "fun(x) {};", expectedParams: []string{"x"}},
		{input: "fun(x, y, z) {};", expectedParams: []string{"x", "y", "z"}},
	}

	for _, tt := range tests {
		l := lex(tt.input)
		p := New(l)
		program := p.ParseProgram()
		assertNoErrors(t, p)

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		fun := stmt.Expression.(*ast.FunctionLiteral)
		assert.Len(t, fun.Parameters, len(tt.expectedParams))

		for i, ident := range tt.expectedParams {
			assertLiteralExpression(t, fun.Parameters[i], ident)
		}
	}
}

func TestCallExpressionParsing(t *testing.T) {
	input := "add(1, 2 * 3, 4 + 5);"

	l := lex(input)
	p := New(l)
	program := p.ParseProgram()
	assertNoErrors(t, p)

	assert.Len(t, program.Statements, 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	exp, ok := stmt.Expression.(*ast.CallExpression)
	assert.True(t, ok)

	assertIdentifier(t, exp.Function, "add")

	assert.Len(t, exp.Arguments, 3)

	assertLiteralExpression(t, exp.Arguments[0], 1)
	assertInfixExpression(t, exp.Arguments[1], 2, "*", 3)
	assertInfixExpression(t, exp.Arguments[2], 4, "+", 5)
}

func TestCallErrors(t *testing.T) {
	l := lex("add(1, 2 * 3, 4 + 5;")
	p := New(l)
	p.ParseProgram()

	errors := p.Errors()
	assert.Len(t, errors, 1)
	assert.Equal(t,
		"/path/to/file:1 expected next token to be ), got ; instead",
		errors[0].String(),
	)
}

func TestDirectCallExpressionParsing(t *testing.T) {
	input := "fun(a, b) { a + b }(1, 2 * 3, 4 + 5);"

	l := lex(input)
	p := New(l)
	program := p.ParseProgram()
	assertNoErrors(t, p)

	assert.Len(t, program.Statements, 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	exp, ok := stmt.Expression.(*ast.CallExpression)
	assert.True(t, ok)

	assert.Equal(t, "fun(a, b) (a + b)", exp.Function.String())

	assert.Len(t, exp.Arguments, 3)

	assertLiteralExpression(t, exp.Arguments[0], 1)
	assertInfixExpression(t, exp.Arguments[1], 2, "*", 3)
	assertInfixExpression(t, exp.Arguments[2], 4, "+", 5)
}

func TestAssignExpressionParsing(t *testing.T) {
	input := "a = 1;"

	l := lex(input)
	p := New(l)
	program := p.ParseProgram()
	assertNoErrors(t, p)

	assert.Len(t, program.Statements, 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	exp, ok := stmt.Expression.(*ast.AssignExpression)
	assert.True(t, ok)

	assertIdentifier(t, exp.Name, "a")
}

func TestComments(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"1 + 1 # inline comment",
			"(1 + 1)",
		},
		{
			`
			# full line comment
			1 + 1
			`,
			"(1 + 1)",
		},
	}

	for _, tt := range tests {
		l := lex(tt.input)
		p := New(l)
		program := p.ParseProgram()
		assertNoErrors(t, p)

		assert.Len(t, program.Statements, 1)
		actual := program.String()
		assert.Equal(t, tt.expected, actual)
	}
}

func TestNil(t *testing.T) {
	l := lex(`let a = nil`)
	p := New(l)
	program := p.ParseProgram()
	assertNoErrors(t, p)

	assert.Len(t, program.Statements, 1)

	let, ok := program.Statements[0].(*ast.LetStatement)
	assert.True(t, ok)

	_, ok = let.Value.(*ast.NilLiteral)
	assert.True(t, ok)
}

func TestArrays(t *testing.T) {
	l := lex(`[1, 2, "name"]`)
	p := New(l)
	program := p.ParseProgram()
	assertNoErrors(t, p)

	assert.Len(t, program.Statements, 1)

	exp, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	array, ok := exp.Expression.(*ast.ArrayExpression)
	assert.True(t, ok)

	assert.Len(t, array.Expressions, 3)
	assertIntegerLiteral(t, array.Expressions[0], 1)
	assertIntegerLiteral(t, array.Expressions[1], 2)

	str, ok := array.Expressions[2].(*ast.StringLiteral)
	assert.True(t, ok)
	assert.Equal(t, "name", str.Value)
}

func TestIndexExpressions(t *testing.T) {
	l := lex(`array[1]`)
	p := New(l)
	program := p.ParseProgram()
	assertNoErrors(t, p)

	assert.Len(t, program.Statements, 1)

	exp, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	indexp, ok := exp.Expression.(*ast.IndexExpression)
	assert.True(t, ok)

	assertIdentifier(t, indexp.Left, "array")
	assertIntegerLiteral(t, indexp.Index, 1)
}
