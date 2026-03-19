package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers
	IDENT   = "IDENT"
	INTEGER = "INTEGER"
	DOUBLE  = "DOUBLE"
	STRING  = "STRING"
	BOOLEAN = "BOOLEAN"

	// Operators
	ASSIGN    = "="
	EQ        = "=="
	NOT_EQ    = "!="
	PLUS      = "+"
	INCREMENT = "++"
	MINUS     = "-"
	DECREMENT = "--"
	DIVIDE    = "/"
	MULTIPLY  = "*"
	BANG      = "!"
	LT        = "<"
	GT        = ">"
	LE        = "<="
	GE        = ">="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	DOT       = "."

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	VAR      = "VAR"
	RETURN   = "RETURN"
	FOR      = "FOR"
	BREAK    = "BREAK"
	CONTINUE = "CONTINUE"
	STRUCT   = "STRUCT"
	CONST    = "CONST"
	IF       = "IF"
	ELSE     = "ELSE"
	PRINT    = "PRINT"
	LEN      = "LEN"
	FIRST    = "FIRST"
	LAST     = "LAST"
	PUSH     = "PUSH"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)
