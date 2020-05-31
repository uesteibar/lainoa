package evaluator

import (
	"fmt"

	"github.com/uesteibar/lainoa/pkg/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return object.NewError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}

			default:
				return object.NewError("argument to `len` not supported, got %s", arg.Type())
			}
		},
	},
	"puts": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return object.NewError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				fmt.Println(arg.Value)
				return arg
			case *object.Integer:
				fmt.Println(arg.Value)
				return arg
			case *object.Boolean:
				fmt.Println(arg.Value)
				return arg
			default:
				return object.NewError("argument to `puts` not supported, got %s", arg.Type())
			}
		},
	},
}
