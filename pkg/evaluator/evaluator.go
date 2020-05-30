package evaluator

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node.Statements)
	case *ast.BlockStatement:
		return evalBlockStatement(node.Statements)
	case *ast.ReturnStatement:
		return evalReturnStatement(node)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.PrefixExpression:
		return evalPrefix(node)
	case *ast.InfixExpression:
		return evalInfix(node)
	case *ast.IfExpression:
		return evalIfExpression(node)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return evalBoolean(node)
	default:
		return nil
	}
}

func evalProgram(statements []ast.Statement) object.Object {
	var res object.Object

	for _, stmt := range statements {
		res = Eval(stmt)

		if returnValue, ok := res.(*object.ReturnValue); ok {
			return returnValue.Value
		}
	}

	return res
}

func evalBlockStatement(statements []ast.Statement) object.Object {
	var res object.Object

	for _, stmt := range statements {
		res = Eval(stmt)

		if res != nil && res.Type() == object.RETURN_VALUE_OBJECT {
			return res
		}
	}

	return res
}
