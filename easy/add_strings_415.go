package easy

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/add-strings/description/
type AddStrings415 struct{}

/*
Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.
You must solve the problem without using any built-in library for handling large integers (such as BigInteger).
You must also not convert the inputs to integers directly.
*/

// N1 is length of nums1
// N2 is length of nums2
// Time Complexity : O(max(N1, N2))
// Space Complexity : O(max(N1, N2)+1)
func (s *AddStrings415) addStrings(num1, num2 string) string {
	sb := strings.Builder{}

	i, j, carry, sum := len(num1)-1, len(num2)-1, 0, 0

	for i >= 0 || j >= 0 || carry > 0 {
		d1, d2 := 0, 0
		if i >= 0 {
			d1 = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			d2 = int(num2[j] - '0')
			j--
		}

		sum = d1 + d2 + carry

		sb.WriteString(strconv.Itoa(sum % 10))
		carry = sum / 10
	}

	if carry != 0 {
		sb.WriteString(strconv.Itoa(carry))
	}

	// reverse string
	tmp := sb.String()
	sb.Reset()
	for i := len(tmp) - 1; i >= 0; i-- {
		sb.WriteByte(tmp[i])
	}

	return sb.String()
}

func (s *AddStrings415) Test() {
	num1, num2 := "11", "123"
	r := s.addStrings(num1, num2)
	println(r)

}
