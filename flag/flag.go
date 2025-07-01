package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define the --file flag
	filePath := flag.String("file", "", "Path to the file to be printed")

	// Parse the command-line flags
	flag.Parse()

	// Check if the file path was provided
	if *filePath == "" {
		fmt.Println("Please provide a file path using the --file flag.")
		flag.Usage() // Print usage instructions
		os.Exit(1)
	}

	// Read the content of the file
	content, err := os.ReadFile(*filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", *filePath, err)
		os.Exit(1)
	}

	// Print the content of the file
	fmt.Printf("Content of %s:\n%s\n", *filePath, content)
}
