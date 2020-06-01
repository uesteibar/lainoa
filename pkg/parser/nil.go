package parser

import (
	"github.com/uesteibar/lainoa/pkg/ast"
)

func (p *Parser) parseNil() ast.Expression {
	return &ast.NilLiteral{Token: p.curToken}
}
