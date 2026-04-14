package ast

type StmtVisitor interface {
	VisitExpressionStmt(stmt *ExpressionStmt)
	VisitPrintStmt(stmt *PrintStmt)
	VisitVarStmt(stmt *VarStmt)
	VisitBlockStmt(stmt *BlockStmt)
	VisitIfStmt(stmt *IfStmt)
	VisitWhileStmt(stmt *WhileStmt)
	VisitFunctionStmt(stmt *FunctionStmt)
	VisitReturnStmt(stmt *ReturnStmt)
}
