package ast

import (
	"fmt"

	"github.com/uesteibar/lainoa/pkg/token"
)

type StringLiteral struct {
	Token token.Token // token.STRING
	Value string
}

func (s *StringLiteral) expressionNode()      {}
func (s *StringLiteral) TokenLiteral() string { return s.Token.Literal }
func (s *StringLiteral) String() string       { return fmt.Sprintf("\"%s\"", s.Value) }
