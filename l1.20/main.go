package main

import "fmt"

func reverseWords(input string) string {
	var words []string
	word := ""
	for _, char := range input {
		if char == ' ' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}

	if word != "" {
		words = append(words, word)
	}

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	result := ""
	for i, word := range words {
		if i > 0 {
			result += " "
		}
		result += word
	}

	return result
}

func main() {
	input := "snow dog sun"
	result := reverseWords(input)
	fmt.Println("Исходная строка:", input)
	fmt.Println("Перевёрнутая строка:", result)
}
