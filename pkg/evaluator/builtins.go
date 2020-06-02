package evaluator

import (
	"fmt"
	"strconv"

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
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}

			default:
				return object.NewError("argument to `len` not supported, got %s", arg.Type())
			}
		},
	},
	"to_string": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return object.NewError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return arg
			case *object.Integer:
				return &object.String{Value: strconv.Itoa(int(arg.Value))}
			default:
				return object.NewError("argument to `to_string` not supported, got %s", arg.Type())
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
			case *object.Nil:
				fmt.Println("nil")
				return arg
			default:
				return object.NewError("argument to `puts` not supported, got %s", arg.Type())
			}
		},
	},

	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return object.NewError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJECT {
				return object.NewError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"head": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return object.NewError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) > 0 {
					return arg.Elements[0]
				}
				return NIL
			default:
				return object.NewError("argument to `first` must be ARRAY, got %s", arg.Type())
			}
		},
	},
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return object.NewError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJECT {
				return object.NewError("argument to `rest` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NIL
		},
	},
}
