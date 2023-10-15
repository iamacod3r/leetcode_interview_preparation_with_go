package easy

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/missing-number/

type MissingNumber struct {
}

func (m *MissingNumber) missingNumber(nums []int) int {
	sort.Ints(nums)
	size := len(nums)
	if nums[size-1] != size {
		return size
	} else if nums[0] != 0 {
		return 0
	}

	for i := 1; i <= size; i++ {
		result := nums[i-1] + 1
		if result != nums[i] {
			return result
		}
	}

	return -1
}

func (m *MissingNumber) missingNumberSimple(nums []int) int {
	s := len(nums)
	for i := range nums {
		s += i - nums[i]
	}
	return s
}

// https://leetcode.com/problems/missing-number/editorial/
func (m *MissingNumber) missingNumberWithGaussFormula(nums []int) int {
	size := len(nums)
	expectedSum := size * (size + 1) / 2
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return expectedSum - sum
}

func (m *MissingNumber) Test() {
	nums := [3]int{3, 0, 1}
	e := 2
	r := m.missingNumberSimple(nums[:])

	fmt.Println(e, r)
}
