package scanner

import (
	"fmt"

	"github.com/mateusprt/lotus/token"
)

type Scanner struct {
	Source  string
	Tokens  []token.Token
	start   int
	current int
	line    int
}

var keywords = map[string]token.TokenType{
	"var":      token.VAR,
	"const":    token.CONST,
	"if":       token.IF,
	"else":     token.ELSE,
	"for":      token.FOR,
	"break":    token.BREAK,
	"continue": token.CONTINUE,
	"print":    token.PRINT,
	"function": token.FUNCTION,
	"return":   token.RETURN,
	"struct":   token.STRUCT,
	"len":      token.LEN,
	"first":    token.FIRST,
	"last":     token.LAST,
	"push":     token.PUSH,
	"true":     token.TRUE,
	"false":    token.FALSE,
	"null":     token.NULL,
}

func New(source []byte) *Scanner {
	return &Scanner{
		Source:  string(source),
		start:   0,
		current: 0,
		line:    1,
	}
}

func ScanTokens(s *Scanner) []token.Token {
	for !isEOF(s) {
		s.start = s.current
		scanToken(s)
	}
	s.Tokens = append(s.Tokens, token.Token{
		Type: token.EOF,
		Line: s.line,
	})
	return s.Tokens
}

func scanToken(s *Scanner) {
	charScanned := getNextChar(s)
	fmt.Printf("char escaneado: %s\n", charScanned)
	switch charScanned {
	case "(":
		addToken(s, token.LPAREN)
	case ")":
		addToken(s, token.RPAREN)
	case "{":
		addToken(s, token.LBRACE)
	case "}":
		addToken(s, token.RBRACE)
	case "[":
		addToken(s, token.LBRACKET)
	case "]":
		addToken(s, token.RBRACKET)
	case ",":
		addToken(s, token.COMMA)
	case ";":
		addToken(s, token.SEMICOLON)
	case ":":
		addToken(s, token.COLON)
	case ".":
		addToken(s, token.DOT)
	case "=":
		if advanceIfMatch(s, "=") {
			addToken(s, token.EQ)
		} else {
			addToken(s, token.ASSIGN)
		}
	case "+":
		if advanceIfMatch(s, "+") {
			addToken(s, token.INCREMENT)
		} else {
			addToken(s, token.PLUS)
		}
	case "-":
		if advanceIfMatch(s, "-") {
			addToken(s, token.DECREMENT)
		} else {
			addToken(s, token.MINUS)
		}
	case "*":
		addToken(s, token.MULTIPLY)
	case "/":
		if advanceIfMatch(s, "/") {
			for peek(s) != "\n" && !isEOF(s) {
				s.current++
			}
		} else {
			addToken(s, token.DIVIDE)
		}
	case "!":
		if advanceIfMatch(s, "=") {
			addToken(s, token.NOT_EQ)
		} else {
			addToken(s, token.BANG)
		}
	case "<":
		if advanceIfMatch(s, "=") {
			addToken(s, token.LE)
		} else {
			addToken(s, token.LT)
		}
	case ">":
		if advanceIfMatch(s, "=") {
			addToken(s, token.GE)
		} else {
			addToken(s, token.GT)
		}
	case " ":
	case "\r":
	case "\t":
	case "\n":
		s.line++
	case "\"":
		getString(s)
	default:
		if isDigit(charScanned) {
			getNumber(s)
			return
		}
		if isAlpha(charScanned) {
			getIdentifier(s)
			return
		}
		Error(s.line, fmt.Sprintf("SyntaxError: Invalid or unexpected symbol: %s", charScanned))
	}
}

func getIdentifier(s *Scanner) {
	for isAlphaNumeric(peek(s)) {
		s.current++
	}
	text := s.Source[s.start:s.current]
	tokenType := keywords[text]
	if tokenType == "" {
		tokenType = token.IDENTIFIER
	}
	addToken(s, tokenType)
}

func isAlpha(ch string) bool {
	return (ch >= "a" && ch <= "z") || (ch >= "A" && ch <= "Z") || ch == "_"
}

func isAlphaNumeric(ch string) bool {
	return isAlpha(ch) || isDigit(ch)
}

func getNumber(s *Scanner) {
	for isDigit(peek(s)) {
		s.current++
	}
	if peek(s) == "." && isDigit(peekNext(s)) {
		s.current++
		for isDigit(peek(s)) {
			s.current++
		}
	}
	addToken(s, token.NUMBER, s.Source[s.start:s.current])
}

func peekNext(s *Scanner) string {
	if s.current+1 >= len(s.Source) {
		return ""
	}
	return string(s.Source[s.current+1])
}

func isDigit(ch string) bool {
	return ch >= "0" && ch <= "9"
}

func getString(s *Scanner) {
	for peek(s) != "\"" && !isEOF(s) {
		if peek(s) == "\n" {
			s.line++
		}
		s.current++
	}
	if isEOF(s) {
		Error(s.line, "Unterminated string.")
		return
	}
	s.current++
	literal := s.Source[s.start+1 : s.current-1]
	addToken(s, token.STRING, literal)
}

func peek(s *Scanner) string {
	if isEOF(s) {
		return ""
	}
	return string(s.Source[s.current])
}

func getNextChar(s *Scanner) string {
	ch := s.Source[s.current]
	s.current++
	return string(ch)
}

func addToken(s *Scanner, tokenType token.TokenType, literal ...interface{}) {
	text := s.Source[s.start:s.current]
	var lit interface{}
	if len(literal) > 0 {
		lit = literal[0]
	} else {
		lit = nil
	}
	s.Tokens = append(s.Tokens, token.Token{
		Type:    tokenType,
		Lexeme:  text,
		Literal: lit,
		Line:    s.line,
	})
}

func Error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Printf("[line %d] Error%s: %s\n", line, where, message)
}

func isEOF(s *Scanner) bool {
	return s.current >= len(s.Source)
}

func advanceIfMatch(s *Scanner, expected string) bool {
	if isEOF(s) {
		return false
	}
	if string(s.Source[s.current]) != expected {
		return false
	}
	s.current++
	return true
}
