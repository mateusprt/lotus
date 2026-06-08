package interpreter

import (
	"fmt"

	"github.com/mateusprt/lotus/errors"
	"github.com/mateusprt/lotus/token"
)

type LotusStruct struct {
	Name   string
	Fields []string
}

type LotusInstance struct {
	str    *LotusStruct
	fields map[string]interface{}
}

func (s *LotusInstance) String() string {
	result := "{"
	for _, field := range s.str.Fields {
		val := s.fields[field]
		result += " " + field + ": " + fmt.Sprintf("%v", val)
	}
	result += " }"
	return result
}

func (s *LotusStruct) Call(interpreter *Interpreter, arguments []interface{}) interface{} {
	return NewStructInstance(s)
}

func (s *LotusStruct) Arity() int {
	return 0
}

func (structInstance *LotusInstance) Get(name token.Token) interface{} {
	if value, ok := structInstance.fields[name.Lexeme]; ok {
		return value
	}
	errors.ThrowRuntimeError(name, "Undefined property '"+name.Lexeme+"'.")
	return nil
}

func (s *LotusInstance) Set(name token.Token, value interface{}) {
	s.fields[name.Lexeme] = value
}

func (s *LotusInstance) SetField(name string, value interface{}) {
	s.fields[name] = value
}

func NewStructInstance(str *LotusStruct) *LotusInstance {
	fields := make(map[string]interface{})
	for _, field := range str.Fields {
		fields[field] = nil
	}
	return &LotusInstance{str: str, fields: fields}
}

func NewLotusStruct(name string, fields []string) *LotusStruct { // Corrigido: recebe []string
	return &LotusStruct{Name: name, Fields: fields}
}
