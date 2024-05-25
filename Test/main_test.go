package main

import (
	colour "asci-art/Color/Functions"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	Ascii "asci-art/Banner"
)

// TestLoadBanner tests the LoadBanner function to ensure it correctly loads banner characters from a file.
func TestLoadBanner(t *testing.T) {
	testCases := []struct {
		fiLename      string
		expectedChars []rune
	}{
		{"standard", []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '+', '`', '~', '?', ',', '\'', '"', ';', '=', '_', '/', '\\', '.', '<', '>', '|', '[', ']', '{', '}', ':'}},   // Example test case with multiple characters
		{"thinkertoy", []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '+', '`', '~', '?', ',', '\'', '"', ';', '=', '_', '/', '\\', '.', '<', '>', '|', '[', ']', '{', '}', ':'}}, // New test case
		{"shadow", []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '+', '`', '~', '?', ',', '\'', '"', ';', '=', '_', '/', '\\', '.', '<', '>', '|', '[', ']', '{', '}', ':'}},
	}

	for _, tc := range testCases {
		banner := Ascii.LoadBanner(tc.fiLename)

		// Check if the loaded banner is not empty
		if len(banner) == 0 {
			t.Errorf("For file %q: Expected banner to load characters, but got none", tc.fiLename)
		}

		// Check if all expected characters are in the loaded banner
		for _, char := range tc.expectedChars {
			if _, ok := banner[char]; !ok {
				t.Errorf("For file %q: Expected to find character '%c' in banner, but it was not found", tc.fiLename, char)
			}
		}
	}
}

