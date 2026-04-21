package parser

import (
	"github.com/mateusprt/lotus/ast"
	"github.com/mateusprt/lotus/token"
)

func structDeclaration(p *Parser) ast.Stmt {
	name, _ := consume(p, token.IDENTIFIER, "Expect struct name.")
	consume(p, token.LBRACE, "Expect '{' after struct name.")
	var fields []token.Token
	for !check(p, token.RBRACE) && !isAtEnd(p) {
		field, _ := consume(p, token.IDENTIFIER, "Expect field name.")
		fields = append(fields, field)
		consume(p, token.SEMICOLON, "Expect ';' after field declaration.")
	}
	consume(p, token.RBRACE, "Expect '}' after struct body.")
	return &ast.StructStmt{Name: name, Fields: fields}
}
