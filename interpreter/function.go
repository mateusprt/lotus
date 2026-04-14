package interpreter

import (
	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/environment"
	"github.com/mateusprt/lotus/errors"
)

type LoxFunction struct {
	declaration *ast.FunctionStmt
}

func NewLoxFunction(declaration *ast.FunctionStmt) *LoxFunction {
	return &LoxFunction{declaration: declaration}
}

func (f *LoxFunction) Call(i *Interpreter, arguments []interface{}) interface{} {
	env := environment.NewEnclosed(i.Globals)

	for i, param := range f.declaration.Params {
		environment.Define(env, param.Lexeme, arguments[i])
	}

	var returnValue interface{} = nil
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
	return returnValue
}

func (f *LoxFunction) Arity() int {
	return len(f.declaration.Params)
}

func (f *LoxFunction) String() string {
	return "<fn " + f.declaration.Name.Lexeme + ">"
}
