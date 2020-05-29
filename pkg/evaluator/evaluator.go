package evaluator

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/object"
	"github.com/uesteibar/lainoa/pkg/token"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.PrefixExpression:
		return evalPrefix(node)
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

func evalPrefix(prefix *ast.PrefixExpression) object.Object {
	right := Eval(prefix.Right)

	switch prefix.Token.Type {
	case token.BANG:
		return evalBangOperation(right)
	}

	return nil
}

func evalBangOperation(right object.Object) *object.Boolean {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalBoolean(boolean *ast.Boolean) *object.Boolean {
	if boolean.Value {
		return TRUE
	}
	return FALSE
}
