package main

import (
	"fmt"
	"os"
	"training-go/sorting_structures/adapters/cli"
)

func main() {
	cliHandler := cli.NewCLIHandler()
	if err := cliHandler.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
