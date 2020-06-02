package parser

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/token"
)

func (p *Parser) parseArray() ast.Expression {
	array := &ast.ArrayExpression{
		Token:       p.curToken,
		Expressions: []ast.Expression{},
	}

	p.nextToken()
	if p.curTokenIs(token.RBRACKET) {
		return array
	}

	exp := p.parseExpression(LOWEST)
	array.Expressions = append(array.Expressions, exp)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()

		exp := p.parseExpression(LOWEST)
		array.Expressions = append(array.Expressions, exp)
	}

	if !p.expectPeek(token.RBRACKET) {
		return nil
	}

	return array
}
