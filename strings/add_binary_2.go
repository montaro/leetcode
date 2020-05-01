package main

import (
	"fmt"
)

func main() {

	//TestBitAdd()
}

func addBinary2(a string, b string) string {
	lenA := len(a)
	lenB := len(b)

	result := ""
	// loop over b
	for i:=0; i < lenB; i++ {
		carry := "0"
		sum, carry := bitAdd(string(a[i]), carry)
		sum, carry = bitAdd(string(b[i]), sum)
		result += sum
	}
	if lenA > lenB {
	} else {
		// loop over a
		for i, _ := range a {
			carry := "0"
			sum, carry := bitAdd(string(a[i]), carry)
			sum, carry = bitAdd(string(b[i]), sum)
			result += sum
		}

	}
	return result
}

func bitAdd(a, b string) (sum, carry string) {
	carry = "0"
	if a == "0" && b == "0" {
		sum = "0"
	} else if a == "1" && b == "0" {
		sum = "1"
	} else if a == "0" && b == "1" {
		sum = "1"
	} else {
		sum = "1"
		carry = "1"
	}
	return sum, carry
}

func TestBitAdd() {
	fmt.Println(bitAdd("0", "0"))
	fmt.Println(bitAdd("0", "1"))
	fmt.Println(bitAdd("1", "0"))
	fmt.Println(bitAdd("1", "1"))
}
