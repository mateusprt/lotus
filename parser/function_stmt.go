package parser

import (
	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/errors"
	"github.com/mateusprt/lotus/token"
)

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