//TestPrintBanner tests the PrintBanner function to ensure it correctly prints the input string using the loaded banner characters.
func TestPrintBanner(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"A", "           \n    /\\     \n   /  \\    \n  / /\\ \\   \n / ____ \\  \n/_/    \\_\\ \n           \n           \n"},
		{"!", "_  \n| | \n| | \n| | \n|_| \n(_) \n    \n    \n"},                   // New test case
		{"\"", " _ _  \n( | ) \n V V  \n      \n      \n      \n      \n      \n"}, // New test case
		{"$", "  _   \n | |  \n/ __) \n\\__ \\ \n(   / \n |_|  \n      \n      \n"},
		{"%", " _   __ \n(_) / / \n   / /  \n  / /   \n / / _  \n/_/ (_) \n        \n        \n"},
		{"&", "         \n  ___    \n ( _ )   \n / _ \\/\\ \n| (_>  < \n \\___/\\/ \n         \n         \n"},
		{"(", "  __ \n / / \n| |  \n| |  \n| |  \n| |  \n \\_\\ \n     \n"},
		{")", "__   \n\\ \\  \n | | \n | | \n | | \n | | \n/_/  \n     \n"},
		{"*", "    _     \n /\\| |/\\  \n \\ ` ' /  \n|_     _| \n / , . \\  \n \\/|_|\\/  \n          \n          \n"},
		{"+", "        \n   _    \n _| |_  \n|_   _| \n  |_|   \n        \n        \n        \n"},
		{"'", " _  \n( ) \n|/  \n    \n    \n    \n    \n    \n"},
		{"-", "         \n         \n ______  \n|______| \n         \n         \n         \n         \n"},
		{".", "    \n    \n    \n    \n _  \n(_) \n    \n    \n"},
		{"/", "     __ \n    / / \n   / /  \n  / /   \n / /    \n/_/     \n        \n        \n"},
		{"0", "        \n  ___   \n / _ \\  \n| | | | \n| |_| | \n \\___/  \n        \n        \n"},
		{"<", "   __ \n  / / \n / /  \n< <   \n \\ \\  \n  \\_\\ \n      \n      \n"},
		{"=", "         \n ______  \n|______| \n ______  \n|______| \n         \n         \n         \n"},
		{">", "__    \n\\ \\   \n \\ \\  \n  > > \n / /  \n/_/   \n      \n      \n"},
		{"?", " ___   \n|__ \\  \n   ) | \n  / /  \n |_|   \n (_)   \n       \n       \n"},
		{"@", "          \n   ____   \n  / __ \\  \n / / _` | \n| | (_| | \n \\ \\__,_| \n  \\____/  \n          \n"},
		{"A", "           \n    /\\     \n   /  \\    \n  / /\\ \\   \n / ____ \\  \n/_/    \\_\\ \n           \n           \n"},
		{"B", " ____   \n|  _ \\  \n| |_) | \n|  _ <  \n| |_) | \n|____/  \n        \n        \n"},
		{"C", "  _____  \n / ____| \n| |      \n| |      \n| |____  \n \\_____| \n         \n         \n"},
		{"D", " _____   \n|  __ \\  \n| |  | | \n| |  | | \n| |__| | \n|_____/  \n         \n         \n"},
		{"E", " ______  \n|  ____| \n| |__    \n|  __|   \n| |____  \n|______| \n         \n         \n"},
		{"F", " ______  \n|  ____| \n| |__    \n|  __|   \n| |      \n|_|      \n         \n         \n"},
		{"G", "  _____  \n / ____| \n| |  __  \n| | |_ | \n| |__| | \n \\_____| \n         \n         \n"},
		{"H", " _    _  \n| |  | | \n| |__| | \n|  __  | \n| |  | | \n|_|  |_| \n         \n         \n"},
		{"I", " _____  \n|_   _| \n  | |   \n  | |   \n _| |_  \n|_____| \n        \n        \n"},
		{"J", "      _  \n     | | \n     | | \n _   | | \n| |__| | \n \\____/  \n         \n         \n"},
		{"K", " _  __ \n| |/ / \n| ' /  \n|  <   \n| . \\  \n|_|\\_\\ \n       \n       \n"},
		{"L", " _       \n| |      \n| |      \n| |      \n| |____  \n|______| \n         \n         \n"},
		{"M", " __  __  \n|  \\/  | \n| \\  / | \n| |\\/| | \n| |  | | \n|_|  |_| \n         \n         \n"},
		{"N", " _   _  \n| \\ | | \n|  \\| | \n| . ` | \n| |\\  | \n|_| \\_| \n        \n        \n"},
		{"O", "  ____   \n / __ \\  \n| |  | | \n| |  | | \n| |__| | \n \\____/  \n         \n         \n"},
		{"P", " _____   \n|  __ \\  \n| |__) | \n|  ___/  \n| |      \n|_|      \n         \n         \n"},
		{"Q", "  ____   \n / __ \\  \n| |  | | \n| |  | | \n| |__| | \n \\___\\_\\ \n         \n         \n"},
		{"R", " _____   \n|  __ \\  \n| |__) | \n|  _  /  \n| | \\ \\  \n|_|  \\_\\ \n         \n         \n"},
		{"S", "  _____  \n / ____| \n| (___   \n \\___ \\  \n ____) | \n|_____/  \n         \n         \n"},
		{"T", " _______  \n|__   __| \n   | |    \n   | |    \n   | |    \n   |_|    \n          \n          \n"},
		{"U", " _    _  \n| |  | | \n| |  | | \n| |  | | \n| |__| | \n \\____/  \n         \n         \n"},
		{"V", "__      __ \n\\ \\    / / \n \\ \\  / /  \n  \\ \\/ /   \n   \\  /    \n    \\/     \n           \n           \n"},
		{"W", "__          __ \n\\ \\        / / \n \\ \\  /\\  / /  \n  \\ \\/  \\/ /   \n   \\  /\\  /    \n    \\/  \\/     \n               \n               \n"},
		{"X", "__   __ \n\\ \\ / / \n \\ V /  \n  > <   \n / . \\  \n/_/ \\_\\ \n        \n        \n"},
		{"Y", "__     __ \n\\ \\   / / \n \\ \\_/ /  \n  \\   /   \n   | |    \n   |_|    \n          \n          \n"},
		{"Z", " ______ \n|___  / \n   / /  \n  / /   \n / /__  \n/_____| \n        \n        \n"},
		{"[", " ___  \n|  _| \n| |   \n| |   \n| |   \n| |_  \n|___| \n      \n"},
		{"\\", "__      \n\\ \\     \n \\ \\    \n  \\ \\   \n   \\ \\  \n    \\_\\ \n        \n        \n"},
		{"]", " ___  \n|_  | \n  | | \n  | | \n  | | \n _| | \n|___| \n      \n"},
		{"^", " /\\  \n|/\\| \n     \n     \n     \n     \n     \n     \n"},
		{"_", "         \n         \n         \n         \n         \n         \n ______  \n|______| \n"},
		{"`", " _  \n( ) \n \\| \n    \n    \n    \n    \n    \n"},
		{"a", "        \n        \n  __ _  \n / _` | \n| (_| | \n \\__,_| \n        \n        \n"},
		{"b", " _      \n| |     \n| |__   \n| '_ \\  \n| |_) | \n|_.__/  \n        \n        \n"},
		{"c", "       \n       \n  ___  \n / __| \n| (__  \n \\___| \n       \n       \n"},
		{"d", "     _  \n    | | \n  __| | \n / _` | \n| (_| | \n \\__,_| \n        \n        \n"},
		{"e", "       \n       \n  ___  \n / _ \\ \n|  __/ \n \\___| \n       \n       \n"},
		{"f", "  __  \n / _| \n| |_  \n|  _| \n| |   \n|_|   \n      \n      \n"},
		{"g", "        \n        \n  __ _  \n / _` | \n| (_| | \n \\__, | \n  __/ | \n |___/  \n"},
		{"h", " _      \n| |     \n| |__   \n|  _ \\  \n| | | | \n|_| |_| \n        \n        \n"},
		{"i", " _  \n(_) \n _  \n| | \n| | \n|_| \n    \n    \n"},
		{"j", "   _  \n  (_) \n   _  \n  | | \n  | | \n  | | \n _/ | \n|__/  \n"},
		{"k", "       \n _     \n| | _  \n| |/ / \n|   <  \n|_|\\_\\ \n       \n       \n"},
		{"l", " _  \n| | \n| | \n| | \n| | \n|_| \n    \n    \n"},
		{"m", "            \n            \n _ __ ___   \n| '_ ` _ \\  \n| | | | | | \n|_| |_| |_| \n            \n            \n"},
		{"n", "        \n        \n _ __   \n| '_ \\  \n| | | | \n|_| |_| \n        \n        \n"},
		{"o", "        \n        \n  ___   \n / _ \\  \n| (_) | \n \\___/  \n        \n        \n"},
		{"p", "        \n        \n _ __   \n| '_ \\  \n| |_) | \n| .__/  \n| |     \n|_|     \n"},
		{"q", "        \n        \n  __ _  \n / _` | \n| (_| | \n \\__, | \n    | | \n    |_| \n"},
		{"r", "       \n       \n _ __  \n| '__| \n| |    \n|_|    \n       \n       \n"},
		{"s", "      \n      \n ___  \n/ __| \n\\__ \\ \n|___/ \n      \n      \n"},
		{"t", " _    \n| |   \n| |_  \n| __| \n\\ |_  \n \\__| \n      \n      \n"},
		{"u", "        \n        \n _   _  \n| | | | \n| |_| | \n \\__,_| \n        \n        \n"},
		{"v", "        \n        \n__   __ \n\\ \\ / / \n \\ V /  \n  \\_/   \n        \n        \n"},
		{"w", "           \n           \n__      __ \n\\ \\ /\\ / / \n \\ V  V /  \n  \\_/\\_/   \n           \n           \n"},
		{"x", "       \n       \n__  __ \n\\ \\/ / \n >  <  \n/_/\\_\\ \n       \n       \n"},
		{"y", "        \n        \n _   _  \n| | | | \n| |_| | \n \\__, | \n __/ /  \n|___/   \n"},
		{"z", "      \n      \n ____ \n|_  / \n / /  \n/___| \n      \n      \n"},
		{"{", "   __ \n  / / \n | |  \n/ /   \n\\ \\   \n | |  \n  \\_\\ \n      \n"},
		{"|", " _  \n| | \n| | \n| | \n| | \n| | \n| | \n|_| \n"},
		{"}", "__    \n\\ \\   \n | |  \n  \\ \\ \n  / / \n | |  \n/_/   \n      \n"},
		{"~", " /\\/| \n|/\\/  \n      \n      \n      \n      \n      \n      \n"},
	}

	for _, tc := range testCases {
		// Redirecting standard output to capture the PrintBanner output
		old := os.Stdout // keep backup of the real stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		Ascii.PrintBanner(tc.input) // Pass the input from the test case

		w.Close()
		out, _ := io.ReadAll(r)
		os.Stdout = old // restoring the real stdout

		if !strings.Contains(string(out), tc.expectedOutput) {
			t.Errorf("Input %q: Expected output to contain %q, got %q", tc.input, tc.expectedOutput, string(out))
		}
	}
}

