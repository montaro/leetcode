package main

func main() {
	
}

func Solution(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	result := numbers[0]
	for _, n := range numbers {
		if n > result {
			result = n
		}
	}
	return result
}