package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		subMap := make(map[string]interface{})
		subMap[string(s[i])] = ""
		for j := i + 1; j < len(s); j++ {
			if _, ok := subMap[string(s[j])]; !ok {
				subMap[string(s[j])] = ""
			} else {
				break
			}
		}
		if len(subMap) > res {
			res = len(subMap)
		}
	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcbcb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(" "))
}
