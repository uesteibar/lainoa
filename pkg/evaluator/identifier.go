package evaluator

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/object"
)

func evalIdentifier(ident *ast.Identifier, env *object.Environment) object.Object {
	val, exists := env.Get(ident.Value)

	if !exists {
		return object.NewError("identifier not found: %s", ident.Value)
	}

	return val
}
