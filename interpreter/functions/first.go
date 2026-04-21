package functions

import (
	"github.com/mateusprt/lotus/interpreter"
)

type FirstFunction struct{}

func (c *FirstFunction) Arity() int {
	return 1
}

func (c *FirstFunction) Call(interpreter *interpreter.Interpreter, arguments []interface{}) interface{} {
	if len(arguments) != 1 {
		panic("first() expects 1 argument")
	}
	str, ok := arguments[0].([]interface{})
	if !ok {
		panic("first() expects an array as argument")
	}
	if len(str) == 0 {
		return nil
	}
	return str[0]
}

func (c *FirstFunction) String() string {
	return "<native fn>"
}
