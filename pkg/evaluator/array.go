package evaluator

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/object"
)

func evalArray(array *ast.ArrayExpression, env *object.Environment) object.Object {
	elements, err := evalExpressions(array.Expressions, env)
	if err != nil {
		return err
	}

	return &object.Array{Elements: elements}
}
