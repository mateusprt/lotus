package environment

import (
	"github.com/mateusprt/lotus/errors"
	"github.com/mateusprt/lotus/token"
)

type Environment struct {
	values    map[string]interface{}
	enclosing *Environment
}

func New() *Environment {
	return &Environment{values: make(map[string]interface{})}
}

func Define(e *Environment, name string, value interface{}) {
	e.values[name] = value
}

func Get(e *Environment, name token.Token) interface{} {
	if e.values[name.Lexeme] != nil {
		return e.values[name.Lexeme]
	}

	if e.enclosing != nil {
		return Get(e.enclosing, name)
	}

	errors.ThrowRuntimeError(name, "Undefined variable '"+name.Lexeme+"'.")
	return nil
}

func GetAt(e *Environment, distance int, name string) interface{} {
	return ancestor(e, distance).values[name]
}

func ancestor(e *Environment, distance int) *Environment {
	env := e
	for i := 0; i < distance; i++ {
		env = env.enclosing
	}
	return env
}

func AssignAt(e *Environment, distance int, name token.Token, value interface{}) {
	ancestor(e, distance).values[name.Lexeme] = value
}

func Assign(e *Environment, name token.Token, value interface{}) {
	if _, ok := e.values[name.Lexeme]; ok {
		e.values[name.Lexeme] = value
		return
	}

	if e.enclosing != nil {
		Assign(e.enclosing, name, value)
		return
	}

	errors.ThrowRuntimeError(name, "Undefined variable '"+name.Lexeme+"'.")
}

func NewEnclosed(enclosing *Environment) *Environment {
	return &Environment{
		values:    make(map[string]interface{}),
		enclosing: enclosing,
	}
}
