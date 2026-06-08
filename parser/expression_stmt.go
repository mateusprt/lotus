package parser

import (
	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/token"
)

func expressionStatement(p *Parser) *ast.ExpressionStmt {
	expr, err := expression(p)
	if err != nil {
		panic(err)
	}
	if _, err := consume(p, token.SEMICOLON, "Expect ';' after expression."); err != nil {
		panic(err)
	}

	return &ast.ExpressionStmt{Expression: expr}
}
