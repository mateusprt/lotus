package token

type TokenType string

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

const (
	// Identifiers
	EOF        = "EOF"
	IDENTIFIER = "IDENTIFIER"
	NUMBER     = "NUMBER"
	STRING     = "STRING"
	BOOLEAN    = "BOOLEAN"

	// Operators
	ASSIGN    = "ASSIGN"
	PLUS      = "PLUS"
	MINUS     = "MINUS"
	STAR      = "STAR"
	SLASH     = "SLASH"
	BANG      = "BANG"
	LT        = "LT"
	GT        = "GT"
	EQ        = "EQ"
	NOT_EQ    = "NOT_EQ"
	INCREMENT = "INCREMENT"
	DECREMENT = "DECREMENT"
	LE        = "LE"
	GE        = "GE"

	// Delimiters
	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"
	COLON     = "COLON"
	DOT       = "DOT"

	LPAREN   = "LPAREN"
	RPAREN   = "RPAREN"
	LBRACE   = "LBRACE"
	RBRACE   = "RBRACE"
	LBRACKET = "LBRACKET"
	RBRACKET = "RBRACKET"

	// Keywords
	VAR    = "VAR"
	CONST  = "CONST"
	IF     = "IF"
	ELSE   = "ELSE"
	FOR    = "FOR"
	WHILE  = "WHILE"
	PRINT  = "PRINT"
	FN     = "FN"
	RETURN = "RETURN"
	STRUCT = "STRUCT"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	NULL   = "NULL"
	AND    = "AND"
	OR     = "OR"
)
