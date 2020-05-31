package evaluator

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/object"
)

func evalFunctionLiteral(fun *ast.FunctionLiteral, env *object.Environment) object.Object {
	return &object.Function{
		Parameters: fun.Parameters,
		Body:       fun.Body,
		Env:        env,
	}
}

func evalFunctionCall(call *ast.CallExpression, env *object.Environment) object.Object {
	fun := Eval(call.Function, env)
	if object.IsError(fun) {
		return fun
	}

	args, err := evalExpressions(call.Arguments, env)
	if err != nil {
		return err
	}

	return applyFunction(fun, args)
}

func evalExpressions(exps []ast.Expression, env *object.Environment) ([]object.Object, *object.Error) {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if err, ok := evaluated.(*object.Error); ok {
			return result, err
		}
		result = append(result, evaluated)
	}

	return result, nil
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {

	case *object.Function:
		env, err := envWithArgs(fn, args)
		if err != nil {
			return err
		}
		evaluated := Eval(fn.Body, env)
		return unwrapReturnValue(evaluated)

	case *object.Builtin:
		return fn.Fn(args...)

	default:
		return object.NewError("expected %s to be a function, got %s", fn.Inspect(), fn.Type())
	}
}

func envWithArgs(fn *object.Function, args []object.Object) (*object.Environment, *object.Error) {
	env := object.NewEnclosedEnvironment(fn.Env)

	for paramIdx, param := range fn.Parameters {
		res := env.Set(param.Value, args[paramIdx])

		if err, ok := res.(*object.Error); ok {
			return env, err
		}
	}

	return env, nil
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}

	return obj
}
