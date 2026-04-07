package ast

import "github.com/mateusprt/lotus/token"

type Stmt interface {
	Accept(visitor StmtVisitor)
}

type ExpressionStmt struct {
	Expression Expression
}

func (e *ExpressionStmt) Accept(visitor StmtVisitor) {
	visitor.VisitExpressionStmt(e)
}

type PrintStmt struct {
	Expression Expression
}

func (p *PrintStmt) Accept(visitor StmtVisitor) {
	visitor.VisitPrintStmt(p)
}

type VarStmt struct {
	Name        token.Token
	Initializer Expression
}

func (v *VarStmt) Accept(visitor StmtVisitor) {
	visitor.VisitVarStmt(v)
}
