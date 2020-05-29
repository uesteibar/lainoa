package evaluator

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/object"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return evalBoolean(node)
	default:
		return nil
	}
}

func evalStatements(statements []ast.Statement) object.Object {
	var res object.Object

	for _, stmt := range statements {
		res = Eval(stmt)
	}

	return res
}

func evalBoolean(boolean *ast.Boolean) *object.Boolean {
	if boolean.Value {
		return TRUE
	} else {
		return FALSE
	}
}
