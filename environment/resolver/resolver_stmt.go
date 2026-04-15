package resolver

import "github.com/mateusprt/lotus/ast"

func (r *Resolver) VisitExpressionStmt(stmt *ast.ExpressionStmt) {
}

func (r *Resolver) VisitPrintStmt(stmt *ast.PrintStmt) {

}

func (r *Resolver) VisitVarStmt(stmt *ast.VarStmt) {

}

func (r *Resolver) VisitBlockStmt(stmt *ast.BlockStmt) {
	beginScope(r)
	Resolve(r, stmt.Statements)
	endScope(r)
	return
}

func (r *Resolver) VisitIfStmt(stmt *ast.IfStmt) {
}

func (r *Resolver) VisitWhileStmt(stmt *ast.WhileStmt) {
}

func (r *Resolver) VisitFunctionStmt(stmt *ast.FunctionStmt) {
}

func (r *Resolver) VisitReturnStmt(stmt *ast.ReturnStmt) {

}

func Resolve(r *Resolver, statements []ast.Stmt) {
	for _, statement := range statements {
		resolveStmt(r, statement)
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