// // // TestGetLine tests the GetLine function to ensure it correctly reads a specified line from a file.
// func TestGetLine(t *testing.T) {
// 	// Test cases for the GetLine function
// 	tests := []struct {
// 		num      int    // Line number to retrieve
// 		filename string // File name
// 		expected string // Expected line content
// 	}{
// 		{23, "standard", "This is line 1"},
// 		{40, "standard", " \\__ \\ "},
// 		{48, "standard", "    / /  "},
// 		{99, "standard", "This is line 15"},
// 		{120, "standard", "  ______  "},
// 	}

// 	for _, test := range tests {
// 		t.Run(fmt.Sprintf("Line %d", test.num), func(t *testing.T) {
// 			line := Ascii.GetLine(test.num, test.filename)
// 			if line != test.expected {
// 				t.Errorf("Got: %s, Expected: %s", line, test.expected)
// 			}
// 		})
// 	}
// }package main

// Define a custom struct for test cases
type testCase struct {
	str1     string
	str2     string
	expected struct {
		found    bool
		startInd int
		endInd   int
	}
}

func TestToColor(t *testing.T) {
	// Test cases for the ToColor function
	tests := []testCase{
		{
			str1: "hello world",
			str2: "hello",
			expected: struct {
				found    bool
				startInd int
				endInd   int
			}{true, 0, 4},
		},
		{
			str1: "hello world",
			str2: "world",
			expected: struct {
				found    bool
				startInd int
				endInd   int
			}{true, 6, 10},
		},
		{
			str1: "hello world",
			str2: "foo",
			expected: struct {
				found    bool
				startInd int
				endInd   int
			}{true, 0, 0},
		},
		{
			str1: "hello world",
			str2: "l",
			expected: struct {
				found    bool
				startInd int
				endInd   int
			}{true, 0, 0},
		},
	}

	// Iterate over each test case
	for _, test := range tests {
		// Run each test case as a sub-test
		t.Run(fmt.Sprintf("ToColor(%q, %q)", test.str1, test.str2), func(t *testing.T) {
			// Call the function under test
			found, startInd, endInd := colour.ToColor(test.str1, test.str2)
			expected := test.expected

			// Compare the actual result with the expected result
			if found != expected.found || startInd != expected.startInd || endInd != expected.endInd {
				t.Errorf("Expected (%t, %d, %d), got (%t, %d, %d)", expected.found, expected.startInd, expected.endInd, found, startInd, endInd)
			}
		})
	}
}
