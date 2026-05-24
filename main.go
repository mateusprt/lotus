package main

import (
	"fmt"
	"os"

	"github.com/mateusprt/lotus/cmd"
)

const version = "1.0"

func main() {

	if os.Args[1] == "--version" || os.Args[1] == "-v" {
		fmt.Printf("Lotus %s\n", version)
		return
	}

	if len(os.Args[1:]) > 1 {
		fmt.Println("Usage: lotus [script]")
		os.Exit(64)
	}

	if len(os.Args[1:]) == 0 {
		cmd.RunPrompt()
		return
	}

	cmd.RunFile(os.Args[1])
}
