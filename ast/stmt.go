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

type BlockStmt struct {
	Statements []Stmt
}

func (b *BlockStmt) Accept(visitor StmtVisitor) {
	visitor.VisitBlockStmt(b)
}

type IfStmt struct {
	Condition Expression
	Then      Stmt
	Else      Stmt
}

func (i *IfStmt) Accept(visitor StmtVisitor) {
	visitor.VisitIfStmt(i)
}

type WhileStmt struct {
	Condition Expression
	Body      Stmt
}

func (w *WhileStmt) Accept(visitor StmtVisitor) {
	visitor.VisitWhileStmt(w)
}

type FunctionStmt struct {
	Name   token.Token
	Params []token.Token
	Body   []Stmt
}

func (f *FunctionStmt) Accept(visitor StmtVisitor) {
	visitor.VisitFunctionStmt(f)
}
