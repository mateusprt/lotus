package parser

import (
	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/token"
)

func printStatement(p *Parser) *ast.PrintStmt {
	consume(p, token.LPAREN, "Expect '(' after print.")
	expr, err := expression(p)
	if err != nil {
		panic(err)
	}
	consume(p, token.RPAREN, "Expect ')' after value.")
	consume(p, token.SEMICOLON, "Expect ';' after value.")
	return &ast.PrintStmt{Expression: expr}
}
