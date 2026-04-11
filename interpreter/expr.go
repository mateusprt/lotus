package interpreter

import (
	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/environment"
	"github.com/mateusprt/lotus/errors"
	"github.com/mateusprt/lotus/token"
)

func (i *Interpreter) VisitBinary(expr *ast.Binary) interface{} {
	left := evaluate(expr.Left, i)
	right := evaluate(expr.Right, i)

	switch expr.Operator.Type {
	case token.GT:
		checkNumberOperands(expr.Operator, left, right)
		return toFloat(left) > toFloat(right)
	case token.GE:
		checkNumberOperands(expr.Operator, left, right)
		return toFloat(left) >= toFloat(right)
	case token.LT:
		checkNumberOperands(expr.Operator, left, right)
		return toFloat(left) < toFloat(right)
	case token.LE:
		checkNumberOperands(expr.Operator, left, right)
		return toFloat(left) <= toFloat(right)
	case token.EQ:
		return isEqual(left, right)
	case token.NOT_EQ:
		return !isEqual(left, right)
	case token.MINUS:
		checkNumberOperands(expr.Operator, left, right)
		return toFloat(left) - toFloat(right)
	case token.PLUS:
		if isNumber(left) && isNumber(right) {
			return toFloat(left) + toFloat(right)
		}
		if isString(left) && isString(right) {
			return toString(left) + toString(right)
		}
		errors.ThrowRuntimeError(expr.Operator, "Operands must be two numbers or two strings.")
	case token.SLASH:
		checkNumberOperands(expr.Operator, left, right)
		return toFloat(left) / toFloat(right)
	case token.STAR:
		checkNumberOperands(expr.Operator, left, right)
		return toFloat(left) * toFloat(right)
	}
	return nil
}

func (i *Interpreter) VisitGrouping(expr *ast.Grouping) interface{} {
	return evaluate(expr.Expression, i)
}

func (i *Interpreter) VisitLiteral(expr *ast.Literal) interface{} {
	return expr.Value
}

func (i *Interpreter) VisitUnary(expr *ast.Unary) interface{} {
	right := evaluate(expr.Right, i)

	switch expr.Operator.Type {
	case token.MINUS:
		checkNumberOperand(expr.Operator, right)
		return -toFloat(right)
	case token.BANG:
		return !isTruthy(right)
	}
	return nil
}

func (i *Interpreter) VisitVariable(expr *ast.Variable) interface{} {
	return environment.Get(i.environment, expr.Name)
}

func (i *Interpreter) VisitAssign(expr *ast.Assign) interface{} {
	value := evaluate(expr.Value, i)
	environment.Assign(i.environment, expr.Name, value)
	return value
}

func (i *Interpreter) VisitLogical(expr *ast.Logical) interface{} {
	left := evaluate(expr.Left, i)

	if expr.Operator.Type == token.OR {
		if isTruthy(left) {
			return left
		}
	} else {
		if !isTruthy(left) {
			return left
		}
	}

	return evaluate(expr.Right, i)
}
