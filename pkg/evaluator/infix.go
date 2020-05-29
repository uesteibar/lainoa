package evaluator

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/object"
	"github.com/uesteibar/lainoa/pkg/token"
)

func evalInfix(infix *ast.InfixExpression) object.Object {
	left := Eval(infix.Left)
	right := Eval(infix.Right)

	switch {
	case left.Type() == object.INTEGER_OBJECT && right.Type() == object.INTEGER_OBJECT:
		left := left.(*object.Integer)
		right := right.(*object.Integer)
		return evalIntegerInfixExpression(left, infix.Operator, right)
	case infix.Operator == token.EQ:
		return nativeBoolToBoolean(left == right)
	case infix.Operator == token.NOT_EQ:
		return nativeBoolToBoolean(left != right)
	default:
		return NULL
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
		return NULL
	}
}
