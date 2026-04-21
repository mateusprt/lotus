package functions

import (
	"github.com/mateusprt/lotus/interpreter"
)

type LastFunction struct{}

func (c *LastFunction) Arity() int {
	return 1
}

func (c *LastFunction) Call(interpreter *interpreter.Interpreter, arguments []interface{}) interface{} {
	if len(arguments) != 1 {
		panic("last() expects 1 argument")
	}
	str, ok := arguments[0].([]interface{})
	if !ok {
		panic("last() expects an array as argument")
	}
	if len(str) == 0 {
		return nil
	}
	return str[len(str)-1]
}

func (c *LastFunction) String() string {
	return "<native fn>"
}
