package easy

import (
	"fmt"
	"log"
)

type FirstBadVersion struct{}

// https://leetcode.com/problems/first-bad-version/description/
func (*FirstBadVersion) fistBadVersion(n int) int {

	left := 1
	right := n
	pivot := 0
	counter := 1
	for left < right {
		log.Println(counter)
		pivot = left + (right-left)/2
		if isBadVersion(pivot) {
			right = pivot
		} else {
			left = pivot + 1
		}

		counter++
	}
	return left
}

func (f *FirstBadVersion) Test() {

	fmt.Println(f.fistBadVersion(5))
}

func isBadVersion(pivot int) bool {
	return pivot > 3
}
