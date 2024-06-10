package main

import (
	"fmt"
	"os"

	"github.com/michalslomczynski/gyara"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <yara_file>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	ruleContents := string(data)
	parsedRule := gyara.ParseRule(ruleContents)

	fmt.Println(parsedRule)
}
