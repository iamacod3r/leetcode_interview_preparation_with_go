package easy

import "interview_go/common"

// https://leetcode.com/problems/two-sum/description/
type TwoSum1 struct {
}

func (t *TwoSum1) twoSum(nums []int, target int) []int {
	remainMap := map[int]int{}
	result := []int{}

	for i := 0; i < len(nums); i++ {

		if v, ok := remainMap[target-nums[i]]; ok {
			result = append(result, v, i)
			break
		}

		remainMap[nums[i]] = i
	}

	return result
}

func (t *TwoSum1) Test() {
	nums := []int{2, 7, 11, 15}
	target := 9
	e := []int{0, 1}
	r := t.twoSum(nums, target)
	common.PrintIntSlice(r, e)
}
