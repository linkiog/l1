package main

import (
	"fmt"
	"strings"
)

func hasUniqueChars(input string) bool {

	seen := make(map[rune]bool)

	input = strings.ToLower(input)

	for _, char := range input {
		if seen[char] {
			return false
		}
		seen[char] = true
	}

	return true
}

func main() {
	tests := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
	}

	for _, test := range tests {
		fmt.Printf("Строка: %s, Уникальные символы: %t\n", test, hasUniqueChars(test))
	}
}

// если без map
// func hasUniqueChars(input string) bool {
// 	input = strings.ToLower(input)
// 	for i := 0; i < len(input); i++ {
// 		for j := i + 1; j < len(input); j++ {
// 			if input[i] == input[j] {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }
