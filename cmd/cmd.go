package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
			fmt.Printf("Error reading input: %s\n", err)
			break
		}
		run(line)
	}
}

func RunFile(path string) {
	byteSequence, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}
	run(byteSequence)
}

func run(byteSequence []byte) {
	fmt.Printf("Running code:\n%s\n", string(byteSequence))
}
