package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/mateusprt/lotus/environment"
	"github.com/mateusprt/lotus/environment/resolver"
	"github.com/mateusprt/lotus/errors"
	"github.com/mateusprt/lotus/interpreter"
	"github.com/mateusprt/lotus/interpreter/functions"
	"github.com/mateusprt/lotus/parser"
	"github.com/mateusprt/lotus/scanner"
	"github.com/peterh/liner"
)

func initEnv() *environment.Environment {
	env := environment.New()
	environment.Define(env, "now", &functions.NowFunction{})
	environment.Define(env, "len", &functions.LenFunction{})
	environment.Define(env, "first", &functions.FirstFunction{})
	environment.Define(env, "last", &functions.LastFunction{})
	environment.Define(env, "push", &functions.PushFunction{})
	environment.Define(env, "pop", &functions.PopFunction{})
	return env
}

func RunPrompt() {
	fmt.Println("Welcome to the Lotus REPL!")
	fmt.Println("Press Ctrl+D to exit.")

	lineReader := liner.NewLiner()
	defer lineReader.Close()
	lineReader.SetCtrlCAborts(true)

	interp := interpreter.New(initEnv())

	var buffer string
	openBraces := 0

	for {
		var prompt string
		if openBraces == 0 {
			prompt = "> "
		} else {
			prompt = ". "
		}
		line, err := lineReader.Prompt(prompt)

		if err != nil {
			if err == io.EOF {
				fmt.Println("\nGoodbye!")
			}
			if err == liner.ErrPromptAborted {
				if (openBraces % 2) != 0 {
					openBraces--
				}
				continue
			}
			break
		}
		lineReader.AppendHistory(line)
		buffer += line + "\n"

		// Conta chaves abertas e fechadas
		for _, b := range line {
			if b == '{' {
				openBraces++
			}
			if b == '}' {
				openBraces--
			}
		}

		// Só executa quando todos os blocos estão fechados
		if openBraces <= 0 && len(buffer) > 0 {
			if err := run([]byte(buffer), interp); err != nil {
				if runtimeErr, ok := err.(*errors.RuntimeError); ok {
					errors.PrintRuntimeError(runtimeErr)
				} else {
					fmt.Println(err)
				}
			}
			buffer = ""
			openBraces = 0
		}
	}
}

func RunFile(path string) {
	ext := filepath.Ext(path)
	if ext != ".lt" {
		fmt.Println("Error: only files with .lt extension are allowed.")
		os.Exit(65)
	}
	byteSequence, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(65)
	}
	if err := run(byteSequence, interpreter.New(initEnv())); err != nil {
		if runtimeErr, ok := err.(*errors.RuntimeError); ok {
			errors.PrintRuntimeError(runtimeErr)
		} else {
			fmt.Println(err)
		}
		os.Exit(65)
	}
}

func run(byteSequence []byte, interp *interpreter.Interpreter) error {
	sc := scanner.New(byteSequence)
	tokens := scanner.ScanTokens(sc)
	p := parser.New(tokens)
	statements := parser.Parse(p)

	res := resolver.New(interp)
	resolver.Resolve(res, statements)
	return interp.Interpret(statements)
}
