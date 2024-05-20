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
			for i := 0; i < 8; i++ {
				for _, letter := range word {
					line := Ascii.GetLine(1+int(letter-' ')*9+i, fileName)
					oK, toClr := ToColor(word, os.Args[2])

					if len(input) > 2 {
						words := strings.Split(word, " ")
						if (len(input) == 4 || len(os.Args) == 4) && oK && Check(words, toClr) && strings.ContainsRune(toClr, letter) {
							fmt.Print(colour.ESCseq(color), line)
						} else {
							fmt.Print(colour.ESCseq("white"), line)
						}

					} else {
						fmt.Print(colour.ESCseq(color), line)
					}
				}
				fmt.Printf("\n")
			}
		}
	}

}

func ToColor(str1 string, str2 string) (bool, string) {
	str := strings.Split(str1, " ")

	//return strings.Contains(str1, str2)

	for i := range str {
		if str[i] == str2 {
			return true, str[i]
		}

	}
	return false, ""

}
func Check(s []string, str string) bool {

	for i := range s {
		if s[i] == str {
			return true
		}
	}
	return false
}
