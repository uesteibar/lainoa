package evaluator

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/object"
)

func evalIfExpression(ifexp *ast.IfExpression) object.Object {
	condition := Eval(ifexp.Condition)
	if object.IsError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ifexp.Consequence)
	} else if ifexp.Alternative != nil {
		return Eval(ifexp.Alternative)
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
