package sorting

import (
	"math"
	"sort"
)

func IntroSort(slice []int) []int {
	maxDepth := 2 * int(math.Floor(math.Log2(float64(len(slice)))))
	introSortRec(slice, 0, len(slice)-1, maxDepth)
	return slice
}

func introSortRec(slice []int, start, end, maxDepth int) {
	if end-start < 16 {
		insertionSort(slice[start : end+1])
	} else if maxDepth == 0 {
		HeapSort(slice[start : end+1])
	} else {
		p := partitionIntro(slice, start, end)
		introSortRec(slice, start, p-1, maxDepth-1)
		introSortRec(slice, p+1, end, maxDepth-1)
	}
}

func insertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1
		for j >= 0 && slice[j] > key {
			slice[j+1] = slice[j]
			j--
		}
		slice[j+1] = key
	}
}

func partitionIntro(slice []int, low, high int) int {
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

func IntroSortGeneric(slice []interface{}, less func(i, j int) bool) []interface{} {
	sort.Slice(slice, less)
	return slice
}

func IntroSortString(slice []string) []string {
	sort.Strings(slice)
	return slice
}
