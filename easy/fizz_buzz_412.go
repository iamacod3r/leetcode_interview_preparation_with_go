package easy

import (
	"fmt"
	"strconv"
)

type FizzBuzz_412 struct{}

// https://leetcode.com/problems/fizz-buzz/
func (*FizzBuzz_412) fizzBuzz(n int) []string {
	var result []string

	for i := 1; i <= n; i++ {
		dividedBy3 := (i%3 == 0)
		dividedBy5 := (i%5 == 0)
		tmp := ""
		if dividedBy3 {
			tmp += "Fizz"
		}
		if dividedBy5 {
			tmp += "Buzz"
		}

		if len(tmp) == 0 {
			tmp += strconv.Itoa(i)
		}

		result = append(result, tmp)
	}
	return result
}

func (*FizzBuzz_412) fizzBuzzAdvance(n int) []string {
	var result []string

	wordDict := make(map[int]string)
	wordDict[3] = "Fizz"
	wordDict[5] = "Buzz"

	for i := 1; i <= n; i++ {
		tmp := ""
		for k, v := range wordDict {
			if i%k == 0 {
				tmp += v
			}
		}

		if len(tmp) == 0 {
			tmp += strconv.Itoa(i)
		}

		result = append(result, tmp)
	}
	return result
}

func (f *FizzBuzz_412) Test() {
	fmt.Println(f.fizzBuzzAdvance(15))
}
