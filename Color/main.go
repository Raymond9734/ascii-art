package main

import (
	"fmt"
	"os"
	"strings"

	Ascii "asci-art/Banner"
	colour "asci-art/Color/Functions"
)

func main() {
	color := ""
	str := ""
	fileName := ""
	// Check for the correct number of arguments.
	input := os.Args[1:]
	if len(input) < 2 || len(input) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING]	")
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
		fmt.Println("Usage: go run . [OPTION] [STRING]	")
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
			PrintAsciiColor(word, os.Args[2], fileName, color)

		}
	}

}

func ToColor(str1 string, str2 string) (bool, int, int) {
	str := strings.Split(str1, " ")

	//return strings.Contains(str1, str2)
	startIndex := strings.Index(str1, str2)

	endIndex := startIndex + len(str2)
	for i := range str {
		if str[i] == str2 {
			return true, startIndex, endIndex
		}
	}
	if strings.Contains(str1, str2) {
		return true, 0, 0
	}

	return false, 0, 0

}

func PrintAsciiColor(word string, toColor, fileName, color string) {

	oK, startInd, endIdex := ToColor(word, toColor)
	r, g, b, err := colour.RgbExtract(color)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < 8; i++ {

		for j, letter := range word {
			line := Ascii.GetLine(1+int(letter-' ')*9+i, fileName)

			if oK && startInd > 0 {
				if j >= startInd && j <= endIdex {
					fmt.Print(colour.ESCseq(r, g, b), line)
				} else {
					fmt.Print(colour.ESCseq(255, 255, 255), line)
				}
			} else if oK && strings.ContainsRune(toColor, letter) {
				fmt.Print(colour.ESCseq(r, g, b), line)
			} else {
				fmt.Print(colour.ESCseq(255, 255, 255), line)
			}

		}
		fmt.Printf("\n")
	}

}
