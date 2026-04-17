package ast

type Visitor interface {
	VisitBinary(expr *Binary) interface{}
	VisitGrouping(expr *Grouping) interface{}
	VisitLiteral(expr *Literal) interface{}
	VisitUnary(expr *Unary) interface{}
	VisitVariable(expr *Variable) interface{}
	VisitAssign(expr *Assign) interface{}
	VisitLogical(expr *Logical) interface{}
	VisitCall(expr *Call) interface{}
	VisitGet(expr *Get) interface{}
}
