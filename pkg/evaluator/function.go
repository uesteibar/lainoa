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

func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		env, err := envWithArgs(fn.Parameters, args, fn.Env)
		if err != nil {
			return err
		}
		if len(args) < len(fn.Parameters) {
			return curryFunction(fn, env, args)
		}

		evaluated := Eval(fn.Body, env)
		return unwrapReturnValue(evaluated)
	case *object.CurriedFunction:
		env, err := envWithArgs(fn.ParametersLeft, args, fn.Env)
		if err != nil {
			return err
		}
		if len(args) < len(fn.ParametersLeft) {
			return recurryFunction(fn, env, args)
		}

		evaluated := Eval(fn.Fn.Body, env)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		return fn.Fn(args...)

	default:
		return object.NewError("expected %s to be a function, got %s", fn.Inspect(), fn.Type())
	}
}

func curryFunction(fn *object.Function, env *object.Environment, args []object.Object) *object.CurriedFunction {
	return &object.CurriedFunction{
		Fn:             fn,
		Env:            env,
		ParametersLeft: fn.Parameters[len(args):len(fn.Parameters)],
	}
}

func recurryFunction(cur *object.CurriedFunction, env *object.Environment, args []object.Object) *object.CurriedFunction {
	return &object.CurriedFunction{
		Fn:             cur.Fn,
		Env:            env,
		ParametersLeft: cur.ParametersLeft[len(args):len(cur.ParametersLeft)],
	}
}

func envWithArgs(params []*ast.Identifier, args []object.Object, env *object.Environment) (*object.Environment, *object.Error) {
	newEnv := object.NewEnclosedEnvironment(env)

	for paramIdx, param := range params {
		if paramIdx >= len(args) {
			return newEnv, nil
		}
		res := newEnv.Set(param.Value, args[paramIdx])

		if err, ok := res.(*object.Error); ok {
			return newEnv, err
		}
	}

	return newEnv, nil
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}

	return obj
}
