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
	case token.MINUS:
		return evalMinusOperation(right)
	case token.PLUS:
		return evalPlusOperation(right)
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

func evalPlusOperation(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJECT {
		return NULL
	}

	val := right.(*object.Integer).Value
	return &object.Integer{Value: val}
}

func evalMinusOperation(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJECT {
		return NULL
	}

	val := right.(*object.Integer).Value
	return &object.Integer{Value: -val}
}

func evalBoolean(boolean *ast.Boolean) *object.Boolean {
	if boolean.Value {
		return TRUE
	}
	return FALSE
}
