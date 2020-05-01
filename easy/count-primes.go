package main

import "fmt"

func main() {
	fmt.Println(countPrimes(10))
}

func countPrimes(n int) int {
	result := 0
	for i := 0; i < n; i++ {
		if isPrime(i) {
			result++
		}
	}
	return result
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i * i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}
