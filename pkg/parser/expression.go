package parser

import (
	"fmt"

	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/token"
)

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix, exists := p.prefixParseFns[p.curToken.Type]
	if !exists {
		p.addError(fmt.Sprintf("prefix operation %s not recognized", p.curToken.Literal))
		return nil
	}

	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix, exists := p.infixParseFns[p.peekToken.Type]
		if !exists {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)

	}

	return leftExp
}
