package object

import (
	"bytes"
	"strings"
)

type Array struct {
	Elements []Object
}

func (a *Array) Inspect() string {
	var out bytes.Buffer
	out.WriteString("[")

	expressions := []string{}
	for _, el := range a.Elements {
		expressions = append(expressions, el.Inspect())
	}
	out.WriteString(strings.Join(expressions, ", "))

	out.WriteString("]")
	return out.String()
}
func (a *Array) Type() ObjectType { return ARRAY_OBJECT }
