package sorting

// MergeSort sorts the given slice using the mergesort algorithm.
func MergeSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}

	mid := len(slice) / 2
	left := make([]int, mid)
	right := make([]int, len(slice)-mid)

	copy(left, slice[:mid])
	copy(right, slice[mid:])

	MergeSort(left)
	MergeSort(right)

	merge(slice, left, right)

	return slice
}

func merge(slice, left, right []int) {
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		slice[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		slice[k] = right[j]
		j++
		k++
	}
}

// MergeSortString sorts the given slice of strings using the mergesort algorithm.
func MergeSortString(slice []string) []string {
	if len(slice) <= 1 {
		return slice
	}

	mid := len(slice) / 2
	left := make([]string, mid)
	right := make([]string, len(slice)-mid)

	copy(left, slice[:mid])
	copy(right, slice[mid:])

	MergeSortString(left)
	MergeSortString(right)

	mergeString(slice, left, right)

	return slice
}

func mergeString(slice, left, right []string) {
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		slice[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		slice[k] = right[j]
		j++
		k++
	}
}

// MergeSortGeneric sorts the given slice using the mergesort algorithm.
func MergeSortGeneric(slice []interface{}, less func(i, j int) bool) []interface{} {
	if len(slice) <= 1 {
		return slice
	}

	mid := len(slice) / 2
	left := make([]interface{}, mid)
	right := make([]interface{}, len(slice)-mid)

	copy(left, slice[:mid])
	copy(right, slice[mid:])

	MergeSortGeneric(left, less)
	MergeSortGeneric(right, less)

	mergeGeneric(slice, left, right, less)

	return slice
}

func mergeGeneric(slice, left, right []interface{}, less func(i, j int) bool) {
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if less(i, j) {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		slice[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		slice[k] = right[j]
		j++
		k++
	}
}
