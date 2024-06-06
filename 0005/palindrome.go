/*
Given a string s, return the longest palindromic substring in s.
Constraints:
    1 <= s.length <= 1000
    s consist of only digits and English letters.
*/

package main

import (
	"fmt"
)

func main() {
	s1, s2, s3 := "babad", "bab", " "
	fmt.Println(longestPalindrome(s1))
	fmt.Println(longestPalindrome(s2))
	fmt.Println(longestPalindrome(s3))
}

func longestPalindrome(s string) string {
	result := ""
	updatePalindromeStartingAt := func(leftIdx, rightIdx int) {
		for l, r := leftIdx, rightIdx; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			if (r - l + 1) > len(result) {
				result = s[l : r+1]
			}
		}
	}
	for i := range len(s) {
		updatePalindromeStartingAt(i, i)
		updatePalindromeStartingAt(i, i+1)
	}

	return result
}
