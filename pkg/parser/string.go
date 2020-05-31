package parser

import (
	"github.com/uesteibar/lainoa/pkg/ast"
)

func (p *Parser) parseString() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}
