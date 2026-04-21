package parser

import (
	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/token"
)

func ifStatement(p *Parser) *ast.IfStmt {
	consume(p, token.LPAREN, "Expect '(' after 'if'.")
	condition, err := expression(p)
	if err != nil {
		panic(err)
	}
	consume(p, token.RPAREN, "Expect ')' after if condition.")
	consume(p, token.LBRACE, "Expect '{' before if body.")

	thenBranch := statement(p)

	consume(p, token.RBRACE, "Expect '}' after if body.")
	var elseBranch ast.Stmt
	if match(p, token.ELSE) {
		consume(p, token.LBRACE, "Expect '{' before if body.")
		elseBranch = statement(p)
		consume(p, token.RBRACE, "Expect '}' after if body.")
	}

	return &ast.IfStmt{
		Condition: condition,
		Then:      thenBranch,
		Else:      elseBranch,
	}
}
