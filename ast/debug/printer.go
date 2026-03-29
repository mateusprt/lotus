package debug

import (
	"fmt"

	"github.com/mateusprt/lotus/ast"
)

type AstPrinter struct{}

func (a *AstPrinter) VisitBinary(expr *ast.Binary) interface{} {
	return a.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (a *AstPrinter) VisitGrouping(expr *ast.Grouping) interface{} {
	return a.parenthesize("group", expr.Expression)
}

func (a *AstPrinter) VisitLiteral(expr *ast.Literal) interface{} {
	if expr.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", expr.Value)
}

func (a *AstPrinter) VisitUnary(expr *ast.Unary) interface{} {
	return a.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (a *AstPrinter) parenthesize(name string, exprs ...ast.Expression) string {
	result := "(" + name
	for _, expr := range exprs {
		result += " " + expr.Accept(a).(string)
	}
	result += ")"
	return result
}

func (a *AstPrinter) Print(expr ast.Expression) string {
	return expr.Accept(a).(string)
}
