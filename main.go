package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		input    = flag.String("input", "", "Input string to generate identicon from (required)")
		output   = flag.String("output", "", "Output PNG file path (optional, defaults to input.png)")
		help     = flag.Bool("help", false, "Show help message")
	)
	
	flag.Parse()
	
	if *help || *input == "" {
		printUsage()
		if *input == "" {
			os.Exit(1)
		}
		return
	}
	
	// Set default output filename if not provided
	outputFile := *output
	if outputFile == "" {
		outputFile = *input + ".png"
	}
	
	fmt.Printf("Generating identicon for: %s\n", *input)
	
	// Generate identicon
	identicon := New(*input)
	
	// Save to file
	err := identicon.Save(outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error saving identicon: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("Identicon saved to: %s\n", outputFile)
}

func printUsage() {
	fmt.Println("Identicon Generator")
	fmt.Println("Generates unique identicon images from input strings")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Printf("  %s -input=<string> [-output=<filename>]\n", os.Args[0])
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -input    Input string to generate identicon from (required)")
	fmt.Println("  -output   Output PNG file path (optional, defaults to input.png)")
	fmt.Println("  -help     Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Printf("  %s -input=john.doe@example.com\n", os.Args[0])
	fmt.Printf("  %s -input=myusername -output=avatar.png\n", os.Args[0])
}