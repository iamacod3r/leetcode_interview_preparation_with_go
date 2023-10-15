package medium

import "interview_go/common"

// https://leetcode.com/problems/ways-to-split-array-into-three-subarrays/
// https://leetcode.com/problems/ways-to-split-array-into-three-subarrays/discuss/999257/C%2B%2BJavaPython-O(n)-with-picture

type WaysToSplit_1712 struct{}

func (*WaysToSplit_1712) WaysToSplit(nums []int) int {

	// https://github.com/cihandokur

	size := len(nums)
	result := 0

	if size < 3 {
		return result
	}

	modValue := 1_000_000_007

	for i := 1; i < size; i++ {
		nums[i] += nums[i-1]
	}

	i := 0
	j := 1
	k := 1
	for ; i < size-2; i++ {
		if j < i+1 {
			j = i + 1
		}

		for j < size-1 && nums[i]*2 > nums[j] {
			j++
		}

		if k < j {
			k = j
		}

		for k < size-1 && nums[size-1]-nums[k] >= nums[k]-nums[i] {
			k++
		}

		result = (result + k - j) % modValue
	}

	return result
}

func (w *WaysToSplit_1712) Test() {
	nums1 := [3]int{1, 1, 1}
	e1 := 1
	r1 := w.WaysToSplit(nums1[:])
	common.PrintInt(r1, e1)

	nums2 := [6]int{1, 2, 2, 2, 5, 0}
	e2 := 3
	r2 := w.WaysToSplit(nums2[:])
	common.PrintInt(r2, e2)

	nums3 := [6]int{3, 2, 1}
	e3 := 0
	r3 := w.WaysToSplit(nums3[:])
	common.PrintInt(r3, e3)
}
