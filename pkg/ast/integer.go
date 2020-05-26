package ast

import (
	"strconv"

	"github.com/uesteibar/lainoa/pkg/token"
)

type IntegerLiteral struct {
	Token token.Token // token.INT
	Value int64
}

func (i *IntegerLiteral) expressionNode()      {}
func (i *IntegerLiteral) TokenLiteral() string { return i.Token.Literal }
func (i *IntegerLiteral) String() string       { return strconv.Itoa(int(i.Value)) }
