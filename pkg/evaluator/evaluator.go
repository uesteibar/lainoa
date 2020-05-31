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

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node.Statements, env)
	case *ast.BlockStatement:
		return evalBlockStatement(node.Statements, env)
	case *ast.ReturnStatement:
		return evalReturnStatement(node, env)
	case *ast.LetStatement:
		return evalLetStatement(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.FunctionLiteral:
		return evalFunctionLiteral(node, env)
	case *ast.CallExpression:
		return evalFunctionCall(node, env)
	case *ast.PrefixExpression:
		return evalPrefix(node, env)
	case *ast.InfixExpression:
		return evalInfix(node, env)
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.AssignExpression:
		return evalAssign(node, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return evalBoolean(node)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	default:
		return nil
	}
}

func evalProgram(statements []ast.Statement, env *object.Environment) object.Object {
	var res object.Object

	for _, stmt := range statements {
		res = Eval(stmt, env)

		if returnValue, ok := res.(*object.ReturnValue); ok {
			return returnValue.Value
		}
		if err, ok := res.(*object.Error); ok {
			return err
		}
	}

	return res
}

func evalBlockStatement(statements []ast.Statement, env *object.Environment) object.Object {
	var res object.Object

	for _, stmt := range statements {
		res = Eval(stmt, env)

		if res != nil && res.Type() == object.RETURN_VALUE_OBJECT {
			return res
		}
		if object.IsError(res) {
			return res
		}
	}

	return res
}
