package functions

import (
	"github.com/mateusprt/lotus/interpreter"
)

type PopFunction struct{}

func (c *PopFunction) Arity() int {
	return 1
}

func (c *PopFunction) Call(i *interpreter.Interpreter, arguments []interface{}) interface{} {
	if len(arguments) != 1 {
		panic("pop() expects 1 argument")
	}
	arr, ok := arguments[0].([]interface{})
	if !ok {
		panic("pop() expects an array as argument")
	}

	if len(arr) == 0 {
		panic("pop() cannot be called on an empty array")
	}

	newInstance := interpreter.NewStructInstance(interpreter.NewLotusStruct("Array", []string{"arr", "value"}))
	newInstance.SetField("array", arr[:len(arr)-1])
	newInstance.SetField("value", arr[len(arr)-1])
	return newInstance
}

func (c *PopFunction) String() string {
	return "<native fn>"
}
