package parser

import (
	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/token"
)

func varDeclaration(p *Parser) *ast.VarStmt {
	name, err := consume(p, token.IDENTIFIER, "Expect variable name.")
	if err != nil {
		panic(err)
	}

	var initializer ast.Expression
	if match(p, token.ASSIGN) {
		expr, err := expression(p)
		if err != nil {
			panic(err)
		}
		initializer = expr
	}

	consume(p, token.SEMICOLON, "Expect ';' after variable declaration.")
	return &ast.VarStmt{Name: name, Initializer: initializer}
}
