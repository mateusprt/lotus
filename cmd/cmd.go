package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/mateusprt/lotus/ast/debug"
	"github.com/mateusprt/lotus/interpreter"
	"github.com/mateusprt/lotus/parser"
	"github.com/mateusprt/lotus/scanner"
)

func RunPrompt() {
	fmt.Println("Welcome to the Lotus REPL!")
	fmt.Println("Press Ctrl+D to exit.")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("\nGoodbye!")
				break
			}
			os.Exit(65)
		}
		run(line)
	}
}

func RunFile(path string) {
	byteSequence, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(65)
	}
	run(byteSequence)
}

func run(byteSequence []byte) {
	sc := scanner.New(byteSequence)
	interp := interpreter.New()
	tokens := scanner.ScanTokens(sc)

	p := parser.New(tokens)
	ast := parser.Parse(p)
	if ast == nil {
		os.Exit(65)
	}

	interp.Interpret(ast)

	print := debug.NewAstPrinter()
	print.Print(ast)
	fmt.Printf("Tokens: %v\n", tokens)
}
