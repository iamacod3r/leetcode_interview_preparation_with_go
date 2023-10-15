package easy

import "interview_go/common"

// https://leetcode.com/problems/unique-number-of-occurrences/description/
type UniqueNumberOfOccurrences struct{}

func (u *UniqueNumberOfOccurrences) uniqueOccurrences(arr []int) bool {

	freqMap := map[int]int{}

	for i := 0; i < len(arr); i++ {
		freqMap[arr[i]]++
	}

	for k, v := range freqMap {
		for x, y := range freqMap {
			if k != x && v == y {
				return false
			}
		}
	}

	return true
}

func (u *UniqueNumberOfOccurrences) Test() {
	arr := []int{1, 2, 2, 1, 1, 3}
	e := true
	r := u.uniqueOccurrences(arr)
	common.PrintBool(r, e)
	arr2 := []int{1, 2}
	e2 := false
	r2 := u.uniqueOccurrences(arr2)
	common.PrintBool(r2, e2)

	arr3 := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	e3 := true
	r3 := u.uniqueOccurrences(arr3)
	common.PrintBool(r3, e3)
}
