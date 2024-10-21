package main

import (
	"fmt"
	"os"

	"car-pooling/work"
)

func main() {
	// Check if a file path was provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide an input file as a command-line argument.")
		return
	}

	// Open the file for reading, using the file path provided as a command-line argument
	filePath := os.Args[1]
	results, err := work.ProcessFile(filePath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Output the results for all test cases
	for i, result := range results {
		fmt.Printf("Case #%d: %s\n", i+1, result)
	}
}
