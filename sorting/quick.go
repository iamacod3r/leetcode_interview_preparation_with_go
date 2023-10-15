package sorting

// https://www.programiz.com/dsa/quick-sort
/*
	Time Complexity : Best  O(n * log_n)	Worst O(n^2)	Avarage O(n * log_n)
	Space Complexity : O(log_n)
*/
func Quick(slice []int) {
	high := len(slice) - 1
	quickSort(slice, 0, high)
}

func quickSort(slice []int, low, high int) {
	if low < high {

		// find pivot element such that
		// elements smaller than pivot are on the left
		// elements greater than pivot are on the right

		pivot := partition(slice, low, high)
		quickSort(slice, low, pivot-1)
		quickSort(slice, pivot+1, high)
	}
}

func partition(slice []int, low, high int) int {

	// chose the rightmost element as pivot
	pivot := slice[high]
	// pointer for greater element
	i := low - 1

	// traverse through all elements
	// compare each element with pivot
	for j := low; j < high; j++ {
		if slice[j] <= pivot {
			// if element smaller than pivot is found
			// swap it with the greater element pointed by i
			i++
			// swapping element at i with element at j
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	// swap the pivot element with the greater element specified by i
	slice[i+1], slice[high] = slice[high], slice[i+1]

	return i + 1
}
