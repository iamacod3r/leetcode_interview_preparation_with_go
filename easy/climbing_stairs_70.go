package easy

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/climbing-stairs/
type ClimbingStairs struct{}

// Time Complexity : O(logN)
// Space Complexity : O(1)
func (c *ClimbingStairs) climbStairs_FibAndSqrt(n int) int {
	sqrt5 := math.Sqrt(5)
	phi := (1 + sqrt5) / 2
	psi := (1 - sqrt5) / 2
	nplus := float64(n + 1)
	result := (math.Pow(phi, nplus) - math.Pow(psi, nplus)) / sqrt5

	return int(result)
}

func (c *ClimbingStairs) climStairs_FibMethod(n int) int {
	if n == 1 {
		return 1
	}

	first, second := 1, 2

	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}

	return second
}

func (c *ClimbingStairs) climStairsWithDP(n int) int {
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// recurvise with memoization
func (c *ClimbingStairs) climbStairs(n int) int {
	memo := make([]int, n+1)
	return c.helper(0, n, memo)
}

func (c *ClimbingStairs) helper(i, n int, memo []int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}

	if memo[i] > 0 {
		return memo[i]
	}

	memo[i] = c.helper(i+1, n, memo) + c.helper(i+2, n, memo)
	return memo[i]
}

func (c *ClimbingStairs) Test() {
	n := 2
	r := c.climbStairs_FibAndSqrt(n)
	e := 2
	fmt.Println(r, e)
}
