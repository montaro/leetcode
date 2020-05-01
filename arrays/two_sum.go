package main

import "fmt"

func main() {
	//nums = [2, 7, 11, 15], target = 9 // return [0, 1]
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		x := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == x {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
