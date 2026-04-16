package resolver

import (
	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/ds"
	"github.com/mateusprt/lotus/interpreter"
)

type FunctionType int

const (
	NoneFunction FunctionType = iota
	Function
)

type Resolver struct {
	interpreter     *interpreter.Interpreter
	scopes          *ds.Stack[map[string]bool]
	currentFunction FunctionType
}

func New(interpreter *interpreter.Interpreter) *Resolver {
	return &Resolver{interpreter: interpreter, scopes: ds.NewStack[map[string]bool]()}
}

func Resolve(r *Resolver, statements []ast.Stmt) {
	for _, statement := range statements {
		resolveStmt(r, statement)
	}
}
