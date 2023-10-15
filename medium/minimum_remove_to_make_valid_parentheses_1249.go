package medium

import (
	"interview_go/common"
	"strings"
)

// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/description/
type MinimumRemoveToMakeValidParentheses struct {
}

func (p *MinimumRemoveToMakeValidParentheses) minRemoveToMakeValid(s string) string {
	sb := strings.Builder{}
	open, balance := 0, 0

	for i := 0; i < len(s); i++ {
		curr := s[i]
		if curr == '(' {
			balance++
			open++
		}

		if curr == ')' {
			if balance == 0 {
				continue
			}
			balance--
		}

		sb.WriteByte(curr)
	}

	result := strings.Builder{}
	keep := open - balance

	sbStr := sb.String()
	for i := 0; i < len(sbStr); i++ {
		c := sbStr[i]
		if c == '(' {
			keep--
			if keep < 0 {
				continue
			}
		}
		result.WriteByte(c)
	}

	return result.String()
}

// Time Complexity : O(n)
// Space Complexity : O(n)
func (p *MinimumRemoveToMakeValidParentheses) removeInvalidClosing(s string, open, close byte) strings.Builder {
	sb := strings.Builder{}
	balance := 0
	for i := 0; i < len(s); i++ {
		curr := s[i]
		if curr == open {
			balance++
		}

		if curr == close {
			if balance == 0 {
				continue
			}
			balance--
		}
		sb.WriteByte(curr)
	}

	return sb
}

func (p *MinimumRemoveToMakeValidParentheses) ReverseString(str string) string {
	sb := strings.Builder{}

	for i := len(str) - 1; i >= 0; i-- {
		sb.WriteByte(str[i])
	}
	return sb.String()
}

func (p *MinimumRemoveToMakeValidParentheses) minRemoveToMakeValid_with_Two_StringBuilders(s string) string {
	result := p.removeInvalidClosing(s, '(', ')')
	result = p.removeInvalidClosing(p.ReverseString(result.String()), ')', '(')
	return p.ReverseString(result.String())

}

// n is the length of the input string
// Time Complexity : O(n)
// Space Complexity : O(n)
// this one is my submission on leetcode
func (p *MinimumRemoveToMakeValidParentheses) minRemoveToMakeValid_with_Stack(s string) string {
	if len(s) == 0 {
		return s
	}

	removeIndexes := map[int]struct{}{}
	stack := []int{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		}
		if s[i] == ')' {
			if len(stack) == 0 {
				removeIndexes[i] = struct{}{}
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	for len(stack) > 0 {
		removeIndexes[stack[0]] = struct{}{}
		stack = stack[1:]
	}
	sb := strings.Builder{}

	for i := 0; i < len(s); i++ {

		if _, ok := removeIndexes[i]; !ok {
			sb.WriteByte(s[i])
		}
	}

	return sb.String()
}

func (p *MinimumRemoveToMakeValidParentheses) minRemoveToMakeValid_byte(s string) string {

	b := []byte(s)
	open := 0
	for i := 0; i < len(b); i++ {

		if b[i] == '(' {
			open++
		}

		if b[i] == ')' {
			if open == 0 {
				b = append(b[:i], b[i+1:]...)
				i--
			} else {
				open--
			}
		}
	}

	closed := 0
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] == ')' {
			closed++
		}

		if b[i] == '(' {
			if closed == 0 {
				b = append(b[:i], b[i+1:]...)
			} else {
				closed--
			}

		}
	}
	return string(b)
}

func (p *MinimumRemoveToMakeValidParentheses) minRemoveToMakeValid_Stack_And_String(s string) string {
	stack := make([]int, 0)
	str := []byte(s)
	// identify and mark as '*' wrong parenthses
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			stack = append(stack, i)
		} else if str[i] == ')' {
			if len(stack) == 0 {
				str[i] = '*'
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	// remove open parentheses if there
	for len(stack) > 0 {
		index := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		str[index] = '*'
	}
	var res strings.Builder
	for i := 0; i < len(str); i++ {
		if str[i] != '*' {
			res.WriteByte(str[i])
		}
	}
	return res.String()
}

func (p *MinimumRemoveToMakeValidParentheses) Test() {
	s1 := "lee(t(c)o)de)"
	e1 := "lee(t(c)o)de"
	r1 := p.minRemoveToMakeValid_byte(s1)
	common.PrintStr(r1, e1)

	s2 := "a)b(c)d"
	e2 := "ab(c)d"
	r2 := p.minRemoveToMakeValid_byte(s2)
	common.PrintStr(r2, e2)

	s3 := "))(("
	e3 := ""
	r3 := p.minRemoveToMakeValid_byte(s3)
	common.PrintStr(r3, e3)
}
