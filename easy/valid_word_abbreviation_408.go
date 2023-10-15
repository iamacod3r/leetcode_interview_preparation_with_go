package easy

import (
	"interview_go/common"
	"strconv"
)

// https://leetcode.com/problems/valid-word-abbreviation/
type ValidWordAbbreviation struct {
}

func (v *ValidWordAbbreviation) validWordAbbreviation(word string, abbr string) bool {

	// if len(abbr) > len(word) {
	// 	return false
	// }

	/*
	   0 - i	0 - i
	   1 - n	1 - 1
	   2 - t	2 - 2
	   3 - e	3 - i
	   4 - r	4 - z
	   5 - n	5 - 4
	   6 - a	6 - n
	   7 - t
	   8 - i
	   9 - o
	   10 - n
	   11 - a
	   12 - l
	   13 - i
	   14 - z
	   15 - a
	   16 - t
	   17 - i
	   18 - o
	   19 - n
	   20 - i
	*/
	a, w := 0, 0
	for len(abbr) > a && len(word) > w {

		if word[w] == abbr[a] {
			w++
			a++
		} else {
			if abbr[a] == '0' {
				return false
			}
			prev := a
			for len(abbr) > a && abbr[a] >= '0' && abbr[a] <= '9' {
				a++
			}

			num, err := strconv.Atoi(abbr[prev:a])
			if err != nil {
				return false
			}
			w += num
		}
	}

	return len(word) == w && len(abbr) == a
}

func (v *ValidWordAbbreviation) Test() {
	word := "abbreviation"
	abbr := "a10n"
	e := true
	r := v.validWordAbbreviation(word, abbr)
	common.PrintBool(r, e)

	// word2 := "apple"
	// abbr2 := "a2e"
	// e2 := false
	// r2 := v.validWordAbbreviation(word2, abbr2)
	// common.PrintBool(r2, e2)
}
