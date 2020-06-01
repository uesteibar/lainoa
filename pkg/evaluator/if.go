package evaluator

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/object"
)

func evalIfExpression(ifexp *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ifexp.Condition, env)
	if object.IsError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ifexp.Consequence, object.NewEnclosedEnvironment(env))
	} else if ifexp.Alternative != nil {
		return Eval(ifexp.Alternative, object.NewEnclosedEnvironment(env))
	} else {
		return NIL
	}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NIL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}
