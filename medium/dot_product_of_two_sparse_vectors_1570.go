package medium

import "interview_go/common"

// https://leetcode.com/problems/dot-product-of-two-sparse-vectors/description/
type DotProductOfTwoSparseVectors struct {
}

type SparseVectorWithMap struct {
	// index, value
	numMap *map[int]int
}

func ConstructorForMap(nums []int) SparseVectorWithMap {

	numMap := map[int]int{}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			numMap[i] = nums[i]
		}
	}

	return SparseVectorWithMap{numMap: &numMap}
}

// Return the dotProduct of two sparse vectors
func (s *SparseVectorWithMap) dotProduct(vec SparseVectorWithMap) int {

	result := 0
	// smallMap := s.numMap
	// bigMap := vec.numMap
	// if len(*s.numMap) > len(*vec.numMap) {
	// 	smallMap = vec.numMap
	// 	bigMap = s.numMap
	// }
	for k, v := range *s.numMap {

		if z, ok := (*vec.numMap)[k]; ok {
			result += v * z
		}
	}

	return result
}

type IndexValue struct {
	Index int
	Value int
}

type SparseVectorWithPair struct {
	pairs []IndexValue
	// or
	// pairs [][]int
}

func ConstructorForPair(nums []int) SparseVectorWithPair {
	pairs := []IndexValue{}

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			pairs = append(pairs, IndexValue{i, nums[i]})
		}
	}

	return SparseVectorWithPair{pairs: pairs}
}

func (s *SparseVectorWithPair) dotProduct(vec SparseVectorWithPair) int {
	result, p, q := 0, 0, 0

	for p < len(s.pairs) && q < len(vec.pairs) {

		if s.pairs[p].Index == vec.pairs[q].Index {
			result += s.pairs[p].Value * vec.pairs[q].Value
			p++
			q++
		} else if s.pairs[p].Value > vec.pairs[q].Value {
			q++
		} else {
			p++
		}
	}

	return result
}

func (t *DotProductOfTwoSparseVectors) Test() {
	nums1 := []int{1, 0, 0, 2, 3}
	nums2 := []int{0, 3, 0, 4, 0}
	v1 := ConstructorForPair(nums1)
	v2 := ConstructorForPair(nums2)
	e1 := 8
	r1 := v1.dotProduct(v2)
	common.PrintInt(r1, e1)

	nums1 = []int{0, 1, 0, 0, 0}
	nums2 = []int{0, 0, 0, 0, 2}
	v1 = ConstructorForPair(nums1)
	v2 = ConstructorForPair(nums2)
	e1 = 0
	r1 = v1.dotProduct(v2)
	common.PrintInt(r1, e1)

	nums1 = []int{0, 1, 0, 0, 2, 0, 0}
	nums2 = []int{1, 0, 0, 0, 3, 0, 4}
	v1 = ConstructorForPair(nums1)
	v2 = ConstructorForPair(nums2)
	e1 = 6
	r1 = v1.dotProduct(v2)
	common.PrintInt(r1, e1)

}
