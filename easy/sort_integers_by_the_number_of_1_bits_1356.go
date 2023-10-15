package easy

import (
	"interview_go/common"
	"math/bits"
	"sort"
)

// https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/
type SortIntegersByTheNumberOf1Bits struct{}

// Time Complexity : O(n * logn)
// Space Complexity : O(logn)
func (s *SortIntegersByTheNumberOf1Bits) sortByBits(arr []int) []int {

	sort.SliceStable(arr, func(second, first int) bool {
		secondBits := bits.OnesCount(uint(arr[second]))
		firstBits := bits.OnesCount(uint(arr[first]))

		if secondBits == firstBits {
			return arr[second] < arr[first]
		}
		return secondBits < firstBits
	})

	return arr
}

func (s *SortIntegersByTheNumberOf1Bits) countBits(x int) (cnt int) {
	for i := 0; i < 15; i++ {
		if (x & (1 << i)) > 0 {
			cnt += 1
		}
	}
	return
}

func (s *SortIntegersByTheNumberOf1Bits) Test() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	e := []int{0, 1, 2, 4, 8, 3, 5, 6, 7}
	r := s.sortByBits(arr)
	common.PrintIntSlice(r, e)

	arr2 := []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
	e2 := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	r2 := s.sortByBits(arr2)
	common.PrintIntSlice(r2, e2)
}
