package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("Ahmed: demha"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("car"))
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	s = strings.ToLower(s)
	var newS []rune
	for _, c := range s {
		if isValid(c) {
			newS = append(newS, c)
		}
	}

	var reversedS []rune
	for i := len(newS) - 1; i >= 0; i-- {
		reversedS = append(reversedS, newS[i])
	}

	for i := 0; i < len(newS); i++ {
		if newS[i] != reversedS[i] {
			return false
		}
	}
	return true
}

func isValid(b rune) bool {
	return 'a' <= b && b <= 'z' || '0' <= b && b <= '9'
}
