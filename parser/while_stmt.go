package parser

import (
	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/token"
)

func whileStatement(p *Parser) ast.Stmt {
	consume(p, token.LPAREN, "Expect '(' after 'while'.")
	condition, err := expression(p)
	if err != nil {
		panic(err)
	}
	consume(p, token.RPAREN, "Expect ')' after condition.")
	body := statement(p)
	return &ast.WhileStmt{
		Condition: condition,
		Body:      body,
	}
}
