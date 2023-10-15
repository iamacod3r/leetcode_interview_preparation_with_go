package easy

import "interview_go/common"

// https://leetcode.com/problems/running-sum-of-1d-array/
type RunningSumOf1dArray struct {
}

func (r *RunningSumOf1dArray) runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

func (s *RunningSumOf1dArray) Test() {
	nums := []int{1, 2, 3, 4}
	e := []int{1, 3, 6, 10}
	r := s.runningSum(nums)
	common.PrintIntSlice(r, e)
}
