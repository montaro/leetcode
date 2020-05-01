package main

import "fmt"

func main() {
	fmt.Println(isAnagram("ahmed", "hamada"))
	fmt.Println(isAnagram("ahmed", "ahmed"))
	fmt.Println(isAnagram("cat", "tac"))
	fmt.Println(isAnagram("cat", "taco"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sm := make(map[int32]int)
	for _, c := range s {
		if v, ok := sm[c]; ok {
			sm[c] = v + 1
		} else {
			sm[c] = 1
		}
	}

	tm := make(map[int32]int)
	for _, c := range t {
		if v, ok := tm[c]; ok {
			tm[c] = v + 1
		} else {
			tm[c] = 1
		}
	}

	for k, v := range sm {
		if tm[k] != v {
			return false
		}
	}

	return true
}
