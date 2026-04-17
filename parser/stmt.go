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

func structDeclaration(p *Parser) ast.Stmt {
	name, _ := consume(p, token.IDENTIFIER, "Expect struct name.")
	consume(p, token.LBRACE, "Expect '{' after struct name.")
	var fields []token.Token
	for !check(p, token.RBRACE) && !isAtEnd(p) {
		field, _ := consume(p, token.IDENTIFIER, "Expect field name.")
		fields = append(fields, field)
		consume(p, token.SEMICOLON, "Expect ';' after field declaration.")
	}
	consume(p, token.RBRACE, "Expect '}' after struct body.")
	return &ast.StructStmt{Name: name, Fields: fields}
}

func function(p *Parser, kind string) *ast.FunctionStmt {
	name, _ := consume(p, token.IDENTIFIER, "Expect "+kind+" name.")
	consume(p, token.LPAREN, "Expect '(' after "+kind+" name.")
	var parameters []token.Token
	if !check(p, token.RPAREN) {
		for {
			if len(parameters) >= 255 {
				panic(errors.NewParseError(peek(p), "Can't have more than 255 parameters."))
			}
			identifier, _ := consume(p, token.IDENTIFIER, "Expect parameter name.")
			parameters = append(parameters, identifier)
			if !match(p, token.COMMA) {
				break
			}
		}
	}
	consume(p, token.RPAREN, "Expect ')' after parameters.")
	consume(p, token.LBRACE, "Expect '{' before "+kind+" body.")
	body := block(p)
	return &ast.FunctionStmt{Name: name, Params: parameters, Body: body}
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

func block(p *Parser) []ast.Stmt {
	var statements []ast.Stmt
	for !check(p, token.RBRACE) && !isAtEnd(p) {
		statements = append(statements, declaration(p))
	}
	consume(p, token.RBRACE, "Expect '}' after block.")
	return statements
}

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

func expressionStatement(p *Parser) *ast.ExpressionStmt {
	expr, err := expression(p)
	if err != nil {
		panic(err)
	}
	consume(p, token.SEMICOLON, "Expect ';' after expression.")
	return &ast.ExpressionStmt{Expression: expr}
}
