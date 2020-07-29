package insertionSort

import "sort_algorithm/sequence"

func insertionSort(arr []int) sequence.Sequence {
	for i := range arr {
		preIndex := i - 1
		current := arr[i]

		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex + 1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex + 1] = current
	}

	return arr
}
