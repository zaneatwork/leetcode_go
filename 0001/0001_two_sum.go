// Given an array of integers "nums" and an integer "target" return indices of
// the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may
// not use the same element twice.
// You can return the answer in any order.

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
  target := 9
  result := twoSum(nums, target)
  fmt.Printf("nums: %v, target: %v, result: %v\n", nums, target, result)

	nums = []int{3, 2, 4}
  target = 6
  result = twoSum(nums, target)
  fmt.Printf("nums: %v, target: %v, result: %v\n", nums, target, result)

	nums = []int{3, 3}
  target = 6
  result = twoSum(nums, target)
  fmt.Printf("nums: %v, target: %v, result: %v\n", nums, target, result)
}

func twoSum(nums []int, target int) []int {
	var numToIndex = make(map[int]int)
	for idx, num := range nums {
		if otherIdx, ok := numToIndex[target-num]; ok {
			return []int{otherIdx, idx}
		}
		numToIndex[num] = idx
	}
	return []int{}
}
