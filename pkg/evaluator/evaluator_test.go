package evaluator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesteibar/lainoa/pkg/lexer"
	"github.com/uesteibar/lainoa/pkg/object"
	"github.com/uesteibar/lainoa/pkg/parser"
)

func eval(input string) object.Object {
	l := lexer.New(input, "/path/to/file")
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func assertIntegerObject(t *testing.T, obj object.Object, expected int64) {
	integer, ok := obj.(*object.Integer)
	assert.True(t, ok)

	assert.Equal(t, expected, integer.Value)
}

func assertStringObject(t *testing.T, obj object.Object, expected string) {
	str, ok := obj.(*object.String)
	assert.True(t, ok)

	assert.Equal(t, expected, str.Value)
}

func assertBooleanObject(t *testing.T, obj object.Object, expected bool) {
	integer, ok := obj.(*object.Boolean)
	assert.True(t, ok)

	assert.Equal(t, expected, integer.Value)
}

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"-5", -5},
		{"10", 10},
		{"--10", 10},
		{"---10", -10},
		{"+10", 10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		evaluated := eval(tt.input)
		assertIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalStringExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`"5"`, "5"},
		{`"unai"`, "unai"},
		{`"unai" + " " + "esteibar"`, "unai esteibar"},
	}

	for _, tt := range tests {
		evaluated := eval(tt.input)
		assertStringObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
		{"nil == nil", true},
		{"let a = nil; a == nil", true},
		{"(1 > 2) == nil", false},
		{"\"nil\" != nil", true},
	}

	for _, tt := range tests {
		evaluated := eval(tt.input)
		assertBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := eval(tt.input)
		assertBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	}

	for _, tt := range tests {
		evaluated := eval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			assertIntegerObject(t, evaluated, int64(integer))
		} else {
			assert.Equal(t, NIL, evaluated)
		}
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{`
		if (10 > 1) {
		  if (10 > 1) {
			return 10;
		  }

		  return 1;
		}`, 10},
	}

	for _, tt := range tests {
		evaluated := eval(tt.input)
		assertIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"+true",
			"unknown operator: +BOOLEAN",
		},
		{
			"true + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`"unai" - "ai"; true;`,
			"unknown operator: STRING - STRING",
		},
		{
			"5; true + false; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`if (10 > 1) {
				if (10 > 1) {
					return true + false;
				}

				return 1;
			}`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"identifier not found: foobar",
		},
		{
			"if(true) { let scoped = 1 }; scoped;",
			"identifier not found: scoped",
		},
		{
			"if(false) {} else { let scoped = 1 }; scoped;",
			"identifier not found: scoped",
		},
		{
			"let a = 1; if(false) { let scoped = a + 1 }; scoped;",
			"identifier not found: scoped",
		},
		{
			"if(true) { a = 1 };",
			"can't assign identier `a` because it doesn't exist, you need to do `let a = 1` first",
		},
		{
			"let a = 1; let a = 10;",
			"can't re-bind already bound identifier `a`",
		},
	}

	for _, tt := range tests {
		evaluated := eval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		assert.True(t, ok)

		assert.Equal(t, tt.expectedMessage, errObj.Message)
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
		{"let a = 5; if (a == 5) { a + 10 };", 15},
		{"let a = 5; a = 10; a;", 10},
		{"let a = 5; if (a == 5) { a = 10 }; a;", 10},
	}

	for _, tt := range tests {
		assertIntegerObject(t, eval(tt.input), tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "fun(x) { x + 2; };"

	evaluated := eval(input)
	fn, ok := evaluated.(*object.Function)
	assert.True(t, ok)

	assert.Len(t, fn.Parameters, 1)
	assert.Equal(t, "x", fn.Parameters[0].String())
	assert.Equal(t, "(x + 2)", fn.Body.String())
}

func TestFunctionCalls(t *testing.T) {
	evaluated := eval(`
		let multiply = fun(num) {
			let number = num
			return fun(multiplyer) {
				number * multiplyer
			}
		}

		multiply(2)(5)
	`)

	assertIntegerObject(t, evaluated, 10)
}

func TestFunctionCurrying(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{
			`
			let multiply = fun(a, b) {
				a * b
			}

			let by_five = multiply(5)

			by_five(10)
			`,
			50,
		},
		{
			`
			let multiply = fun(a, b, c, d) {
				a * b * c * d
			}

			let by_fifty = multiply(5, 10)

			by_fifty(3)(2)
			`,
			300,
		},
	}

	for _, tt := range tests {
		evaluated := eval(tt.input)

		assertIntegerObject(t, evaluated, tt.expected)
	}
}

func TestFunctionCallErrors(t *testing.T) {
	evaluated := eval(`
		let multiply = fun(num) {
			let number = num

			return number * 2
		}

		multiply(2)(5)
	`)

	errObj, ok := evaluated.(*object.Error)
	assert.True(t, ok)

	assert.Equal(t, "expected 4 to be a function, got INTEGER", errObj.Message)
}

func TestFunctionArgumentErrors(t *testing.T) {
	evaluated := eval(`
		let num = 2
		let multiply = fun(num) {
			return num * 2
		}

		multiply(num)
	`)

	errObj, ok := evaluated.(*object.Error)
	assert.True(t, ok)

	assert.Equal(t, "can't re-bind already bound identifier `num`", errObj.Message)
}

func TestComplexProgram(t *testing.T) {
	evaluated := eval(`
	let main = fun() {
		let result = 0

		let add = fun(a) {
		  let number = a

		  return fun(b) {
			return number + b
		  }
		}

		let addFive = add(5)

		result = addFive(10)

		# result = 10

		result
	}

	main()
	`)

	assertIntegerObject(t, evaluated, 15)
}

func TestLenBuiltin(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world")`, 11},
		{`len(1)`, "argument to `len` not supported, got INTEGER"},
		{`len("one", "two")`, "wrong number of arguments. got=2, want=1"},
	}

	for _, tt := range tests {
		evaluated := eval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			assertIntegerObject(t, evaluated, int64(expected))
		case string:
			errObj, ok := evaluated.(*object.Error)
			assert.True(t, ok)
			assert.Equal(t, expected, errObj.Message)
		}
	}
}

