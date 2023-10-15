package easy

import "fmt"

// https://leetcode.com/problems/single-number/
type SingleNumber struct {
}

func (s *SingleNumber) singleNumber(nums []int) int {

	check := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		check[nums[i]]++
	}

	fmt.Println(check)
	for i := 0; i < len(nums); i++ {
		if check[nums[i]] == 1 {
			return nums[i]
		}
	}

	return 0
}

func (s *SingleNumber) Test() {
	nums := []int{2, 2, 1}
	e := 1
	r := s.singleNumber(nums)
	fmt.Println(r, e)
}
