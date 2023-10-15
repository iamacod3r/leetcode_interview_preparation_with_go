package easy

import "interview_go/common"

// https://leetcode.com/problems/valid-anagram/description/
type ValidAnagram struct {
}

func (h *ValidAnagram) isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	anagramMap := map[byte]int{}

	for i := 0; i < len(s); i++ {
		anagramMap[s[i]]++
	}

	for i := 0; i < len(t); i++ {

		if v, ok := anagramMap[t[i]]; !ok || v < 1 {
			return false
		}
		anagramMap[t[i]]--
	}

	return true
}

func (h *ValidAnagram) Test() {
	s, t := "anagram", "nagaram"
	e := true
	r := h.isAnagram(s, t)
	common.PrintBool(r, e)

	s2, t2 := "rat", "car"
	e2 := false
	r2 := h.isAnagram(s2, t2)
	common.PrintBool(r2, e2)
}
