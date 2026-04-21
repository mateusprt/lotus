package parser

import (
	"strconv"

	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/errors"
	"github.com/mateusprt/lotus/token"
)

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
	if match(p, token.LBRACKET) {
		var elements []ast.Expression
		if !check(p, token.RBRACKET) {
			element, err := expression(p)
			if err != nil {
				return nil, err
			}
			elements = append(elements, element)
			for match(p, token.COMMA) {
				element, err := expression(p)
				if err != nil {
					return nil, err
				}
				elements = append(elements, element)
			}
		}
		consume(p, token.RBRACKET, "Expect ']' after array elements.")
		return &ast.ArrayLiteral{Elements: elements}, nil
	}

	return nil, errors.NewParseError(peek(p), "Expect expression.")
}
