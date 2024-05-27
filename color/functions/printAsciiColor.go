package colour

import (
	Ascii "asci-art/banner"
	"fmt"
	"os"
	"strings"
)

// ToColor searches for a specific word within a string and returns its start and end indices if found.
// It also checks if the word contains any characters from another string.
func ToColor(str1 string, str2 string) (bool, int, int) {
	str := strings.Split(str1, " ")
	var startIndex int
	var endIndex int

	for i := range str {
		// Find the index of the second string in the first string.
		startIndex = strings.Index(str1, str2)
		endIndex = startIndex + len(str2) - 1
		if str[i] == str2 {
			return true, startIndex, endIndex
		}
	}
	// If no exact match was found, reset start and end indices.
	startIndex = 0
	endIndex = 0
	// If the second string's length is greater than 1, attempt to find it as a substring.
	if len(str2) > 1 {
		startIndex, endIndex = FindSubstringIndx(str1, str2)
		if startIndex >= 0 && endIndex > 0 {
			return true, startIndex, endIndex

		}
	}
	for _, ch := range str1 {
		if strings.ContainsRune(str2, ch) {
			return true, 0, 0
		}

	}
	fmt.Printf("%s was not found in the text\n", str2)

	return false, 0, 0

}

// FindSubstringIndx finds the start and end indices of a substring within a string.
func FindSubstringIndx(str, subStr string) (int, int) {

	for i := 0; i <= len(str)-len(subStr); i++ {
		if str[i:i+len(subStr)] == subStr {
			if len(subStr) > 1 {
				return i, i + len(subStr) - 1
			} else {
				return i, i + len(subStr)
			}

		}

	}
	return -1, -1

}

// PrintAsciiColor prints a word in ASCII art with optional coloring for a specific substring.
func PrintAsciiColor(word string, toColor map[rune]string, wordToColor, fileName string) {
	// Determines if the word needs to be colored, and finds the start and end indices for coloring.
	oK, startInd, endIdex := ToColor(word, wordToColor)
	var r, g, b int
	var err error

	for i := 0; i < 8; i++ {

		for j, letter := range word {
			// Get the ASCII art line for the current letter.
			line := Ascii.GetLine(1+int(letter-' ')*9+i, fileName)
			// Check if the word needs to be colored and if the current character falls within the start and end indices.
			if oK && startInd >= 0 && endIdex > 0 {
				if color, found := toColor[letter]; found {
					r, g, b, err = RgbExtract(color)
					if err != nil {

						fmt.Println(err)
						os.Exit(1)
					}
				}
				if j >= startInd && j <= endIdex {

					fmt.Print(ESCseq(r, g, b), line)
				} else {
					fmt.Print(ESCseq(255, 255, 255), line)
				}
			} else if oK && strings.ContainsRune(wordToColor, letter) {
				if color, found := toColor[letter]; found {
					r, g, b, err = RgbExtract(color)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				}
				fmt.Print(ESCseq(r, g, b), line)
			} else {
				// Print the character in white if no specific color is assigned.
				fmt.Print(ESCseq(255, 255, 255), line)
			}

		}
		fmt.Println()
	}

}
