package evaluator

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/object"
)

func evalBoolean(boolean *ast.Boolean) *object.Boolean {
	return nativeBoolToBoolean(boolean.Value)
}

func nativeBoolToBoolean(boolean bool) *object.Boolean {
	if boolean {
		return TRUE
	}
	return FALSE
}
