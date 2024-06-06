/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return
the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).

Constraints:
    nums1.length == m
    nums2.length == n
    0 <= m <= 1000
    0 <= n <= 1000
    1 <= m + n <= 2000
    -106 <= nums1[i], nums2[i] <= 106
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	n1 := []int{1, 3}
	n2 := []int{2}
	result := FindMedianSortedArrays(n1, n2)
	fmt.Printf("Median of:\n%v\nAnd:\n%v\nis:\n%v\n\n", n1, n2, result)
	// Should be 2.0000

	n1 = []int{3, 4}
	n2 = []int{1, 2}
	// Should be 2.5000
	result = FindMedianSortedArrays(n1, n2)
	fmt.Printf("Median of:\n%v\nAnd:\n%v\nis:\n%v\n\n", n1, n2, result)

	n1 = []int{0, 0}
	n2 = []int{0, 0}
	// Should be 0
	result = FindMedianSortedArrays(n1, n2)
	fmt.Printf("Median of:\n%v\nAnd:\n%v\nis:\n%v\n\n", n1, n2, result)

	n1 = []int{1, 2}
	n2 = []int{-1, 3}
	// Should be 1.5000
	result = FindMedianSortedArrays(n1, n2)
	fmt.Printf("Median of:\n%v\nAnd:\n%v\nis:\n%v\n\n", n1, n2, result)

	n1 = []int{}
	n2 = []int{2, 3}
	// Should be 2.5000
	result = FindMedianSortedArrays(n1, n2)
	fmt.Printf("Median of:\n%v\nAnd:\n%v\nis:\n%v\n\n", n1, n2, result)
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays(nums1, nums2)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
  // Always work on the smallest of the two arrays
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	total := len(nums1) + len(nums2)
	half := int(math.Floor(float64(total) / 2))

	leftPartitionStart, rightPartitionEnd := 0, len(nums1)-1
	for {
    var nums1LeftValue, nums1RightValue, nums1PartitionIdx int
		nums1PartitionIdx = int(math.Floor(float64(leftPartitionStart+rightPartitionEnd) / 2))
		if nums1PartitionIdx >= 0 {
			nums1LeftValue = nums1[nums1PartitionIdx]
		} else {
		  nums1LeftValue = math.MinInt
    }
		if nums1PartitionIdx+1 < len(nums1) {
			nums1RightValue = nums1[nums1PartitionIdx+1]
		} else {
		  nums1RightValue = math.MaxInt
    }

		var nums2LeftValue, nums2RightValue, nums2PartitionIdx int
		nums2PartitionIdx = half - (nums1PartitionIdx + 1) - 1
		if nums2PartitionIdx >= 0 {
			nums2LeftValue = nums2[nums2PartitionIdx]
		} else {
      nums2LeftValue = math.MinInt
    }
		if nums2PartitionIdx+1 < len(nums2) {
			nums2RightValue = nums2[nums2PartitionIdx+1]
		} else {
      nums2RightValue = math.MaxInt
    }

		// If we picked the correct partition index, left-hand-side of both arrays
    // are less than the right hand side values.
		if nums1LeftValue <= nums2RightValue && nums2LeftValue <= nums1RightValue {
			// If we have an odd num of elements in the combined array, return middle
			if total%2 > 0 {
				return float64(min(nums1RightValue, nums2RightValue))
			}
			// If we have an even number, we have to get half between the two.
			return float64(max(nums1LeftValue, nums2LeftValue)+min(nums1RightValue, nums2RightValue)) / 2
		} else if nums1LeftValue > nums2RightValue {
      // Move our partition point on nums1 back 1
			rightPartitionEnd = nums1PartitionIdx - 1
		} else {
			leftPartitionStart = nums1PartitionIdx + 1
		}
	}
}
