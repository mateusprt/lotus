package resolver

import (
	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/errors"
	"github.com/mateusprt/lotus/token"
)

func (r *Resolver) VisitExpressionStmt(stmt *ast.ExpressionStmt) {
	resolveExpr(r, stmt.Expression)
}

func (r *Resolver) VisitPrintStmt(stmt *ast.PrintStmt) {
	resolveExpr(r, stmt.Expression)
}

func (r *Resolver) VisitVarStmt(stmt *ast.VarStmt) {
	declare(r, stmt.Name)
	if stmt.Initializer != nil {
		resolveExpr(r, stmt.Initializer)
	}
	define(r, stmt.Name)
}

func declare(r *Resolver, name token.Token) {
	if r.scopes.IsEmpty() {
		return
	}

	scope := r.scopes.Peek()
	if _, ok := scope[name.Lexeme]; ok {
		errors.Error("Variable with this name already declared in this scope.")
	}
	scope[name.Lexeme] = false
}

func define(r *Resolver, name token.Token) {
	if r.scopes.IsEmpty() {
		return
	}

	r.scopes.Peek()[name.Lexeme] = true
}

func (r *Resolver) VisitBlockStmt(stmt *ast.BlockStmt) {
	beginScope(r)
	Resolve(r, stmt.Statements)
	endScope(r)
}

func (r *Resolver) VisitIfStmt(stmt *ast.IfStmt) {
	resolveExpr(r, stmt.Condition)
	resolveStmt(r, stmt.Then)
	if stmt.Else != nil {
		resolveStmt(r, stmt.Else)
	}
}

func (r *Resolver) VisitWhileStmt(stmt *ast.WhileStmt) {
	resolveExpr(r, stmt.Condition)
	resolveStmt(r, stmt.Body)
}

func (r *Resolver) VisitFunctionStmt(stmt *ast.FunctionStmt) {
	declare(r, stmt.Name)
	define(r, stmt.Name)

	resolveFunction(r, stmt, Function)
}

func (r *Resolver) VisitReturnStmt(stmt *ast.ReturnStmt) {
	if stmt.Value != nil {
		resolveExpr(r, stmt.Value)
	}
}

func resolveStmt(r *Resolver, stmt ast.Stmt) {
	stmt.Accept(r)
}

func resolveExpr(r *Resolver, expr ast.Expression) {
	expr.Accept(r)
}

func beginScope(r *Resolver) {
	r.scopes.Push(map[string]bool{})
}

func endScope(r *Resolver) {
	r.scopes.Pop()
}

func resolveFunction(r *Resolver, function *ast.FunctionStmt, functionType FunctionType) {
	enclosingFunction := r.currentFunction
	r.currentFunction = functionType

	beginScope(r)
	for _, param := range function.Params {
		declare(r, param)
		define(r, param)
	}
	Resolve(r, function.Body)
	endScope(r)
	r.currentFunction = enclosingFunction
}

func (r *Resolver) VisitStructStmt(stmt *ast.StructStmt) {
	declare(r, stmt.Name)
	define(r, stmt.Name)
}
