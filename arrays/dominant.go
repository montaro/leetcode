package main

import "fmt"

func main() {
	fmt.Println(dominantIndex([]int{0, 0, 0, 1}))
}

func dominantIndex(nums []int) int {
	allDoubles := doubleElements(nums)
	for i, num := range nums {
		var leftSide = make([]int, len(allDoubles[:i]))
		var rightSide = make([]int, len(allDoubles[i+1:]))
		copy(leftSide, allDoubles[:i])
		copy(rightSide, allDoubles[i+1:])
		doublesToCompare := append(leftSide, rightSide...)
		if isGreaterThanAll(num, doublesToCompare) {
			return i
		}
	}
	return -1
}

func doubleElements(nums []int) []int {
	doubles := make([]int, len(nums))
	for i, num := range nums {
		doubles[i] = num * 2
	}
	return doubles
}

func isGreaterThanAll(num int, nums []int) bool {
	for _, n := range nums {
		if num < n {
			return false
		}
	}
	return true
}