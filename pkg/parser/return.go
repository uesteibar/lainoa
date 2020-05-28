package parser

import (
	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/token"
)

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)

	for p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
