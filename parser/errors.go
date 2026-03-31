package parser

import (
	"fmt"

	"github.com/mateusprt/lotus/token"
)

type ParseError struct {
	Type    string
	Message string
}

func (e *ParseError) Error() string {
	return e.Message
}

func NewParseError(currentToken token.Token, message string) *ParseError {
	errorMessage := fmt.Sprintf("%d at end: %s", currentToken.Line, message)
	if currentToken.Type == token.EOF {
		report(errorMessage)
		return &ParseError{Message: errorMessage, Type: "ParseError"}
	}
	errorMessage = fmt.Sprintf("%d at '%s': %s", currentToken.Line, currentToken.Lexeme, message)
	report(errorMessage)
	return &ParseError{Message: errorMessage, Type: "ParseError"}
}

func report(message string) {
	fmt.Println(message)
}
