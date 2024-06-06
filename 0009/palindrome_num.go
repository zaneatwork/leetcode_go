/*
Given an integer x, return true if x is a palindrome, and false otherwise.
*/
package main

import "fmt"

func main() {
  fmt.Printf("isPalindrome(%v) is %v\n", 121, isPalindrome(121))
  fmt.Printf("isPalindrome(%v) is %v\n", -121, isPalindrome(-121))
  fmt.Printf("isPalindrome(%v) is %v\n", 10, isPalindrome(10))
}

func isPalindrome(x int) bool {
  original := x
  reversed := 0

  for x > 0 {
    reversed = reversed * 10 + x % 10
    x /= 10
  }

  return original == reversed
}
