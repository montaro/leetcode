package main

import "fmt"

func main() {
	n := []int{0, 1, 0, 3, 12}
	moveZeroes(n)
	fmt.Println(n)
}

func moveZeroes(nums []int) {
	lastIndexOfNonZero := 0
	for _, n := range nums {
		if n != 0 {
			nums[lastIndexOfNonZero] = n
			lastIndexOfNonZero++
		}
	}
	n := len(nums)
	for i := lastIndexOfNonZero; i < n; i++ {
		nums[i] = 0
	}
}
