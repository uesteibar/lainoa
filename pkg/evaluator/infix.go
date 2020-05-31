package evaluator

import (
	"fmt"

	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/object"
	"github.com/uesteibar/lainoa/pkg/token"
)

func evalInfix(infix *ast.InfixExpression, env *object.Environment) object.Object {
	left := Eval(infix.Left, env)
	if object.IsError(left) {
		return left
	}
	right := Eval(infix.Right, env)
	if object.IsError(right) {
		return right
	}

	switch {
	case left.Type() == object.INTEGER_OBJECT && right.Type() == object.INTEGER_OBJECT:
		left := left.(*object.Integer)
		right := right.(*object.Integer)
		return evalIntegerInfixExpression(left, infix.Operator, right)
	case left.Type() == object.STRING_OBJECT && right.Type() == object.STRING_OBJECT:
		left := left.(*object.String)
		right := right.(*object.String)
		return evalStringInfixExpression(left, infix.Operator, right)
	case infix.Operator == token.EQ:
		return nativeBoolToBoolean(left == right)
	case infix.Operator == token.NOT_EQ:
		return nativeBoolToBoolean(left != right)
	case left.Type() != right.Type():
		return object.NewError(
			"type mismatch: %s %s %s",
			left.Type(), infix.Operator, right.Type())
	default:
		return object.NewError(
			"unknown operator: %s %s %s",
			left.Type(), infix.Operator, right.Type())
	}
}

func evalIntegerInfixExpression(left *object.Integer, operator string, right *object.Integer) object.Object {
	switch operator {
	case token.PLUS:
		return &object.Integer{Value: left.Value + right.Value}
	case token.MINUS:
		return &object.Integer{Value: left.Value - right.Value}
	case token.ASTERISK:
		return &object.Integer{Value: left.Value * right.Value}
	case token.SLASH:
		return &object.Integer{Value: left.Value / right.Value}
	case token.LT:
		return nativeBoolToBoolean(left.Value < right.Value)
	case token.GT:
		return nativeBoolToBoolean(left.Value > right.Value)
	case token.EQ:
		return nativeBoolToBoolean(left.Value == right.Value)
	case token.NOT_EQ:
		return nativeBoolToBoolean(left.Value != right.Value)
	default:
		return object.NewError(
			"unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func evalStringInfixExpression(left *object.String, operator string, right *object.String) object.Object {
	switch operator {
	case token.PLUS:
		return &object.String{Value: fmt.Sprintf("%s%s", left.Value, right.Value)}
	default:
		return object.NewError(
			"unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}
