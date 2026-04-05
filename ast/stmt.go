package ast

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
