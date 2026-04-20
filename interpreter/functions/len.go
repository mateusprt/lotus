package functions

import (
	"github.com/mateusprt/lotus/interpreter"
)

type LenFunction struct{}

func (c *LenFunction) Arity() int {
	return 1
}

func (c *LenFunction) Call(interpreter *interpreter.Interpreter, arguments []interface{}) interface{} {
	str, ok := arguments[0].(string)
	if !ok {
		panic("len() expects a string as argument")
	}
	return float64(len(str))
}

func (c *LenFunction) String() string {
	return "<native fn>"
}
