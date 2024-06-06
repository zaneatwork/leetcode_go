/*
Given a signed 32-bit integer x, return x with its digits reversed.
If reversing x causes the value to go outside the signed 32-bit integer
range [-231, 231 - 1], then return 0.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("The reverse of %v is %v\n", 123, reverse(123))
	fmt.Printf("The reverse of %v is %v\n", -123, reverse(-123))
	fmt.Printf("The reverse of %v is %v\n", 120, reverse(120))
}

func reverse(x int) int {
	result := 0
	for y := x; y != 0; y = y / 10 {
		result = result*10 + y%10
	}
	if result >= math.MaxInt32 || result <= math.MinInt32 {
		return 0
	}
	return result
}
