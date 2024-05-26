package colour

import (
	Ascii "asci-art/Banner"
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

		startIndex = strings.Index(str1, str2)
		endIndex = startIndex + len(str2) - 1
		if str[i] == str2 {
			return true, startIndex, endIndex
		}
	}
	startIndex = 0
	endIndex = 0
	if len(str2) > 1 {
		startIndex, endIndex = FindSubstringIndx(str1, str2)
		if startIndex >= 0 && endIndex >= 0 {
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
func PrintAsciiColor(word string, toColor map[rune]string, wordToColor, fileName, color string) {

	oK, startInd, endIdex := ToColor(word, wordToColor)
	var r, g, b int
	var err error
	// r, g, b, err := RgbExtract(color)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	for i := 0; i < 8; i++ {

		for j, letter := range word {
			// Get the ASCII art line for the current letter.
			line := Ascii.GetLine(1+int(letter-' ')*9+i, fileName)

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

				fmt.Print(ESCseq(255, 255, 255), line)
			}

		}
		fmt.Printf("\n")
	}

}

// func PrintAsciiColor(word string, toColor map[rune]string, fileName string) {
// 	for i := 0; i < 8; i++ {
// 		for _, letter := range word {
// 			// Get the ASCII art line for the current letter.
// 			line := GetLine(1+int(letter-' ')*9+i, fileName)

// 			if color, found := toColor[letter]; found {
// 				r, g, b, err := RgbExtract(color)
// 				if err != nil {
// 					fmt.Println(err)
// 					os.Exit(1)
// 				}
// 				fmt.Print(ESCseq(r, g, b), line)
// 			} else {
// 				// Default color (white)
// 				fmt.Print(ESCseq(255, 255, 255), line)
// 			}
// 		}
// 		fmt.Printf("\n")
// 	}
// }
