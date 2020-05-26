package parser

import (
	"fmt"
	"strconv"

	"github.com/uesteibar/lainoa/pkg/ast"
)

func (p *Parser) parseInteger() ast.Expression {
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		p.addError(fmt.Sprintf("Couldn't parse %q as integer", p.curToken.Literal))
	}
	return &ast.IntegerLiteral{Token: p.curToken, Value: value}
}
