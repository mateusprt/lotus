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
	Globals     *environment.Environment
	environment *environment.Environment
	locals      map[ast.Expression]int
}

func New(e *environment.Environment) *Interpreter {
	return &Interpreter{Globals: e, environment: e, locals: make(map[ast.Expression]int)}
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

func ExecuteBlock(interpreter *Interpreter, statements []ast.Stmt, env *environment.Environment) {
	previous := interpreter.environment
	interpreter.environment = env
	defer func() {
		interpreter.environment = previous
	}()

	for _, statement := range statements {
		execute(statement, interpreter)
	}
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
	if n, ok := value.(int); ok {
		return n != 0
	}
	if n, ok := value.(float64); ok {
		return n != 0
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

func (i *Interpreter) Resolve(expr ast.Expression, depth int) {
	i.locals[expr] = depth
}
