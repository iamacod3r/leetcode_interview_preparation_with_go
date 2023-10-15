package medium

import (
	"interview_go/common"
)

// https://leetcode.com/problems/sum-of-subarray-minimums/
type SumOfSubarrayMinimums struct {
}

// Time Complexity : O(n)
// Space Complexity : O(n)
func (s *SumOfSubarrayMinimums) sumSubarrayMins(arr []int) int {
	MOD := int64(1000000007)
	stack := []int{}
	dp := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {

		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			prev := stack[len(stack)-1]
			dp[i] = dp[prev] + (i-prev)*arr[i]
		} else {
			dp[i] = (i + 1) * arr[i]
		}

		stack = append(stack, i)
	}

	sum := int64(0)

	for _, v := range dp {
		sum += int64(v)
		sum %= MOD
	}

	return int(sum)
}

func (s *SumOfSubarrayMinimums) Test() {
	a1 := []int{8, 6, 3, 5, 4, 9, 2}
	e1 := 100
	r1 := s.sumSubarrayMins(a1)
	common.PrintInt(r1, e1)

	// a2 := []int{11, 81, 94, 43, 3}
	// e2 := 444
	// r2 := s.sumSubarrayMins(a2)
	// common.PrintInt(r2, e2)
}
