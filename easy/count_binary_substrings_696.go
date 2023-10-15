package easy

import (
	"interview_go/common"
	"math"
)

// https://leetcode.com/problems/count-binary-substrings/
type CountBinarySubstrings struct {
}

func (b *CountBinarySubstrings) countBinarySubstrings(s string) int {
	result, prev, curr := float64(0), float64(0), float64(1)
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			result += math.Min(prev, curr)
			prev = curr
			curr = 1
		} else {
			curr++
		}
	}
	return int(result + math.Min(prev, curr))
}

func (b *CountBinarySubstrings) Test() {
	s1 := "00110011"
	e1 := 6
	r1 := b.countBinarySubstrings(s1)
	common.PrintInt(r1, e1)

	s2 := "10101"
	e2 := 4
	r2 := b.countBinarySubstrings(s2)
	common.PrintInt(r2, e2)
}
