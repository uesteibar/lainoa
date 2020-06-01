package ast

import (
	"github.com/uesteibar/lainoa/pkg/token"
)

type NilLiteral struct {
	Token token.Token // token.NIL
}

func (n *NilLiteral) expressionNode()      {}
func (n *NilLiteral) TokenLiteral() string { return n.Token.Literal }
func (n *NilLiteral) String() string       { return "nil" }
