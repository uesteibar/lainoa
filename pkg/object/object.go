package object

type ObjectType string

// Object types
var (
	INTEGER_OBJECT = ObjectType("INTEGER")
	BOOLEAN_OBJECT = ObjectType("BOOLEAN")
	NULL_OBJECT    = ObjectType("NULL")
)

type Object interface {
	Type() ObjectType
	Inspect() string
}
