/*
Given an input string s and a pattern p, implement regular expression matching
with support for '.' and '*' where:

    '.' Matches any single character.
    '*' Matches zero or more of the preceding element.

The matching should cover the entire input string (not partial).
*/

package main

import "fmt"

func main() {
	var s, p string
	s = "aa"
	p = "a"
	fmt.Printf("isMatch(%v, %v) %v\n", s, p, isMatch(s, p))

	s = "aa"
	p = "a*"
	fmt.Printf("isMatch(%v, %v) %v\n", s, p, isMatch(s, p))

	s = "ab"
	p = ".*"
	fmt.Printf("isMatch(%v, %v) %v\n", s, p, isMatch(s, p))
}

// a
// a* -> aa, aaa
// ab*a -> aa, aba, abba

func isMatch(s string, p string) bool {
	// do a recursive depth first search through the string and pattern to check
	// for complete match
	cache := make(map[[2]int]bool) // cache previous results, avoid duplicate work

	var depthFirstSearch func(int, int) bool
	depthFirstSearch = func(stringIdx, patternIdx int) bool {
    key := [2]int{stringIdx, patternIdx}
		if val, ok := cache[key]; ok {
			return val
		}

		if stringIdx >= len(s) && patternIdx >= len(p) {
			return true
		} else if patternIdx >= len(p) {
			return false
		}

		match := stringIdx < len(s) && (s[stringIdx] == p[patternIdx] || p[patternIdx] == '.')

		if (patternIdx+1) < len(p) && p[patternIdx+1] == '*' {
			cache[key] = (depthFirstSearch(stringIdx, patternIdx+2) || // don't use *
				(match && depthFirstSearch(stringIdx+1, patternIdx))) // use *
			return cache[key]
		}

		if match {
			cache[key] = depthFirstSearch(stringIdx+1, patternIdx+1)
			return cache[key]
		}

		cache[key] = false
		return false
	}

	return depthFirstSearch(0, 0)
}
