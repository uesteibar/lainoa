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
	case *ast.InfixExpression:
		return evalInfix(node)
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

func evalInfix(infix *ast.InfixExpression) object.Object {
	left := Eval(infix.Left)
	right := Eval(infix.Right)

	switch {
	case left.Type() == object.INTEGER_OBJECT && right.Type() == object.INTEGER_OBJECT:
		left := left.(*object.Integer)
		right := right.(*object.Integer)
		return evalIntegerInfixExpression(left, infix.Operator, right)
	case infix.Operator == token.EQ:
		return nativeBoolToBoolean(left == right)
	case infix.Operator == token.NOT_EQ:
		return nativeBoolToBoolean(left != right)
	default:
		return NULL
	}
}

func evalIntegerInfixExpression(left *object.Integer, operator string, right *object.Integer) object.Object {
	switch operator {
	case token.PLUS:
		return &object.Integer{Value: left.Value + right.Value}
	case token.MINUS:
		return &object.Integer{Value: left.Value - right.Value}
	case token.ASTERISK:
		return &object.Integer{Value: left.Value * right.Value}
	case token.SLASH:
		return &object.Integer{Value: left.Value / right.Value}
	case token.LT:
		return nativeBoolToBoolean(left.Value < right.Value)
	case token.GT:
		return nativeBoolToBoolean(left.Value > right.Value)
	case token.EQ:
		return nativeBoolToBoolean(left.Value == right.Value)
	case token.NOT_EQ:
		return nativeBoolToBoolean(left.Value != right.Value)
	default:
		return NULL
	}
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
	return nativeBoolToBoolean(boolean.Value)
}

func nativeBoolToBoolean(boolean bool) *object.Boolean {
	if boolean {
		return TRUE
	}
	return FALSE
}
