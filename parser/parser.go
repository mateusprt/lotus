package parser

import (
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

func match(p *Parser, types ...token.TokenType) bool {
	for _, t := range types {
		if check(p, t) {
			advance(p)
			return true
		}
	}
	return false
}

func synchronize(p *Parser) {
	advance(p)

	for !isAtEnd(p) {
		if previous(p).Type == token.SEMICOLON {
			return
		}

		switch peek(p).Type {
		case token.STRUCT,
			token.FN,
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
