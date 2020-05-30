package evaluator

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/object"
)

func evalReturnStatement(ret *ast.ReturnStatement) *object.ReturnValue {
	value := Eval(ret.Value)

	return &object.ReturnValue{Value: value}
}
