package ast

import "github.com/mateusprt/lotus/token"

type Expression interface {
	Accept(visitor Visitor) interface{}
}

type Binary struct {
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (b *Binary) Accept(visitor Visitor) interface{} {
	return visitor.VisitBinary(b)
}

type Grouping struct {
	Expression Expression
}

func (g *Grouping) Accept(visitor Visitor) interface{} {
	return visitor.VisitGrouping(g)
}

type Literal struct {
	Value interface{}
}

func (l *Literal) Accept(visitor Visitor) interface{} {
	return visitor.VisitLiteral(l)
}

type Unary struct {
	Operator token.Token
	Right    Expression
}

func (u *Unary) Accept(visitor Visitor) interface{} {
	return visitor.VisitUnary(u)
}
