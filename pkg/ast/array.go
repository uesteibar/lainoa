package ast

import (
	"bytes"
	"strings"

	"github.com/uesteibar/lainoa/pkg/token"
)

type ArrayExpression struct {
	Token       token.Token // token.RBRACKET '['
	Expressions []Expression
}

func (ae *ArrayExpression) expressionNode()      {}
func (ae *ArrayExpression) TokenLiteral() string { return ae.Token.Literal }

func (ae *ArrayExpression) String() string {
	var out bytes.Buffer

	out.WriteString("[")
	expressions := []string{}
	for _, e := range ae.Expressions {
		expressions = append(expressions, e.String())
	}
	out.WriteString(strings.Join(expressions, ", "))
	out.WriteString("]")

	return out.String()
}
