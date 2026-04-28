package functions

import (
	"github.com/mateusprt/lotus/interpreter"
)

type LenFunction struct{}

func (c *LenFunction) Arity() int {
	return 1
}

func (c *LenFunction) Call(interpreter *interpreter.Interpreter, arguments []interface{}) interface{} {
	str, strOk := arguments[0].(string)
	arr, arrOk := arguments[0].([]interface{})
	if !strOk && !arrOk {
		panic("len() expects a string or an array as argument")
	}
	if strOk {
		return float64(len(str))
	}
	return float64(len(arr))
}
func (c *LenFunction) String() string {
	return "<native fn>"
}
