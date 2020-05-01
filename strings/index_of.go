package main

import "fmt"

func main() {
	s := "montaro"
	rs := ""
	for _, c := range s {
		rs += string(c)
	}
	fmt.Println(rs)
	//println(strStr("mississippi", "mississippi"))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i, c := range haystack {
		if string(c) == string(needle[0]) {
			//check the rest of needle
			if len(needle) > (len(haystack) - i) {
				return -1
			}
			if checkStr(haystack, i, needle) {
				return i
			}
		}
	}
	return -1
}

func checkStr(haystack string, index int, needle string) bool {
	for i := 0; i < len(needle); i++ {
		if string(needle[i]) != string(haystack[index]) {
			return false
		}
		index += 1
	}
	return true
}
