package main

import (
	"fmt"
	"os"
	"strings"

	Ascii "asci-art/Banner"
)

func main() {
	fileName := ""
	// Check for the correct number of arguments.
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Error: Invalid number of arguments")
		os.Exit(1)
	}
	input := os.Args[1]

	if len(os.Args) == 3 {
		fileName = os.Args[2]
	} else {
		fileName = "standard" // Default to "standard" if no file name is provided.
	}

	input = strings.Replace(input, "\\n", "\n", -1)
	//  Removing the non-printable characters in the input string.
	input = Ascii.HandleSpecialCase(input)

	if input == "\n" {
		fmt.Println()
		return
	} else if input == "" {
		return
	}
	// Split the input into lines based on newline characters.

	Input := strings.Split(input, "\n")

	spaceCount := 0
	// Iterate over each line of the input.
	for _, word := range Input {
		if word == "" {
			spaceCount++
			if spaceCount < len(Input) {
				fmt.Println()
			}
		} else {
			PrintBanner(word, fileName)

		}
	}
}
