package sorting

// https://www.programiz.com/dsa/selection-sort
/*
	Time Complexity : Best  O(n^2)	Worst O(n^2)	Avarage O(n^2)
	Space Complexity : O(1)
*/
func Selection(slice []int) {
	size := len(slice)

	for step := 0; step < size; step++ {
		min_idx := step
		swapped := false
		for i := step + 1; i < size; i++ {

			if slice[i] < slice[min_idx] {
				min_idx = i
				swapped = true
			}
		}

		if !swapped {
			break
		}

		slice[step], slice[min_idx] = slice[min_idx], slice[step]
	}
}
