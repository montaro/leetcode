package main

import "fmt"
import "strconv"

func main() {
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) [] string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result[i-1] = "FizzBuzz"
			continue
		} else if i%3 == 0 {
			result[i-1] = "Fizz"
			continue
		} else if i%5 == 0 {
			result[i-1] = "Buzz"
			continue
		}
		result[i-1] = strconv.Itoa(i)
	}
	return result
}
