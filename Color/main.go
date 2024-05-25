package main

import (
	colour "asci-art/Color/Functions"
	"fmt"
	"os"
	"strings"

	Ascii "asci-art/Banner"
)

func main() {
	color := ""
	str := ""
	fileName := ""
	// Check for the correct number of arguments.
	input := os.Args[1:]
	if len(input) < 2 || len(input) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING]  ")
		fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
	}

	if len(input) == 4 && !(input[len(input)-1] == "thinkertoy" || input[len(input)-1] == "shadow" || input[len(input)-1] == "standard") {
		fmt.Println("Usage: go run . [OPTION] [STRING]  ")
		fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
	}
	fileName = "standard"

	for _, v := range input {

		if len(v) >= 8 && v[0:8] == "--color=" {
			color = v[8:]
		} else if len(v) == 2 && v[0:2] == "--" {
			continue
		} else if v == "standard" || v == "thinkertoy" || v == "shadow" {
			fileName = v
		} else {
			if len(input) == 4 {
				str = input[2]
			} else if len(input) == 3 {
				str = input[2]
			} else if len(input) == 2 || (len(input) == 3 && input[len(input)-1] == "standard" || input[len(input)-1] == "thinkertoy" || input[len(input)-1] == "shadow") {
				str = input[1]
			}
		}
	}
	if len(color) == 0 {
		fmt.Println("Usage: go run . [OPTION] [STRING]  ")
		fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
	}

	str = strings.Replace(str, "\\n", "\n", -1)
	//  Removing the non-printable characters in the str string.
	str = Ascii.HandleSpecialCase(str)

	if str == "\n" {
		fmt.Println()
		return
	} else if str == "" {
		return
	}
	// Split the str into lines based on newline characters.

	Input := strings.Split(str, "\n")

	spaceCount := 0
	// Iterate over each line of the input.
	for _, word := range Input {
		if word == "" {
			spaceCount++
			if spaceCount < len(Input) {
				fmt.Println()
			}
		} else {
			colour.PrintAsciiColor(word, os.Args[2], fileName, color)

		}
	}

}
