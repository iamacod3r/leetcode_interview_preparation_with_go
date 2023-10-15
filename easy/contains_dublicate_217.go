package easy

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/contains-duplicate/description/

type ContainsDuplicate struct {
}

func (*ContainsDuplicate) containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	check := make(map[int]struct{})

	for _, n := range nums {
		fmt.Println(n)
		if _, ok := check[n]; ok {
			return true
		}

		check[n] = struct{}{}
	}

	return false
}

func (*ContainsDuplicate) containsDuplicateWithSort(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

func (c *ContainsDuplicate) Test() {

	nums := [4]int{1, 2, 3, 1}
	e := true

	r := c.containsDuplicateWithSort(nums[:])
	fmt.Println(e, r)
}
