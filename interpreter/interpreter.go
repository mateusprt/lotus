package interpreter

import (
	"fmt"
	"os"

	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/environment"
	"github.com/mateusprt/lotus/errors"
	"github.com/mateusprt/lotus/token"
)

type Interpreter struct {
	environment *environment.Environment
}

func New(e *environment.Environment) *Interpreter {
	return &Interpreter{environment: e}
}

func (i *Interpreter) Interpret(statements []ast.Stmt) {
	defer func() {
		if r := recover(); r != nil {
			if runtimeErr, ok := r.(*errors.RuntimeError); ok {
				errors.PrintRuntimeError(runtimeErr)
				os.Exit(65)
			} else {
				panic(r)
			}
		}
	}()
	for _, statement := range statements {
		execute(statement, i)
	}
}

func execute(stmt ast.Stmt, interpreter *Interpreter) {
	stmt.Accept(interpreter)
}

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

func (i *Interpreter) VisitExpressionStmt(stmt *ast.ExpressionStmt) {
	evaluate(stmt.Expression, i)
}

func (i *Interpreter) VisitPrintStmt(stmt *ast.PrintStmt) {
	value := evaluate(stmt.Expression, i)
	fmt.Println(stringify(value))
}

func (i *Interpreter) VisitVarStmt(stmt *ast.VarStmt) {
	var value interface{} = nil
	if stmt.Initializer != nil {
		value = evaluate(stmt.Initializer, i)
	}
	environment.Define(i.environment, stmt.Name.Lexeme, value)
}

func (i *Interpreter) VisitAssign(expr *ast.Assign) interface{} {
	value := evaluate(expr.Value, i)
	environment.Assign(i.environment, expr.Name, value)
	return value
}

func evaluate(expr ast.Expression, interpreter *Interpreter) interface{} {
	return expr.Accept(interpreter)
}

func isTruthy(value interface{}) bool {
	if value == nil {
		return false
	}
	if b, ok := value.(bool); ok {
		return b
	}
	return true
}

func isString(value interface{}) bool {
	_, ok := value.(string)
	return ok
}

func isNumber(value interface{}) bool {
	switch value.(type) {
	case int, float32, float64:
		return true
	}
	return false
}

func isEqual(a, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return false
	}
	return a == b
}

func toFloat(value interface{}) float64 {
	switch v := value.(type) {
	case int:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return v
	default:
		return 0
	}
}

func toString(value interface{}) string {
	if s, ok := value.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", value)
}

func checkNumberOperand(operator token.Token, operand interface{}) {
	if isNumber(operand) {
		return
	}
	errors.ThrowRuntimeError(operator, "Operand must be a number.")
}

func checkNumberOperands(operator token.Token, left, right interface{}) {
	if isNumber(left) && isNumber(right) {
		return
	}
	errors.ThrowRuntimeError(operator, "Operands must be numbers.")
}

func stringify(value interface{}) string {
	if value == nil {
		return "nil"
	}
	switch v := value.(type) {
	case float64:
		text := fmt.Sprintf("%g", v)
		return text
	default:
		return fmt.Sprintf("%v", value)
	}
}
