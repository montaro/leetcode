package main

func main() {
	println(removeElement([]int{1, 2, 3, 4, 5, 4, 4}, 4))
	println(removeElement([]int{3, 2, 2, 3}, 3))
}

func removeElement(nums []int, val int) int {
	k := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
