package evaluator

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/object"
)

func evalIdentifier(ident *ast.Identifier, env *object.Environment) object.Object {
	if val, exists := builtins[ident.Value]; exists {
		return val
	}

	if val, exists := env.Get(ident.Value); exists {
		return val
	}

	return object.NewError("identifier not found: %s", ident.Value)
}
