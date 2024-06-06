package main 

import (
	"testing"
)

func BenchmarkFindMedianOfSortedArrays(b *testing.B) {
  count := 10000
	nums1 := make([]int, 0)
	nums2 := make([]int, 0)
	for i := 1; i <= count; i++ {
		if i%2 == 0 {
			nums1 = append(nums1, i)
		} else {
			nums2 = append(nums2, i)
		}
	}
	b.ResetTimer()

	for i := 0; i < 500000; i++ {
    main.FindMedianSortedArrays(nums1, nums2)
	}
}
