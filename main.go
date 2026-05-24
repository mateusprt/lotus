package main

import (
	"fmt"
	"os"

	"github.com/mateusprt/lotus/cmd"
)

const version = "1.0"

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Printf("Lotus %s\n", version)
		return
	}

	if len(os.Args) > 2 {
		fmt.Println("Usage: lotus [script]")
		os.Exit(64)
	}

	if len(os.Args) == 1 {
		cmd.RunPrompt()
		return
	}

	cmd.RunFile(os.Args[1])
}
