package resolver

import (
	"github.com/mateusprt/lotus/ds"
	"github.com/mateusprt/lotus/interpreter"
)

type Resolver struct {
	interpreter *interpreter.Interpreter
	scopes      *ds.Stack[map[string]bool]
}

func New(interpreter *interpreter.Interpreter) *Resolver {
	return &Resolver{interpreter: interpreter, scopes: ds.NewStack[map[string]bool]()}
}
