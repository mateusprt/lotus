package parser

import (
	"strconv"

	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/errors"
	"github.com/mateusprt/lotus/token"
)

type Parser struct {
	Tokens  []token.Token
	current int
}

func New(tokens []token.Token) *Parser {
	return &Parser{
		Tokens:  tokens,
		current: 0,
	}
}

func Parse(p *Parser) []ast.Stmt {
	statements := make([]ast.Stmt, 0)
	for !isAtEnd(p) {
		statements = append(statements, declaration(p))
	}
	return statements
}

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
	if match(p, token.VAR) {
		return varDeclaration(p)
	}
	return statement(p)
}

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

func statement(p *Parser) ast.Stmt {
	if match(p, token.PRINT) {
		return printStatement(p)
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

func printStatement(p *Parser) *ast.PrintStmt {
	expr, err := expression(p)
	if err != nil {
		panic(err)
	}
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

func expression(p *Parser) (ast.Expression, error) {
	return assignment(p)
}

func assignment(p *Parser) (ast.Expression, error) {
	expr, err := equality(p)
	if err != nil {
		return nil, err
	}

	if match(p, token.ASSIGN) {
		equals := previous(p)
		value, err := assignment(p)
		if err != nil {
			return nil, err
		}

		if variable, ok := expr.(*ast.Variable); ok {
			token := variable.Name
			return &ast.Assign{Name: token, Value: value}, nil
		}
		return nil, errors.NewParseError(equals, "Invalid assignment target.")
	}
	return expr, nil
}

func equality(p *Parser) (ast.Expression, error) {
	expr, err := comparison(p)
	if err != nil {
		return nil, err
	}

	for match(p, token.NOT_EQ, token.EQ) {
		operator := previous(p)
		right, err := comparison(p)
		if err != nil {
			return nil, err
		}
		expr = &ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr, nil
}

func comparison(p *Parser) (ast.Expression, error) {
	expr, err := term(p)
	if err != nil {
		return nil, err
	}

	for match(p, token.GE, token.GT, token.LE, token.LT) {
		operator := previous(p)
		right, err := term(p)
		if err != nil {
			return nil, err
		}
		expr = &ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr, nil
}

func term(p *Parser) (ast.Expression, error) {
	expr, err := factor(p)
	if err != nil {
		return nil, err
	}

	for match(p, token.MINUS, token.PLUS) {
		operator := previous(p)
		right, err := factor(p)
		if err != nil {
			return nil, err
		}
		expr = &ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr, nil
}

func factor(p *Parser) (ast.Expression, error) {
	expr, err := unary(p)
	if err != nil {
		return nil, err
	}

	for match(p, token.SLASH, token.STAR) {
		operator := previous(p)
		right, err := unary(p)
		if err != nil {
			return nil, err
		}
		expr = &ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr, nil
}

func match(p *Parser, types ...token.TokenType) bool {
	for _, t := range types {
		if check(p, t) {
			advance(p)
			return true
		}
	}
	return false
}

func unary(p *Parser) (ast.Expression, error) {
	if match(p, token.BANG, token.MINUS) {
		operator := previous(p)
		right, err := unary(p)
		if err != nil {
			return nil, err
		}
		return &ast.Unary{
			Operator: operator,
			Right:    right,
		}, nil
	}
	return primary(p)
}

func primary(p *Parser) (ast.Expression, error) {
	if match(p, token.FALSE) {
		return &ast.Literal{Value: false}, nil
	}
	if match(p, token.TRUE) {
		return &ast.Literal{Value: true}, nil
	}
	if match(p, token.NULL) {
		return &ast.Literal{Value: nil}, nil
	}
	if match(p, token.NUMBER, token.STRING) {
		if previous(p).Type == token.NUMBER {
			lexem := previous(p).Literal.(string)
			value, err := strconv.ParseFloat(lexem, 64)
			if err != nil {
				return nil, errors.NewParseError(previous(p), "Invalid number format.")
			}
			return &ast.Literal{Value: value}, nil
		}
		return &ast.Literal{Value: previous(p).Literal}, nil
	}
	if match(p, token.IDENTIFIER) {
		return &ast.Variable{Name: previous(p)}, nil
	}
	if match(p, token.LPAREN) {
		expr, err := expression(p)
		if err != nil {
			return nil, err
		}
		consume(p, token.RPAREN, "Expect ')' after expression.")
		return &ast.Grouping{Expression: expr}, nil
	}
	return nil, errors.NewParseError(peek(p), "Expect expression.")
}

func consume(p *Parser, t token.TokenType, message string) (token.Token, error) {
	if check(p, t) {
		return advance(p), nil
	}
	return token.Token{}, errors.NewParseError(peek(p), message)
}

func check(p *Parser, t token.TokenType) bool {
	if isAtEnd(p) {
		return false
	}
	return peek(p).Type == t
}

func advance(p *Parser) token.Token {
	if !isAtEnd(p) {
		p.current++
	}
	return previous(p)
}

func isAtEnd(p *Parser) bool {
	return peek(p).Type == token.EOF
}

func peek(p *Parser) token.Token {
	return p.Tokens[p.current]
}

func previous(p *Parser) token.Token {
	return p.Tokens[p.current-1]
}

func synchronize(p *Parser) {
	advance(p)

	for !isAtEnd(p) {
		if previous(p).Type == token.SEMICOLON {
			return
		}

		switch peek(p).Type {
		case token.STRUCT,
			token.FUNCTION,
			token.VAR,
			token.CONST,
			token.FOR,
			token.IF,
			token.ELSE,
			token.PRINT,
			token.RETURN:
			return
		}

		advance(p)
	}
}
