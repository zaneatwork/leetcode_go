/*
Given a string s, find the length of the longest substring without repeating
characters.

Constraints:

    0 <= s.length <= 5 * 104
    s consists of English letters, digits, symbols and spaces.
*/

package main

import (
	"fmt"
)


func main() {
	s1, s2, s3, s4 := "abcabcbb", "bbbbb", "pwwkew", " "
	fmt.Println(lengthOfLongestSubstring(s1))
	fmt.Println(lengthOfLongestSubstring(s2))
	fmt.Println(lengthOfLongestSubstring(s3))
	fmt.Println(lengthOfLongestSubstring(s4))
}


// Second Try, left and right pointers o(n) complexity
func lengthOfLongestSubstring(s string) int {
  indexes := make(map[byte]int)
  left := 0
  result := 0
  for right := 0; right < len(s); right++ {
    if i, ok := indexes[s[right]]; ok {
      left = max(left, i + 1)
    }
    indexes[s[right]] = right
    result = max(result, right - left + 1)
  }
  return result
}

// First Try, Bruteforcing with goroutines like a dumb-dumb
//func lengthOfLongestSubstring1(s string) int {
//	lens := make(chan int)
//	wg := new(sync.WaitGroup)
//
//	for i := 0; i < len(s); i++ {
//		wg.Add(1)
//		go func(s string) {
//			defer wg.Done()
//			substr := []byte{}
//			for i := 0; i < len(s) && bytes.IndexByte(substr, s[i]) == -1; i++ {
//				substr = append(substr, s[i])
//			}
//			lens <- len(substr)
//		}(s[i:])
//	}
//
//	go func() {
//		wg.Wait()
//		close(lens)
//	}()
//
//	maxLen := 0
//	for l := range lens {
//		if l > maxLen {
//			maxLen = l
//		}
//	}
//	return maxLen
//}
