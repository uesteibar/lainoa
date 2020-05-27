package ast

import (
	"bytes"

	"github.com/uesteibar/lainoa/pkg/token"
)

type BlockStatement struct {
	Token      token.Token // token.RBRACE '{'
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
