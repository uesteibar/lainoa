package evaluator

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/object"
)

func evalIndexOperation(index *ast.IndexExpression, env *object.Environment) object.Object {
	left := Eval(index.Left, env)
	if object.IsError(left) {
		return left
	}
	i := Eval(index.Index, env)
	if object.IsError(i) {
		return i
	}

	switch left := left.(type) {
	case *object.Array:
		return evalArrayIndex(left, i)
	default:
		return object.NewError("type %s doesn't support index operations", left.Type())
	}
}

func evalArrayIndex(array *object.Array, i object.Object) object.Object {
	switch i := i.(type) {
	case *object.Integer:
		if int(i.Value) < len(array.Elements) && i.Value >= 0 {
			return array.Elements[i.Value]
		}

		return NIL
	default:
		return object.NewError("expected %s as index for array, got %s", object.INTEGER_OBJECT, i.Type())
	}
}
