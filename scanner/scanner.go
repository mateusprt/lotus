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
	default:
		Error(s.line, fmt.Sprintf("SyntaxError: Invalid or unexpected symbol: %s", charScanned))
	}
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

func addToken(s *Scanner, tokenType token.TokenType) {
	text := s.Source[s.start:s.current]
	s.Tokens = append(s.Tokens, token.Token{
		Type:    tokenType,
		Literal: text,
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
