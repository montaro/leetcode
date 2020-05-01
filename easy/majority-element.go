package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2, 3, 3, 4, 4, 4 ,4 ,4, 4, 4, 4, 4, 4}))
}

func majorityElement(nums []int) int {
	majority := len(nums) / 2
	key := nums[0]
	m := make(map[int]int)
	for _, n := range nums {
		if v, ok := m[n]; ok {
			m[n] = v + 1
		} else {
			m[n] = 1
		}
		if m[n] > majority {
			key = n
		}
	}
	return key
}