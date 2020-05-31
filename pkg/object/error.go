package object

import "fmt"

type Error struct {
	Message string
}

func (e *Error) Inspect() string  { return fmt.Sprintf("ERROR: %s", e.Message) }
func (e *Error) Type() ObjectType { return ERROR_OBJECT }

func NewError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

func IsError(obj Object) bool {
	if obj != nil {
		return obj.Type() == ERROR_OBJECT
	}
	return false
}
