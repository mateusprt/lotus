package parser

import (
	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/token"
)

func forStatement(p *Parser) ast.Stmt {
	consume(p, token.LPAREN, "Expect '(' after 'for'.")

	var initializer ast.Stmt
	if match(p, token.SEMICOLON) {
		initializer = nil
	} else if match(p, token.VAR) {
		initializer = varDeclaration(p)
	} else {
		initializer = expressionStatement(p)
	}

	var condition ast.Expression
	if !check(p, token.SEMICOLON) {
		expr, err := expression(p)
		if err != nil {
			panic(err)
		}
		condition = expr
	}
	consume(p, token.SEMICOLON, "Expect ';' after loop condition.")

	var increment ast.Expression
	if !check(p, token.RPAREN) {
		expr, err := expression(p)
		if err != nil {
			panic(err)
		}
		increment = expr
	}
	consume(p, token.RPAREN, "Expect ')' after for clauses.")

	body := statement(p)

	// Desugar o loop for em um loop while
	if increment != nil {
		body = &ast.BlockStmt{
			Statements: []ast.Stmt{
				body,
				&ast.ExpressionStmt{Expression: increment},
			},
		}
	}

	if condition == nil {
		condition = &ast.Literal{Value: true}
	}

	body = &ast.WhileStmt{
		Condition: condition,
		Body:      body,
	}

	if initializer != nil {
		body = &ast.BlockStmt{
			Statements: []ast.Stmt{
				initializer,
				body,
			},
		}
	}
	return body
}
