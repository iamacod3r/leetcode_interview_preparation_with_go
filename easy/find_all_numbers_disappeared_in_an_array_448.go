package easy

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
type FindAllNumbersDisappearedInAnArray struct {
}

func (f *FindAllNumbersDisappearedInAnArray) findDisappearedNumbers(nums []int) []int {

	l := len(nums)
	var result []int

	for i := 0; i < l; i++ {
		idx := int(math.Abs(float64(nums[i])) - 1)
		if nums[idx] > 0 {
			nums[idx] *= -1
		}
	}

	fmt.Println(nums)

	for i := 1; i <= l; i++ {
		if nums[i-1] > 0 {
			result = append(result, i)
		}
	}

	return result
}

func (f *FindAllNumbersDisappearedInAnArray) Test() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	e := []int{5, 6}
	r := f.findDisappearedNumbers(nums)

	fmt.Println(e, r)
}
