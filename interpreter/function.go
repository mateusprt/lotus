package interpreter

import (
	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/environment"
	"github.com/mateusprt/lotus/errors"
)

type Function struct {
	declaration *ast.FunctionStmt
	closure     *environment.Environment
}

func NewFunction(declaration *ast.FunctionStmt, closure *environment.Environment) *Function {
	return &Function{declaration: declaration, closure: closure}
}

func (f *Function) Call(i *Interpreter, arguments []interface{}) (returnValue interface{}) {
	env := environment.NewEnclosed(f.closure)
	for i, param := range f.declaration.Params {
		environment.Define(env, param.Lexeme, arguments[i])
	}

	defer func() {
		if r := recover(); r != nil {
			if returnErr, ok := r.(*errors.ReturnError); ok {
				returnValue = returnErr.Value
			} else {
				panic(r)
			}
		}
	}()

	ExecuteBlock(i, f.declaration.Body, env)
	return
}

func (f *Function) Arity() int {
	return len(f.declaration.Params)
}

func (f *Function) String() string {
	return "<fn " + f.declaration.Name.Lexeme + ">"
}
