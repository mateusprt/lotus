package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/mateusprt/lotus/environment"
	"github.com/mateusprt/lotus/environment/resolver"
	"github.com/mateusprt/lotus/interpreter"
	"github.com/mateusprt/lotus/interpreter/functions"
	"github.com/mateusprt/lotus/parser"
	"github.com/mateusprt/lotus/scanner"
)

func RunPrompt() {
	fmt.Println("Welcome to the Lotus REPL!")
	fmt.Println("Press Ctrl+D to exit.")
	reader := bufio.NewReader(os.Stdin)
	env := environment.New()
	interp := interpreter.New(env)
	environment.Define(env, "now", &functions.NowFunction{})

	var buffer []byte
	openBraces := 0

	for {
		if openBraces == 0 {
			fmt.Print("> ")
		} else {
			fmt.Print(". ")
		}
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("\nGoodbye!")
				break
			}
			os.Exit(65)
		}
		buffer = append(buffer, line...)

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
			run(buffer, interp)
			buffer = nil
			openBraces = 0
		}
	}
}

func RunFile(path string) {
	byteSequence, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(65)
	}
	run(byteSequence, interpreter.New(environment.New()))
}

func run(byteSequence []byte, interp *interpreter.Interpreter) {
	sc := scanner.New(byteSequence)
	tokens := scanner.ScanTokens(sc)
	p := parser.New(tokens)
	statements := parser.Parse(p)

	res := resolver.New(interp)
	resolver.Resolve(res, statements)
	interp.Interpret(statements)
}
