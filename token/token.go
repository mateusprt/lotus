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
	EOF     = "EOF"
	INTEGER = "INTEGER"
	FLOAT   = "FLOAT"
	STRING  = "STRING"
	BOOLEAN = "BOOLEAN"

	// Operators
	ASSIGN    = "ASSIGN"
	PLUS      = "PLUS"
	MINUS     = "MINUS"
	MULTIPLY  = "MULTIPLY"
	DIVIDE    = "DIVIDE"
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
	VAR      = "VAR"
	CONST    = "CONST"
	IF       = "IF"
	ELSE     = "ELSE"
	FOR      = "FOR"
	BREAK    = "BREAK"
	CONTINUE = "CONTINUE"
	PRINT    = "PRINT"
	FUNCTION = "FUNCTION"
	RETURN   = "RETURN"
	STRUCT   = "STRUCT"
	LEN      = "LEN"
	FIRST    = "FIRST"
	LAST     = "LAST"
	PUSH     = "PUSH"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	NULL     = "NULL"
)
