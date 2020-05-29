package evaluator

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesteibar/lainoa/pkg/lexer"
	"github.com/uesteibar/lainoa/pkg/object"
	"github.com/uesteibar/lainoa/pkg/parser"
)

func eval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func assertIntegerObject(t *testing.T, obj object.Object, expected int64) {
	integer, ok := obj.(*object.Integer)
	assert.True(t, ok)

	assert.Equal(t, expected, integer.Value)
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
		log.Println(evaluated)
		assertIntegerObject(t, evaluated, tt.expected)
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
	}

	for _, tt := range tests {
		evaluated := eval(tt.input)
		log.Println(evaluated)
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
		log.Println(evaluated)
		assertBooleanObject(t, evaluated, tt.expected)
	}
}
