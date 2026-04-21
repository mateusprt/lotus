package parser

import (
	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/token"
)

func returnStatement(p *Parser) ast.Stmt {
	keyword := previous(p)
	var value ast.Expression
	if !check(p, token.SEMICOLON) {
		expr, err := expression(p)
		if err != nil {
			panic(err)
		}
		value = expr
	}
	consume(p, token.SEMICOLON, "Expect ';' after return value.")
	return &ast.ReturnStmt{Keyword: keyword, Value: value}
}
