package easy

import (
	"fmt"
	"interview_go/common"
)

// https://leetcode.com/problems/summary-ranges/
// Time Complexity : O(N)
// Space Complexity : O(1)
type SummaryRange struct {
}

func (x *SummaryRange) summaryRange(nums []int) []string {
	result := []string{}
	if len(nums) < 1 {
		return result
	}

	for i := 0; i < len(nums); i++ {
		start := nums[i]
		// Keep iterating until the next element is one more than the current element.
		for i+1 < len(nums) && nums[i]+1 == nums[i+1] {
			i++
		}

		if start != nums[i] {
			result = append(result, fmt.Sprintf("%d->%d", start, nums[i]))
		} else {
			result = append(result, fmt.Sprint(start))
		}
	}

	return result
}

func (x *SummaryRange) Test() {
	nums := []int{0, 1, 2, 4, 5, 7}
	e := []string{"0->2", "4->5", "7"}
	r := x.summaryRange(nums)
	common.PrintStrSlice(r, e)
}
