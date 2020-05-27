package parser

import (
	"fmt"

	"github.com/uesteibar/lainoa/pkg/ast"
	"github.com/uesteibar/lainoa/pkg/token"
)

func (p *Parser) parseFunctionLiteral() ast.Expression {
	fun := &ast.FunctionLiteral{
		Token: p.curToken,
	}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	fun.Parameters = p.parseParameters()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	fun.Body = p.parseBlockStatement()

	return fun
}

func (p *Parser) parseParameters() []*ast.Identifier {
	params := []*ast.Identifier{}

	if p.peekTokenIs(token.RPAREN) { // this means there's no params
		p.nextToken()
		return params
	}

	p.nextToken()

	if !p.curTokenIs(token.IDENT) {
		p.addUnexpectedArgumentError()
	}

	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	params = append(params, ident)

	for p.peekTokenIs(token.COMMA) {
		// skip comma, in "a, b" b is 2 tokens away from a
		p.nextToken()
		p.nextToken()

		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		params = append(params, ident)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return params
}

func (p *Parser) addUnexpectedArgumentError() {
	p.addError(fmt.Sprintf(
		"Function parameters can only be identifiers, found '%s' instead",
		p.curToken.Literal,
	))
}
