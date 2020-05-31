package evaluator

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/object"
)

func evalLetStatement(let *ast.LetStatement, env *object.Environment) object.Object {
	val := Eval(let.Value, env)
	if object.IsError(val) {
		return val
	}

	env.Set(let.Name.Value, val)

	return val
}
