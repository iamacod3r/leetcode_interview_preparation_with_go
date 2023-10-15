package medium

import (
	"interview_go/common"
	"math"

	"golang.org/x/exp/slices"
)

// https://leetcode.com/problems/perfect-squares/description/

type PerfectSquares_279 struct {
}

func (p *PerfectSquares_279) Test() {

	n := 12
	e1 := 3
	r1 := p.numSquaresMath(n)
	common.PrintInt(r1, e1)

	n2 := 13
	e2 := 2
	r2 := p.numSquaresMath(n2)
	common.PrintInt(r2, e2)
}

// Time Complexity : O(sqrt(n))
/* ->
In the main loop, we check if the number can be decomposed into the sum of two squares,
which takes O(sqrt(n)) iterations.
In the rest of cases, we do the check in constant time
*/
// Space Complexity : O(1)
// -> The algorithm consumes a constant space, regardless the size of the input number.
func (p *PerfectSquares_279) numSquaresMath(n int) int {
	// https://en.wikipedia.org/wiki/Lagrange%27s_four-square_theorem
	// https://en.wikipedia.org/wiki/Adrien-Marie_Legendre ****

	isSquare := func(n int) bool {
		sq := int(math.Sqrt(float64(n)))
		return n == sq*sq
	}

	// four-square and three-square theorems.
	for n%4 == 0 {
		n /= 4
	}

	if n%8 == 7 {
		return 4
	}

	if isSquare(n) {
		return 1
	}

	// enumeration to check if the number can be decomposed into sum of two squares.
	for i := 1; i*i <= n; i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}

	// bottom case of three-square theorem
	return 3
}

// Time Complexity : O(n^(h/2))
// -> Where h is the height of the N-ary tree.
// Space Complexity : O((sqrt(n))^n)
/* ->
Which is also the maximal number of nodes that can appear at the level h.
As once can see, though we keep a list of square_nums, the main consumption
of the space in the queue variable, which keep track of the remainders to visit
for a given level of N-ary tree.
*/
func (p *PerfectSquares_279) numSquaresGreedyAndBFS(n int) int {

	var square_nums []int
	for i := 1; i*i <= n; i++ {
		square_nums = append(square_nums, i*i)
	}

	queue := make(map[int]struct{})
	queue[n] = struct{}{}
	level := 0

	for len(queue) > 0 {
		level++

		next := make(map[int]struct{})

		for r, _ := range queue {
			for _, s := range square_nums {
				if r == s {
					return level
				} else if r < s {
					break
				} else {
					next[r-s] = struct{}{}
				}
			}
		}

		queue = next
	}
	return level
}

// Time Complexity : O((sqrt(n^(h+1)) -1) / (sqrt(n) - 1))
// -> Where h is the maximal number of recursion that could happen.
// Space Complexity : O(sqrt(n))
/* ->
We keep a list of square_nums, which is of sqrt(n) size.
In addition, we would need additional space for the recursive call stack.
But as we will learn later, the size of the call track would not exceed 4.
*/
func (p *PerfectSquares_279) numSquaresGreedy(n int) int {

	size := int(math.Sqrt(float64(n))) + 1
	square_nums := make(map[int]struct{})

	for i := 1; i <= size; i++ {
		square_nums[i*i] = struct{}{}
	}

	var isDividedBy func(int, int) bool
	isDividedBy = func(n, count int) bool {
		if count == 1 {
			_, ok := square_nums[n]
			return ok
		}

		for k, _ := range square_nums {
			if isDividedBy(n-k, count-1) {
				return true
			}
		}
		return false
	}

	count := 1
	for ; count <= n; count++ {
		if isDividedBy(n, count) {
			return count
		}
	}

	return count
}

// Time Complexity : O(n * sqrt(n))
// -> In main step, we have a nested loop, where the outer loop is of n iterations and in the inner loop it takes at maximum sqrt(n) iterations.
// Space Complexity : O(n)
// -> We keep all the intermediate sub-solutions in the array dp[]
func (p *PerfectSquares_279) minNumSquaresWithDP(n int) int {

	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}

	max_square_index := int(math.Sqrt(float64(n))) + 1
	square_nums := make([]int, max_square_index) // slice

	for i := 1; i < max_square_index; i++ {
		square_nums[i] = i * i
	}

	// https://medium.com/@victorsteven/understanding-data-structures-in-golang-f55205afdcaa
	// this lines for just demo
	var square_nums_slice []int // slice
	/*
		Slices give a more robust interface to sequences compared to arrays.
		Unlike an array, no need to specify the length of the slice when defining it.
	*/
	for i := 1; i < max_square_index; i++ {
		square_nums_slice = append(square_nums_slice, i*i)
	}
	// for array
	var myArray [10]int // array -> size/length is fixed
	myArray[0] = 1
	/*
		This data structure is used to store fixed number of elements.
		So once an array is defined, elements cannot be added or removed from the array.
		Note that, we can set the values for any of the index in the array.
	*/

	for i := 1; i <= n; i++ {
		for s := 1; s < max_square_index; s++ {
			if i < square_nums[s] {
				break
			}
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-square_nums[s]]+1)))
		}
	}

	return dp[n]
}

func (p *PerfectSquares_279) numSquaresBruteForce(n int) int {

	size := int(math.Sqrt(float64(n))) + 1
	square_nums := make([]int, size)

	for i := 1; i <= len(square_nums); i++ {
		square_nums[i-1] = i * i
	}

	var minNumSquares func(int) float64
	minNumSquares = func(k int) float64 {
		if slices.Contains(square_nums, k) {
			return 1
		}

		min_num := math.MaxFloat64

		for _, v := range square_nums {
			if k < v {
				break
			}
			new_num := minNumSquares(k-v) + 1
			min_num = math.Min(min_num, new_num)
		}

		return min_num
	}

	return int(minNumSquares(n))
}
