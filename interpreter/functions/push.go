package functions

import (
	"github.com/mateusprt/lotus/interpreter"
)

type PushFunction struct{}

func (c *PushFunction) Arity() int {
	return 2
}

func (c *PushFunction) Call(interpreter *interpreter.Interpreter, arguments []interface{}) interface{} {
	if len(arguments) != 2 {
		panic("push() expects 2 arguments")
	}
	arr, ok := arguments[0].([]interface{})
	if !ok {
		panic("push() expects an array as argument")
	}

	arg := arguments[1]
	if arg == nil {
		panic("push() expects a value to push")
	}

	return append(arr, arg)
}

func (c *PushFunction) String() string {
	return "<native fn>"
}
