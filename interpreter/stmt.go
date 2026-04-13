package interpreter

import (
	"fmt"

	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/environment"
)

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

func (i *Interpreter) VisitBlockStmt(stmt *ast.BlockStmt) {
	env := environment.NewEnclosed(i.environment)
	ExecuteBlock(i, stmt.Statements, env)
}

func (i *Interpreter) VisitIfStmt(stmt *ast.IfStmt) {
	if isTruthy(evaluate(stmt.Condition, i)) {
		execute(stmt.Then, i)
	} else if stmt.Else != nil {
		execute(stmt.Else, i)
	}
}

func (i *Interpreter) VisitWhileStmt(stmt *ast.WhileStmt) {
	for isTruthy(evaluate(stmt.Condition, i)) {
		execute(stmt.Body, i)
	}
}

func (i *Interpreter) VisitFunctionStmt(stmt *ast.FunctionStmt) {
	function := NewLoxFunction(stmt)
	environment.Define(i.environment, stmt.Name.Lexeme, function)
}