func TestNil(t *testing.T) {
	evaluated := eval(`let a = nil; a`)

	_, ok := evaluated.(*object.Nil)
	assert.True(t, ok)
}

func TestArray(t *testing.T) {
	evaluated := eval(`[3, 2, 1]`)

	res, ok := evaluated.(*object.Array)
	assert.True(t, ok)

	assertIntegerObject(t, res.Elements[0], 3)
	assertIntegerObject(t, res.Elements[1], 2)
	assertIntegerObject(t, res.Elements[2], 1)
}

func TestArrayIndex(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"[][0]",
			nil,
		},
		{
			"[1, 2, 3][0]",
			1,
		},
		{
			"[1, 2, 3][1]",
			2,
		},
		{
			"[1, 2, 3][2]",
			3,
		},
		{
			"let i = 0; [1][i];",
			1,
		},
		{
			"[1, 2, 3][1 + 1];",
			3,
		},
		{
			"let array = [1, 2, 3]; array[2];",
			3,
		},
		{
			"let array = [1, 2, 3]; array[0] + array[1] + array[2];",
			6,
		},
		{
			"let array = [1, 2, 3]; let i = array[0]; array[i]",
			2,
		},
		{
			"[1, 2, 3][3]",
			nil,
		},
		{
			"[1, 2, 3][-1]",
			nil,
		},
	}

	for _, tt := range tests {
		evaluated := eval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			assertIntegerObject(t, evaluated, int64(integer))
		} else {
			assert.Equal(t, NIL, evaluated)
		}
	}
}

func TestIndexError(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"[3, 2, 1][\"3\"]",
			"expected INTEGER as index for array, got STRING",
		},
		{
			"1000[3]",
			"type INTEGER doesn't support index operations",
		},
	}

	for _, tt := range tests {
		evaluated := eval(tt.input)
		err, ok := evaluated.(*object.Error)
		assert.True(t, ok)

		assert.Equal(t, tt.expected, err.Message)
	}
}
func TestArrayPush(t *testing.T) {
	evaluated := eval(`
	let array = [3, 2, 1]
	let new_array = push(array, 0)

	let res = [new_array, array]
	res`)

	res, ok := evaluated.(*object.Array)
	assert.True(t, ok)
	newArray, ok := res.Elements[0].(*object.Array)
	assert.True(t, ok)

	assertIntegerObject(t, newArray.Elements[0], 3)
	assertIntegerObject(t, newArray.Elements[1], 2)
	assertIntegerObject(t, newArray.Elements[2], 1)
	assertIntegerObject(t, newArray.Elements[3], 0)

	oldArray, ok := res.Elements[1].(*object.Array)
	assert.True(t, ok)

	assertIntegerObject(t, oldArray.Elements[0], 3)
	assertIntegerObject(t, oldArray.Elements[1], 2)
	assertIntegerObject(t, oldArray.Elements[2], 1)
}

func TestArrayLen(t *testing.T) {
	evaluated := eval(`len([3, 2, 1])`)

	assertIntegerObject(t, evaluated, 3)
}
