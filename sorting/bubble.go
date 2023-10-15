package sorting

// https://www.programiz.com/dsa/bubble-sort
/*
	Time Complexity : Best  O(n)	Worst O(n^2)	Avarage O(n^2)
	Space Complexity : O(1)
*/
func Bubble(slice []int) {
	size := len(slice)

	for i := 0; i < size-1; i++ {
		swapped := false
		for j := 0; j < size-i-1; j++ {
			if slice[j] > slice[j+1] {
				swapped = true
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}

		if !swapped {
			break
		}
	}

}
