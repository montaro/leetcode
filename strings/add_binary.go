package main

import (
	"fmt"
	"math/bits"
	"strconv"
)

func main() {
	s1 := "1011"
	s2 := "1101"
	fmt.Println(addBinary(s1, s2))
}

func addBinary(a string, b string) string {
	aa, _ := strconv.ParseInt(a, 2, 64)
	bb, _ := strconv.ParseInt(b, 2, 64)
	sum, carry := bits.Add64(uint64(aa), uint64(bb), 0)
	fmt.Println("Sum: ", sum)
	fmt.Println("Carry: ", carry)
	return strconv.FormatUint(sum, 2)
}