package easy

import (
	"interview_go/common"
)

// https://leetcode.com/problems/valid-palindrome-ii/
type ValidPalindromeII struct {
}

func (v *ValidPalindromeII) validPalindromeHelper(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func (v *ValidPalindromeII) validPalindrome(s string) bool {

	l, r := 0, len(s)-1
	for l < r {

		if s[l] != s[r] {
			return v.validPalindromeHelper(s, l, r-1) || v.validPalindromeHelper(s, l+1, r)
		}
		l++
		r--
	}

	return true
}

func (v *ValidPalindromeII) Test() {
	s := "aba"
	e := true
	r := v.validPalindrome(s)
	common.PrintBool(r, e)
	s = "abca"
	e = true
	r = v.validPalindrome(s)
	common.PrintBool(r, e)
	s = "abc"
	e = false
	r = v.validPalindrome(s)
	common.PrintBool(r, e)
	s = "deeee"
	e = true
	r = v.validPalindrome(s)
	common.PrintBool(r, e)
}
