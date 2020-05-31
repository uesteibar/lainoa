package evaluator

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/object"
)

func evalReturnStatement(ret *ast.ReturnStatement) object.Object {
	value := Eval(ret.Value)
	if object.IsError(value) {
		return value
	}

	return &object.ReturnValue{Value: value}
}
