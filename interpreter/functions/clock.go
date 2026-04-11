package functions

import (
	"time"

	"github.com/mateusprt/lotus/interpreter"
)

type NowFunction struct{}

func (c *NowFunction) Arity() int {
	return 0
}

func (c *NowFunction) Call(interpreter *interpreter.Interpreter, arguments []interface{}) interface{} {
	return float64(time.Now().UnixNano()) / 1e9
}

func (c *NowFunction) String() string {
	return "<native fn>"
}
