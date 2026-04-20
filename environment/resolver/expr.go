package resolver

import (
	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/token"
)

func (r *Resolver) VisitBinary(expr *ast.Binary) interface{} {
	resolveExpr(r, expr.Left)
	resolveExpr(r, expr.Right)
	return nil
}

func (r *Resolver) VisitGrouping(expr *ast.Grouping) interface{} {
	resolveExpr(r, expr.Expression)
	return nil
}

func (r *Resolver) VisitLiteral(expr *ast.Literal) interface{} {
	return nil
}

func (r *Resolver) VisitUnary(expr *ast.Unary) interface{} {
	resolveExpr(r, expr.Right)
	return nil
}

func (r *Resolver) VisitVariable(expr *ast.Variable) interface{} {
	if !r.scopes.IsEmpty() && r.scopes.Peek()[expr.Name.Lexeme] == false {
		panic("Cannot read local variable in its own initializer.")
	}

	resolveLocal(r, expr, expr.Name)
	return nil
}

func (r *Resolver) VisitAssign(expr *ast.Assign) interface{} {
	resolveExpr(r, expr.Value)
	resolveLocal(r, expr, expr.Name)
	return nil
}

func (r *Resolver) VisitLogical(expr *ast.Logical) interface{} {
	resolveExpr(r, expr.Left)
	resolveExpr(r, expr.Right)
	return nil
}

func (r *Resolver) VisitCall(expr *ast.Call) interface{} {
	resolveExpr(r, expr.Callee)
	for _, arg := range expr.Arguments {
		resolveExpr(r, arg)
	}
	return nil
}

func (r *Resolver) VisitSet(expr *ast.Set) interface{} {
	resolveExpr(r, expr.Value)
	resolveExpr(r, expr.Object)
	return nil
}

func resolveLocal(r *Resolver, expr ast.Expression, name token.Token) {
	for i := r.scopes.Size() - 1; i >= 0; i-- {
		if _, ok := r.scopes.Get(i)[name.Lexeme]; ok {
			r.interpreter.Resolve(expr, r.scopes.Size()-1-i)
			return
		}
	}
}

func (r *Resolver) VisitArrayLiteral(expr *ast.ArrayLiteral) interface{} {
	for _, element := range expr.Elements {
		resolveExpr(r, element)
	}
	return nil
}

func (r *Resolver) VisitIndex(expr *ast.Index) interface{} {
	resolveExpr(r, expr.Object)
	resolveExpr(r, expr.Index)
	return nil
}

func (r *Resolver) VisitIndexAssign(expr *ast.IndexAssign) interface{} {
	resolveExpr(r, expr.Object)
	resolveExpr(r, expr.Index)
	resolveExpr(r, expr.Value)
	return nil
}
