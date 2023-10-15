package easy

import "interview_go/common"

// https://leetcode.com/problems/fibonacci-number/

type FibonacciNumber struct {
}

func (f *FibonacciNumber) fib(n int) int {
	if n <= 1 {
		return n
	}

	curr := 0
	prev1 := 1
	prev2 := 0

	for i := 2; i <= n; i++ {
		curr = prev1 + prev2
		prev2 = prev1
		prev1 = curr
	}

	return curr
}

func (f *FibonacciNumber) Test() {

	r1 := f.fib(2)
	e1 := 1
	common.PrintInt(r1, e1)
}
