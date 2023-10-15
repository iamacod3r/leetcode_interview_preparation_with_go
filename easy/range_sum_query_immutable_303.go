package easy

import "interview_go/common"

type RangeSumQueryImmutable struct{}

func (r *RangeSumQueryImmutable) Test() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := ConstructorForRangeSumQueryImmutable(nums)
	left, right := 0, 2
	e1 := 1
	r1 := obj.SumRange(left, right)
	common.PrintInt(r1, e1)

	left, right = 2, 5
	e1 = -1
	r1 = obj.SumRange(left, right)
	common.PrintInt(r1, e1)

	left, right = 0, 5
	e1 = -3
	r1 = obj.SumRange(left, right)
	common.PrintInt(r1, e1)
}

// https://leetcode.com/problems/range-sum-query-immutable/
type NumArray struct {
	data  []int
	cache map[CacheKey]int
}

func ConstructorForRangeSumQueryImmutable(nums []int) NumArray {
	obj := &NumArray{}
	obj.data = nums
	obj.cache = make(map[CacheKey]int)
	return *obj
}

func (n *NumArray) SumRange(left int, right int) int {

	if left < 0 {
		left = 0
	}

	if right > len(n.data)-1 {
		right = len(n.data) - 1
	}

	sum := 0

	key := CacheKey{}
	key.Left = left
	key.Right = right

	if v, ok := n.cache[key]; ok {
		return v
	}

	for i := left; i <= right; i++ {
		sum += n.data[i]
	}
	n.cache[key] = sum
	return sum
}

type CacheKey struct {
	Left, Right int
}
