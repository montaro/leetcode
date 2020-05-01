package main

import (
	"fmt"
)

func main() {
	//input := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	//input := []int{6, 1, 4, 5, 3, 9, 0, 1, 9, 5, 1, 8, 6, 7, 0, 5, 5, 4, 9}
	input := []int{9, 9}
	fmt.Println(plusOne(input))
}

func plusOne(digits []int) []int {
	slice := digits
	arraySize := len(digits) - 1
	carry := 0
	for i := arraySize; i >= 0; i-- {
		if digits[i] == 9 {
			slice[i] = 0
			carry = 1
		} else {
			carry = 0
			slice[i] = digits[i] + 1
			break
		}
	}
	if carry == 1 {
		newSlice := make([]int, 1)
		newSlice[0] = 1
		newSlice = append(newSlice, slice...)
		return newSlice
	}
	return slice
}
