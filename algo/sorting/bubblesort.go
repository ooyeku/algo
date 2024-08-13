package sorting

// BubbleSort sorts the given slice using the bubblesort algorithm.
func BubbleSort(slice []int) []int {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}

// BubbleSortString sorts the given slice of strings using the bubblesort algorithm.
func BubbleSortString(slice []string) []string {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}

// BubbleSortGeneric sorts the given slice using the bubblesort algorithm.
func BubbleSortGeneric(slice []interface{}, less func(i, j int) bool) []interface{} {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if less(j+1, j) {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}
