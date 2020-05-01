package main

import "fmt"

func main() {
	//fmt.Println(Solution("({}){}()()[]()()()()"))
	//fmt.Println(Solution("(}{}{()()[]()(){(()()["))
	fmt.Println(Solution("(    )"))
}

func Solution(input string) bool {
	if len(input) == 0 {
		return true
	}
	openB1 := "("
	openB2 := "["
	openB3 := "{"
	closeB1 := ")"
	closeB2 := "]"
	closeB3 := "}"
	if string(input[0]) == closeB1 || string(input[0]) == closeB2 || string(input[0]) == closeB3 {
		return false
	}

	var stack []string
	for _, c := range input {
		cStr := string(c)
		if cStr == " " {
			continue
		}
		if cStr == openB1 || cStr == openB2 || cStr == openB3 {
			stack = append(stack, cStr)
			continue
		}
		//else pop from the stack
		lenStack := len(stack)
		if lenStack == 0 {
			break
		}
		poppedUp := stack[lenStack-1]
		// maintain the stack
		stack = stack[:lenStack-1]
		switch poppedUp {
		case openB1:
			if cStr == closeB1 {
				continue
			} else {
				return false
			}
		case openB2:
			if cStr == closeB2 {
				continue
			} else {
				return false
			}
		case openB3:
			if cStr == closeB3 {
				continue
			} else {
				return false
			}
		}
	}

	return true
}
