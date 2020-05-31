package evaluator

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/object"
)

func evalAssign(assign *ast.AssignExpression, env *object.Environment) object.Object {
	val := Eval(assign.Value, env)
	if object.IsError(val) {
		return val
	}

	return env.Rebind(assign.Name.Value, val)
}
