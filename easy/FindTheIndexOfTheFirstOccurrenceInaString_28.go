package easy

import "interview_go/common"

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
type FindTheIndexOfTheFirstOccurrenceInaString struct {
}

// Sliding Window
// Time Complexity : O (m*n)
// Space Complexity : O(1)
func (f *FindTheIndexOfTheFirstOccurrenceInaString) strStr(haystack string, needle string) int {
	m := len(haystack)
	n := len(needle)

	if m == 0 || n == 0 || n > m {
		return -1
	}

	for w := 0; w < m; w++ {

		if w+n > m {
			break
		}

		for p := 0; p < n; p++ {
			if haystack[w+p] != needle[p] {
				break
			}
			if p == n-1 {
				return w
			}
		}
	}
	return -1
}

func (f *FindTheIndexOfTheFirstOccurrenceInaString) Test() {
	haystack := "sadbutsad"
	needle := "sad"
	e := 0
	r := f.strStr(haystack, needle)
	common.PrintInt(r, e)
}
