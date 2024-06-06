/*
Implement the myAtoi(string s) function, which converts a string to a 32-bit
signed integer.

The algorithm for myAtoi(string s) is as follows:

	Whitespace: Ignore any leading whitespace (" ").
	Signedness: Determine the sign by checking if the next character is '-'
	            or '+', assuming positivity is neither present.
	Conversion: Read the integer by skipping leading zeros until a non-digit
	            character is encountered or the end of the string is reached.
	            If no digits were read, then the result is 0.
	Rounding: If the integer is out of the 32-bit signed integer range
	          [-231, 231 - 1], then round the integer to remain in the range.
	          Specifically, integers less than -231 should be rounded to -231,
	          and integers greater than 231 - 1 should be rounded to 231 - 1.

Return the integer as the final result.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("myAtoi('%v') = %v\n", "42", myAtoi("42"))
	fmt.Printf("myAtoi('%v') = %v\n", "-042", myAtoi("-042"))
	fmt.Printf("myAtoi('%v') = %v\n", "   -042", myAtoi("   -042"))
	fmt.Printf("myAtoi('%v') = %v\n", "1337c0d3", myAtoi("1337c0d3"))
	fmt.Printf("myAtoi('%v') = %v\n", "0-1", myAtoi("0-1"))
	fmt.Printf("myAtoi('%v') = %v\n", "words and 987", myAtoi("words and 987"))
}

func myAtoi(s string) int {
	result := 0
	negative := false
	started := false
loop:
	for _, c := range s {
		switch c {
		case 32:
			if started {
				break loop
			}
		case 43:
			if started {
				break loop
			}
			started = true
		case 45:
			if started {
				break loop
			}
			negative = true
			started = true
		case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57:
			result = result*10 + int(c) - 48
			if result > math.MaxInt32 && negative {
				return math.MinInt32
			} else if result > math.MaxInt32 {
				return math.MaxInt32
			}
			started = true
		default:
			break loop
		}
	}
	if negative {
		result *= -1
	}
	return result
}
