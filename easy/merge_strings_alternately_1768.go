package easy

import (
	"interview_go/common"
	"strings"
)

type MergeStringsAlternately struct{}

func (m *MergeStringsAlternately) mergeAlternately(word1 string, word2 string) string {

	if len(word1) == 0 && len(word2) == 0 {
		return ""
	}

	if len(word1) == 0 {
		return word2
	}
	if len(word2) == 0 {
		return word1
	}

	sb := &strings.Builder{}

	for i := 0; i < len(word1); i++ {

		sb.WriteByte(word1[i])

		if i < len(word2) {
			sb.WriteByte(word2[i])
		}
	}

	if len(word2) > len(word1) {
		sb.WriteString(word2[len(word1):])
	}
	return sb.String()
}

func (m *MergeStringsAlternately) Test() {

	word1, word2 := "cf", "eee"
	e := "cefee"
	r := m.mergeAlternately(word1, word2)
	common.PrintStr(r, e)

	// word3, word4 := "ab", "pqrs"
	// e2 := "apbqrs"
	// r2 := m.mergeAlternately(word3, word4)
	// common.PrintStr(r2, e2)

	// word5, word6 := "abcd", "pq"
	// e3 := "apbqcd"
	// r3 := m.mergeAlternately(word5, word6)
	// common.PrintStr(r3, e3)

	// word7 := "abcd"
	// e4 := "abcd"
	// r4 := m.mergeAlternately(word7, "")
	// common.PrintStr(r4, e4)

	// word8 := "abcd"
	// e5 := "abcd"
	// r5 := m.mergeAlternately("", word8)
	// common.PrintStr(r5, e5)

	// e6 := ""
	// r6 := m.mergeAlternately("", "")
	// common.PrintStr(r6, e6)

}
