package main

import (
	"fmt"
	"strings"
)

func cleanSymbols(text string) string {
	res := ""
	symbols := "!?',;."
	for _, c := range text {
		if strings.Contains(symbols, string(c)) {
			res = res + " "
		} else {
			res = res + string(c)
		}
	}
	return res
}

func toLower(words []string) []string {
	var result []string
	for _, word := range words {
		 result = append(result, strings.ToLower(word))
	}
	return result
}

func searchString(words []string, word string) bool {
	for _, w := range words {
		if w == word {
			return true
		}
	}
	return false
}

func mostCommonWord(paragraph string, banned []string) string {
	paragraph = cleanSymbols(paragraph)
	words := strings.Split(paragraph, " ")
	words = toLower(words)
	banned = toLower(banned)
	frequency := make(map[string]int)
	for _, word := range words {
		if !searchString(banned, word) {
			if v, ok := frequency[word]; ok {
				frequency[word] = v + 1
			} else {
				frequency[word] = 1
			}
		}
	}

	result := ""
	top := 0
	for k, v := range frequency {
		if v > top {
			top = v
			result = k
		}
	}
	fmt.Println(words)
	fmt.Println(frequency)
	return result
}

func main() {
	fmt.Println(cleanSymbols("Are Ahmed and Doaa here?!"))
	fmt.Println(strings.Split("Ahmed and Doaa", " "))
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
	fmt.Println(mostCommonWord("a.", []string{}))
	fmt.Println(mostCommonWord("a, a, a, a, b,b,b,c, c", []string{"a"}))
}
