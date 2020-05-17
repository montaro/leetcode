package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	var index int
	for i := 0; i < k-1; i++ {
		index = largestIndex(nums)
		nums = append(nums[:index], nums[index+1:]...)
	}
	kIndex := largestIndex(nums)
	return nums[kIndex]
}

func largestIndex(nums []int) int {
	largest := nums[0]
	index := 0
	for i, n := range nums {
		if n >= largest {
			largest = n
			index = i
		}
	}
	return index
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
