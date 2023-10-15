package easy

import "interview_go/common"

// https://leetcode.com/problems/two-sum-iii-data-structure-design/
type TwoSumIIIDataStructureDesign struct{}

type TwoSum struct {
	nums map[int]int
}

func ConstructorForTwoSum() TwoSum {
	return TwoSum{nums: make(map[int]int)}
}

func (t *TwoSum) Add(number int) {
	t.nums[number]++
}

func (t *TwoSum) Find(value int) bool {

	for n := range t.nums {

		if count, ok := t.nums[value-n]; ok {
			if n != value-n {
				return true
			} else if count > 1 {
				return true
			}
		}
	}

	return false
}

func (t *TwoSumIIIDataStructureDesign) Test() {
	ts := ConstructorForTwoSum()
	ts.Add(7)
	ts.Add(3)
	ts.Add(3)
	ts.Add(5)
	ts.Add(5)
	e1 := true
	r1 := ts.Find(6)
	common.PrintBool(r1, e1)
	e2 := false
	r2 := ts.Find(7)
	common.PrintBool(r2, e2)
}
