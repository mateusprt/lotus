package parser

import (
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
