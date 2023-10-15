package medium

import (
	"interview_go/common"
)

type FindFirstAndLastPositionOfElementInSortedArray_34 struct {
}

func (p *FindFirstAndLastPositionOfElementInSortedArray_34) Test() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	expected := []int{3, 4}
	result := p.searchRange(nums, target)
	common.PrintIntSlice(result, expected)
	expected = []int{-1, -1}
	result = p.searchRange(nums, 6)
	common.PrintIntSlice(result, expected)
	nums = []int{}
	expected = []int{-1, -1}
	result = p.searchRange(nums, 0)
	common.PrintIntSlice(result, expected)
}

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
func (*FindFirstAndLastPositionOfElementInSortedArray_34) searchRange(nums []int, target int) []int {

	size := len(nums)
	if size == 0 {
		return []int{-1, -1}
	}

	// built-in binarysearch for int[] slices
	// sort.SearchInts(nums, target)

	low, high := 0, size-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {

			begin, end := mid, mid

			for begin > 0 && nums[begin-1] == target {
				begin--
			}
			for end < size-1 && nums[end+1] == target {
				end++
			}
			return []int{begin, end}
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return []int{-1, -1}
}
