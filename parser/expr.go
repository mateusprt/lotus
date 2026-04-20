package parser

import (
	"strconv"

	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/errors"
	"github.com/mateusprt/lotus/token"
)

func expression(p *Parser) (ast.Expression, error) {
	return assignment(p)
}

func assignment(p *Parser) (ast.Expression, error) {
	// or
	expr, err := or(p)
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

		if get, ok := expr.(*ast.Get); ok {
			return &ast.Set{Object: get.Object, Name: get.Name, Value: value}, nil
		}

		if index, ok := expr.(*ast.Index); ok {
			return &ast.IndexAssign{
				Object: index.Object,
				Index:  index.Index,
				Value:  value,
			}, nil
		}
		return nil, errors.NewParseError(equals, "Invalid assignment target.")
	}
	return expr, nil
}

func or(p *Parser) (ast.Expression, error) {
	expr, err := and(p)
	if err != nil {
		return nil, err
	}

	for match(p, token.OR) {
		operator := previous(p)
		right, err := and(p)
		if err != nil {
			return nil, err
		}
		expr = &ast.Logical{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr, nil
}

func and(p *Parser) (ast.Expression, error) {
	expr, err := equality(p)
	if err != nil {
		return nil, err
	}

	for match(p, token.AND) {
		operator := previous(p)
		right, err := equality(p)
		if err != nil {
			return nil, err
		}
		expr = &ast.Logical{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
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
	return call(p)
}

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
