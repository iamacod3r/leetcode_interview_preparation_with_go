package easy

import "interview_go/common"

type CheckIfItIsStraightLine struct {
}

/*
	Slope Formula : ΔY / ΔX
	Slope between (1,2) and (2,3) =  Y2 - Y1 / X2 - X1
								=> 	  3 - 2  /  2 - 1
								=>      1   /   1 = 1 -> 1 is result
	Slope between (1,2) and (3,4) =  Y2 - Y1 / X2 - X1
								=> 	  4 - 2  /  3 - 1
								=> 	    2   /  2 = 1 -> 1 is result
	Slope between (1,2) and (5,6) =  Y2 - Y1 / X2 - X1
								=> 	  6 - 2  /  5 - 1
								=> 	    4   /  4 = 1 -> 1 is result
	Slope between (1,2) and (6,7) =  Y2 - Y1 / X2 - X1
								=> 	  7 - 2  /  6 - 1
								=> 	    5   /  5 = 1 -> 1 is result

	SO IF WE HAVE THREE POINTS p0,p1,p2, and the slope using p0 and p1 is ΔY1 / ΔX1 and the slope beetween is ΔY2 / ΔX2,
	we will check if these two slopes are equal or not, i.e. ΔY1 / ΔX1 == ΔY2 / ΔX2.
	Since ΔX can be zero as well and it that case dividing by it would cause an issue.
	We can / need to tweak the previous equality equation to convert division into MULTIPLICATION to avoid the divided by zero issues.
	So The new equation WOULD BE :
	************************************************************************************
					ΔY1 * ΔX2 == ΔY2 * ΔX1 -> has to be TRUE
	************************************************************************************

	Algorithm
	1. Find the ΔX and ΔY using the points at index 0 and 1.
	2. Iterate over the indices from 2 to the end of the list, and for each index i find the ΔX, ΔY for points at index 0 and i. <- it's NOT 1 it's i!
	3. Compare the slope calculated in step #1 with that of step #2 using the previous equation.
		like -> s.getDelta(coordinates[i], coordinates[0], 1) s.getDelta(coordinates[i], coordinates[0], 1)
	4. If the equation is not satisfied, return false.
	5. Otherwise, at the end of the loop return true.
*/

func (s *CheckIfItIsStraightLine) checkStraightLine(coordinates [][]int) bool {

	x1, y1 := coordinates[0][0], coordinates[0][1]
	// x2 and y2
	deltaX := coordinates[1][0] - x1
	deltaY := coordinates[1][1] - y1

	for i := 2; i < len(coordinates); i++ {
		currX := coordinates[i][0] - x1
		currY := coordinates[i][1] - y1
		// Check if the slope between points 0 and i, is the same as between 0 and 1.
		if deltaY*currX != deltaX*currY {
			return false
		}
	}

	return true
}

func (s *CheckIfItIsStraightLine) Test() {
	c1 := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7}}
	e1 := true
	r1 := s.checkStraightLine(c1)
	common.PrintBool(r1, e1)

	c2 := [][]int{
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 5},
		{5, 6},
		{7, 7},
	}
	e2 := false
	r2 := s.checkStraightLine(c2)
	common.PrintBool(r2, e2)
}
