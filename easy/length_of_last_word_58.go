package easy

import (
	"interview_go/common"
)

// https://leetcode.com/problems/length-of-last-word/
type LengthOfLastWord struct {
}

func (l *LengthOfLastWord) Test() {
	s := "   fly me   to   the moon  "
	e := 4
	r := l.lengthOfLastWord(s)
	common.PrintInt(e, r)

}

func (*LengthOfLastWord) lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	l := len(s) - 1
	space := byte(32)
	for s[l] == space {
		l--
	}
	i := l
	for i >= 0 && s[i] != space {
		i--
	}

	return l - i
}
