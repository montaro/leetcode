package main

import (
"fmt"
)

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for i, n := range nums {
		val, ok := m[nums[i]]
		if ok {
			m[n] = val + 1
		} else {
			m[n] = 1
		}
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}
