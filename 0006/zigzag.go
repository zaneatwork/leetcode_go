/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of
rows like this:

rows  | distance between letters
1       0
2       1
3       3
4       5
5       7
6       9

P   A   H   N
A P L S I I G
Y   I   R

P     I    N
A   L S  I G
Y A   H R
P     I

P       H
A     S I
Y   I   R
P L     I G
A       N

P         R
A       I I
Y     H   N
P   S     G
A I
L

And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number
of rows.

Constraints:

	1 <= s.length <= 1000
	s consists of English letters (lower-case and upper-case), ',' and '.'.
	1 <= numRows <= 1000
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "PAYPALISHIRING"
	fmt.Println(convert(s, 3))
	fmt.Println(convert(s, 4))
}

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
  // String Builder is faster than byte array for multiple appends!
  result := strings.Builder{}

	for row := range numRows {
		jump := 2 * (numRows - 1)
		for i := row; i < len(s); i += jump {
      result.WriteByte(s[i])
			// for middle rows, we need to include the middle zigzag values
			if row > 0 && row < numRows-1 && (i+jump-2*row) < len(s) {
				result.WriteByte(s[i+jump-2*row])
			}
		}
	}

	return result.String()
}
