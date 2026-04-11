package debug

import (
	"fmt"

	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/environment"
)

type AstPrinter struct {
	environment *environment.Environment
}

func NewAstPrinter(environment *environment.Environment) *AstPrinter {
	return &AstPrinter{environment: environment}
}

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

func (i *AstPrinter) VisitVariable(expr *ast.Variable) interface{} {
	return environment.Get(i.environment, expr.Name)
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

func (a *AstPrinter) VisitAssign(expr *ast.Assign) interface{} {
	return a.parenthesize("assign "+expr.Name.Lexeme, expr.Value)
}

func (a *AstPrinter) VisitLogical(expr *ast.Logical) interface{} {
	return a.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (a *AstPrinter) VisitCall(expr *ast.Call) interface{} {
	result := a.parenthesize("call", expr.Callee)
	for _, arg := range expr.Arguments {
		result += " " + arg.Accept(a).(string)
	}
	return result
}
