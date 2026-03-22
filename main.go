package main

import (
	"fmt"
	"os"

	"github.com/mateusprt/lotus/cmd"
)

func main() {
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
