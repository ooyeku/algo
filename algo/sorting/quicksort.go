package sorting

// QuickSort sorts the given slice using the quicksort algorithm.
func QuickSort(slice []int) []int {
	slice = quickSort(slice, 0, len(slice)-1)
	return slice
}

func quickSort(slice []int, low, high int) []int {
	if low < high {
		pi := partition(slice, low, high)
		slice := quickSort(slice, low, pi-1)
		slice = quickSort(slice, pi+1, high)
	}
	return slice
}

func partition(slice []int, low, high int) int {
	pivot := slice[high]
	i := low - 1
	for j := low; j < high; j++ {
		if slice[j] < pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

// QuickSortString sorts the given slice of strings using the quicksort algorithm.
func QuickSortString(slice []string) []string {
	slice = quickSortString(slice, 0, len(slice)-1)
	return slice
}

func quickSortString(slice []string, low, high int) []string {
	if low < high {
		pi := partitionString(slice, low, high)
		slice := quickSortString(slice, low, pi-1)
		slice = quickSortString(slice, pi+1, high)
	}
	return slice
}

func partitionString(slice []string, low, high int) int {
	pivot := slice[high]
	i := low - 1
	for j := low; j < high; j++ {
		if slice[j] < pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

// QuickSortGeneric sorts the given slice using the quicksort algorithm.
func QuickSortGeneric(slice []interface{}, less func(i, j int) bool) []interface{} {
	slice = quickSortGeneric(slice, 0, len(slice)-1, less)
	return slice
}

func quickSortGeneric(slice []interface{}, low, high int, less func(i, j int) bool) []interface{} {
	if low < high {
		pi := partitionGeneric(slice, low, high, less)
		slice := quickSortGeneric(slice, low, pi-1, less)
		slice = quickSortGeneric(slice, pi+1, high, less)
	}
	return slice
}

func partitionGeneric(slice []interface{}, low, high int, less func(i, j int) bool) int {
	_ = slice[high]
	i := low - 1
	for j := low; j < high; j++ {
		if less(j, high) {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}
