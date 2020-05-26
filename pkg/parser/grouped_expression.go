package parser

import (
	"github.com/uesteibar/lainoa/pkg/token"

	"github.com/uesteibar/lainoa/pkg/ast"
)

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}
