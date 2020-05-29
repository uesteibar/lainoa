package evaluator

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/object"
	"github.com/uesteibar/lainoa/pkg/token"
)

func evalPrefix(prefix *ast.PrefixExpression) object.Object {
	right := Eval(prefix.Right)

	switch prefix.Token.Type {
	case token.BANG:
		return evalBangOperation(right)
	case token.MINUS:
		return evalMinusOperation(right)
	case token.PLUS:
		return evalPlusOperation(right)
	}

	return nil
}

func evalBangOperation(right object.Object) *object.Boolean {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalPlusOperation(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJECT {
		return NULL
	}

	val := right.(*object.Integer).Value
	return &object.Integer{Value: val}
}

func evalMinusOperation(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJECT {
		return NULL
	}

	val := right.(*object.Integer).Value
	return &object.Integer{Value: -val}
}
