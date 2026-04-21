package parser

import (
	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/errors"
	"github.com/mateusprt/lotus/token"
)

func declaration(p *Parser) ast.Stmt {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(*errors.ParseError); ok {
				synchronize(p)
			} else {
				panic(r)
			}
		}
	}()
	if match(p, token.STRUCT) {
		return structDeclaration(p)
	}
	if match(p, token.FN) {
		return function(p, "function")
	}
	if match(p, token.VAR) {
		return varDeclaration(p)
	}
	return statement(p)
}

func statement(p *Parser) ast.Stmt {
	if match(p, token.FOR) {
		return forStatement(p)
	}
	if match(p, token.IF) {
		return ifStatement(p)
	}
	if match(p, token.RETURN) {
		return returnStatement(p)
	}
	if match(p, token.PRINT) {
		return printStatement(p)
	}
	if match(p, token.WHILE) {
		return whileStatement(p)
	}
	if match(p, token.LBRACE) {
		return &ast.BlockStmt{Statements: block(p)}
	}
	return expressionStatement(p)
}

func block(p *Parser) []ast.Stmt {
	var statements []ast.Stmt
	for !check(p, token.RBRACE) && !isAtEnd(p) {
		statements = append(statements, declaration(p))
	}
	consume(p, token.RBRACE, "Expect '}' after block.")
	return statements
}
