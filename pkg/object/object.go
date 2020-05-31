package object

type ObjectType string

// Object types
var (
	INTEGER_OBJECT      = ObjectType("INTEGER")
	STRING_OBJECT       = ObjectType("STRING")
	BOOLEAN_OBJECT      = ObjectType("BOOLEAN")
	NULL_OBJECT         = ObjectType("NULL")
	RETURN_VALUE_OBJECT = ObjectType("RETURN_VALUE")
	ERROR_OBJECT        = ObjectType("ERROR")
	FUNCTION_OBJECT     = ObjectType("FUNCTION")
)

type Object interface {
	Type() ObjectType
	Inspect() string
}
