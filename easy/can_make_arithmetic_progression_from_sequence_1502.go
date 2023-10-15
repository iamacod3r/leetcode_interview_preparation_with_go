package easy

import (
	"interview_go/common"
	"sort"
)

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/
type CanMakeArithmeticProgressionFromSequence struct {
}

// Time Complexity : O(NlogN)
// Space Complexity: O(N) or O(logN)
func (p *CanMakeArithmeticProgressionFromSequence) canMakeArithmeticProgression_withSort(arr []int) bool {

	if len(arr) < 3 {
		return true
	}

	sort.Ints(arr)

	diff := arr[1] - arr[0]

	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}

// Time Complexity : O(n)
// Space Complexity : O(1)
func (p *CanMakeArithmeticProgressionFromSequence) canMakeArithmeticProgression_inPlace(arr []int) bool {
	minValue, maxValue := arr[0], arr[0]
	alen := len(arr)
	for i := 1; i < alen; i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}

		if minValue > arr[i] {
			minValue = arr[i]
		}
	}

	if (maxValue-minValue)%(alen-1) != 0 {
		return false
	}

	diff := (maxValue - minValue) / (alen - 1)
	i := 0

	for i < alen {
		if arr[i] == minValue+i*diff {
			i++
		} else if (arr[i]-minValue)%diff != 0 {
			return false
		} else {
			j := (arr[i] - minValue) / diff
			if arr[i] == arr[j] {
				return false
			}

			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return true
}

func (p *CanMakeArithmeticProgressionFromSequence) Test() {

	arr := []int{3, 5, 1}
	e := true
	r := p.canMakeArithmeticProgression_inPlace(arr)
	common.PrintBool(r, e)
	arr2 := []int{1, 2, 4}
	e2 := false
	r2 := p.canMakeArithmeticProgression_inPlace(arr2)
	common.PrintBool(r2, e2)
}
