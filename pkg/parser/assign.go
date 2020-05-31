package parser

import (
	"github.com/uesteibar/lainoa/pkg/ast"
)

func (p *Parser) parseAssignExpression(ident *ast.Identifier) ast.Expression {
	a := &ast.AssignExpression{
		Token: p.curToken,
		Name:  ident,
	}

	// advance to the expression on the right
	p.nextToken()
	a.Value = p.parseExpression(LOWEST)

	return a
}
