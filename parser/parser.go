package parser

import (
	"strconv"

	"github.com/mateusprt/lotus/ast"
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
		statements = append(statements, statement(p))
	}
	return statements
}

func statement(p *Parser) ast.Stmt {
	if match(p, token.PRINT) {
		return printStatement(p)
	}
	return expressionStatement(p)
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
	return equality(p)
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
				return nil, NewParseError(previous(p), "Invalid number format.")
			}
			return &ast.Literal{Value: value}, nil
		}
		return &ast.Literal{Value: previous(p).Literal}, nil
	}
	if match(p, token.LPAREN) {
		expr, err := expression(p)
		if err != nil {
			return nil, err
		}
		consume(p, token.RPAREN, "Expect ')' after expression.")
		return &ast.Grouping{Expression: expr}, nil
	}
	return nil, NewParseError(peek(p), "Expect expression.")
}

func consume(p *Parser, t token.TokenType, message string) (token.Token, error) {
	if check(p, t) {
		return advance(p), nil
	}
	return token.Token{}, NewParseError(peek(p), message)
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
