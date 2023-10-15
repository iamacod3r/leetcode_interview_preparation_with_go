package sorting

// https://www.programiz.com/dsa/merge-sort
/*
	Time Complexity : Best  O(n * log_n)	Worst O(n * log_n)	Avarage O(n * log_n)
	Space Complexity : O(n)
*/
/*
	Merge Sort algorithm using Divide and Conquer technique.
	We divide a problem into subproblems to sole the main problem
	https://www.programiz.com/dsa/divide-and-conquer
*/
func Merge(slice []int) {
	mergeSort(slice, 0, len(slice)-1)
}

// Divide the array into two subarrays, sort them and mergeSort them
func mergeSort(slice []int, left, right int) {
	if left < right {

		pivot := (left + right) / 2

		mergeSort(slice, left, pivot)
		mergeSort(slice, pivot+1, right)
		mergeHalves(slice, left, pivot, right)
	}
}

// mergeHalves two subarrays to into array
func mergeHalves(slice []int, left, pivot, right int) {

	// create leftSlice slice[left..pivot] and rightSlice slice[pivot..right]
	leftSize := pivot - left + 1
	rightSize := right - pivot

	leftPart := make([]int, leftSize)
	rightPart := make([]int, rightSize)

	for l := 0; l < leftSize; l++ {
		leftPart[l] = slice[left+l]
	}
	for r := 0; r < rightSize; r++ {
		rightPart[r] = slice[pivot+1+r]
	}

	// Until we reach either end of either L or M, pick larger among
	// elements L and M and place them in the correct position at slice[left..right]
	i, j, k := 0, 0, left

	for i < leftSize && j < rightSize {
		if leftPart[i] <= rightPart[j] {
			slice[k] = leftPart[i]
			i++
		} else {
			slice[k] = rightPart[j]
			j++
		}
		k++
	}

	// When we run out of elements in either leftSlice or rightSlice,
	// pick up the remaining elements and put in slice[left..right]
	for i < leftSize {
		slice[k] = leftPart[i]
		i++
		k++
	}

	for j < rightSize {
		slice[k] = rightPart[j]
		j++
		k++
	}
}
