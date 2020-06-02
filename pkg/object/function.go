package object

import (
	"bytes"
	"strings"

	"github.com/uesteibar/lainoa/pkg/ast"
)

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJECT }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type CurriedFunction struct {
	Fn             *Function
	Env            *Environment
	ParametersLeft []*ast.Identifier
}

func (f *CurriedFunction) Type() ObjectType { return CURRIED_FUNCTION_OBJECT }
func (f *CurriedFunction) Inspect() string {
	var out bytes.Buffer

	out.WriteString("Curried Function: ")
	out.WriteString(f.Fn.Inspect())

	return out.String()
}
