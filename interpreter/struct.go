package interpreter

import (
	"github.com/mateusprt/lotus/errors"
	"github.com/mateusprt/lotus/token"
)

type StructName struct {
	Name string
}

func NewStructName(name string) *StructName {
	return &StructName{Name: name}
}

func (s *StructName) String() string {
	return s.Name
}

func (s *StructName) Call(interpreter *Interpreter, arguments []interface{}) interface{} {
	return NewStructInstance(s)
}

func (s *StructName) Arity() int {
	return 0
}

type StructInstance struct {
	str    *StructName
	fields map[string]interface{}
}

func NewStructInstance(str *StructName) *StructInstance {
	return &StructInstance{str: str, fields: make(map[string]interface{})}
}

func (s *StructInstance) String() string {
	return s.str.Name + " instance"
}

func (s *StructInstance) Get(name token.Token) interface{} {
	if value, ok := s.fields[name.Lexeme]; ok {
		return value
	}
	errors.ThrowRuntimeError(name, "Undefined property '"+name.Lexeme+"'.")
	return nil
}
