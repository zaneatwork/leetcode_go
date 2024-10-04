
/*
You are given an integer array height of length n. There are n vertical lines 
drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the 
container contains the most water.

Return the maximum amount of water a container can store.
Notice that you may not slant the container.
*/

package main

import "fmt"

func main() {
	var height []int
  
  height = []int{1,8,6,2,5,4,8,3,7}
	fmt.Printf("maxArea(%v) %v\n", height, maxArea(height))

  height = []int{1,1}
	fmt.Printf("maxArea(%v) %v\n", height, maxArea(height))
}

func maxArea(height []int) int {
  // left and right pointers
  left, right, area := 0, len(height)-1, 0

  // walk left and right pointers toward the center of the array, 
  // finding the area of each step
  for right > left {
    area = max(area, (right-left)*min(height[left],height[right]))
    if height[left] < height[right] {
      left++
    } else {
      right--
    }
  }

  return area
}
