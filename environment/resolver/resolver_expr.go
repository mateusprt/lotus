package resolver

import (
	"github.com/mateusprt/lotus/ast"
)

func (r *Resolver) VisitBinary(expr *ast.Binary) interface{} {
}

func (r *Resolver) VisitGrouping(expr *ast.Grouping) interface{} {
}

func (r *Resolver) VisitLiteral(expr *ast.Literal) interface{} {
}

func (r *Resolver) VisitUnary(expr *ast.Unary) interface{} {
}

func (r *Resolver) VisitVariable(expr *ast.Variable) interface{} {
}

func (r *Resolver) VisitAssign(expr *ast.Assign) interface{} {

}

func (r *Resolver) VisitLogical(expr *ast.Logical) interface{} {

}

func (r *Resolver) VisitCall(expr *ast.Call) interface{} {

}
