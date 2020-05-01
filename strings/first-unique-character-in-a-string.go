package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
}

func firstUniqChar(s string) int {
	ms := make(map[int32]int)
	for _, c := range s {
		if v, ok := ms[c]; ok {
			ms[c] = v + 1
		} else {
			ms[c] = 1
		}
	}
	for i, c := range s {
		if ms[c] == 1 {
			return i
		}
	}
	return -1
}
