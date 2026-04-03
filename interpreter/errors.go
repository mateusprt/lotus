package interpreter

import (
	"fmt"

	"github.com/mateusprt/lotus/token"
)

type RuntimeError struct {
	Token   token.Token
	Message string
}

func ThrowRuntimeError(tok token.Token, message string) {
	panic(&RuntimeError{Token: tok, Message: message})
}

func PrintRuntimeError(err *RuntimeError) {
	fmt.Printf("[line %d] RuntimeError: %s\n", err.Token.Line, err.Message)
}
