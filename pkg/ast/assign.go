package ast

import (
	"bytes"

	"github.com/uesteibar/lainoa/pkg/token"
)

type AssignExpression struct {
	Token token.Token // token.ASSIGN
	Name  *Identifier
	Value Expression
}

func (ae *AssignExpression) expressionNode()      {}
func (ae *AssignExpression) TokenLiteral() string { return ae.Token.Literal }

func (ae *AssignExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ae.Name.String())
	out.WriteString(" = ")
	out.WriteString(ae.Value.String())
	out.WriteString(";")

	return out.String()
}
