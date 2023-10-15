package sorting

// https://www.programiz.com/dsa/insertion-sort
/*
	Time Complexity : Best  O(n)	Worst O(n^2)	Avarage O(n^2)
	Space Complexity : O(1)
*/
func Insertion(slice []int) {
	size := len(slice)

	for step := 1; step < size; step++ {
		val := slice[step]

		i := step - 1

		// Compare val with each element on the left of it until an element smaller than
		// it is found.
		// For descending order, change val<slice[i] to val>slice[i].
		for i >= 0 && val < slice[i] {
			slice[i+1] = slice[i]
			i--
		}

		// Place val at after the element just smaller than it.
		slice[i+1] = val
	}
}
