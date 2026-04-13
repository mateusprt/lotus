package interpreter

import (
	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/environment"
)

type LoxFunction struct {
	declaration *ast.FunctionStmt
}

func NewLoxFunction(declaration *ast.FunctionStmt) *LoxFunction {
	return &LoxFunction{declaration: declaration}
}

func (f *LoxFunction) Call(i *Interpreter, arguments []interface{}) interface{} {
	env := environment.NewEnclosed(i.Globals) // ou i.environment, dependendo do contexto

	for i, param := range f.declaration.Params {
		environment.Define(env, param.Lexeme, arguments[i])
	}

	ExecuteBlock(i, f.declaration.Body, env)
	return nil
}

func (f *LoxFunction) Arity() int {
	return len(f.declaration.Params)
}

func (f *LoxFunction) String() string {
	return "<fn " + f.declaration.Name.Lexeme + ">"
}
