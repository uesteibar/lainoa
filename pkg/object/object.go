package object

type ObjectType string

// Object types
var (
	INTEGER_OBJECT      = ObjectType("INTEGER")
	STRING_OBJECT       = ObjectType("STRING")
	BOOLEAN_OBJECT      = ObjectType("BOOLEAN")
	NIL_OBJECT          = ObjectType("NIL")
	RETURN_VALUE_OBJECT = ObjectType("RETURN_VALUE")
	ERROR_OBJECT        = ObjectType("ERROR")
	FUNCTION_OBJECT     = ObjectType("FUNCTION")
	BUILTIN_OBJECT      = ObjectType("BUILTIN")
	ARRAY_OBJECT        = ObjectType("ARRAY")
)

type Object interface {
	Type() ObjectType
	Inspect() string
}
