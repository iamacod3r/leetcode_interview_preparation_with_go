package medium

import (
	"interview_go/common"
)

// https://leetcode.com/problems/basic-calculator-ii/
type BasicCalculatorII_227 struct {
}

func (t *BasicCalculatorII_227) calculate(s string) int {

	if len(s) == 0 {
		return 0
	}

	l := len(s)
	currNum, last := 0, 0
	result := 0
	op := byte('+')

	for i := 0; i < l; i++ {
		c := s[i]

		isDigit := t.isDigit(c)

		if isDigit {
			currNum = currNum*10 + int(c-'0')
		}

		if !isDigit && c != ' ' || i == l-1 {
			if op == '+' || op == '-' {
				result += last
				if op == '-' {
					last = -currNum
				} else {
					last = currNum
				}
			} else if op == '*' {
				last *= currNum
			} else if op == '/' {
				last /= currNum
			}
			currNum = 0
			op = c
		}
	}
	result += last
	return result
}

func (t *BasicCalculatorII_227) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (t *BasicCalculatorII_227) Test() {
	// case 1
	s1 := "3+2*2"
	e1 := 7
	r1 := t.calculate(s1)
	common.PrintInt(r1, e1)

	// case 2
	s2 := " 3/2 "
	e2 := 1
	r2 := t.calculate(s2)
	common.PrintInt(r2, e2)

	// case 3
	s3 := " 3+5 / 2 "
	e3 := 5
	r3 := t.calculate(s3)
	common.PrintInt(r3, e3)
}
