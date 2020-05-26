package parser

import "github.com/uesteibar/lainoa/pkg/ast"

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()

	// advance to the expression on the right
	p.nextToken()
	exp.Right = p.parseExpression(precedence)

	return exp
}
