package easy

import (
	"interview_go/common"
	"strings"
)

// https://leetcode.com/problems/excel-sheet-column-title/description/
type ExcelSheetColumnTitle struct{}

func (e *ExcelSheetColumnTitle) convertToTitle(columnNumber int) string {

	base := 26
	chars := []rune{}

	for columnNumber > 0 {
		columnNumber--
		mod := columnNumber % base
		columnNumber /= base
		chars = append(chars, rune(65+mod))
	}

	sb := &strings.Builder{}

	for i := len(chars) - 1; i >= 0; i-- {
		sb.WriteRune(chars[i])
	}

	return sb.String()
}

func (s *ExcelSheetColumnTitle) Test() {
	columnNumber := 2002
	e := "BXZ"
	r := s.convertToTitle(columnNumber)
	common.PrintStr(r, e)
}
