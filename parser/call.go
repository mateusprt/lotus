package parser

import (
	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/errors"
	"github.com/mateusprt/lotus/token"
)

func call(p *Parser) (ast.Expression, error) {
	expr, err := primary(p)
	if err != nil {
		return nil, err
	}

	for {
		if match(p, token.LPAREN) {
			expr, err = finishCall(p, expr)
			if err != nil {
				return nil, err
			}
		} else if match(p, token.DOT) {
			name, _ := consume(p, token.IDENTIFIER, "Expect property name after '.'.")
			expr = &ast.Get{Object: expr, Name: name}
		} else if match(p, token.LBRACKET) {
			index, err := expression(p)
			if err != nil {
				return nil, err
			}
			consume(p, token.RBRACKET, "Expect ']' after index.")
			expr = &ast.Index{Object: expr, Index: index}
		} else {
			break
		}
	}
	return expr, nil
}

func finishCall(p *Parser, callee ast.Expression) (ast.Expression, error) {
	args := make([]ast.Expression, 0)
	if !check(p, token.RPAREN) {
		for {
			if len(args) >= 255 {
				return nil, errors.NewParseError(peek(p), "Can't have more than 255 arguments.")
			}
			arg, err := expression(p)
			if err != nil {
				return nil, err
			}
			args = append(args, arg)
			if !match(p, token.COMMA) {
				break
			}
		}
	}
	rparen, err := consume(p, token.RPAREN, "Expect ')' after arguments.")
	if err != nil {
		return nil, err
	}
	return &ast.Call{
		Callee:    callee,
		RParen:    rparen,
		Arguments: args,
	}, nil
}
