package parser

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/token"
)

func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	index := &ast.IndexExpression{
		Token: p.curToken,
		Left:  left,
	}

	p.nextToken()

	index.Index = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RBRACKET) {
		return nil
	}

	return index
}
