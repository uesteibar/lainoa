package evaluator

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/object"
)

func evalExpressions(exps []ast.Expression, env *object.Environment) ([]object.Object, *object.Error) {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if err, ok := evaluated.(*object.Error); ok {
			return result, err
		}
		result = append(result, evaluated)
	}

	return result, nil
}
