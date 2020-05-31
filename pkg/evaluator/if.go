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
		env.AddScope()
		res := Eval(ifexp.Consequence, env)
		env.ReleaseScope()
		return res
	} else if ifexp.Alternative != nil {
		env.AddScope()
		res := Eval(ifexp.Alternative, env)
		env.ReleaseScope()
		return res
	} else {
		return NULL
	}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}
