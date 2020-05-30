package parser

import (
	"fmt"

	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/token"
)

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()

		block.Statements = append(block.Statements, stmt)

		p.nextToken()
	}

	if !p.curTokenIs(token.RBRACE) {
		p.addError(fmt.Sprintf("expected } at the end of the block, got %s instead", p.curToken.Type))
		return block
	}

	return block
}
