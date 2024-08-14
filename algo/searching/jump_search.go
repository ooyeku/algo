package searching

import "math"

// JumpSearch performs a jump search on a sorted integer slice
func JumpSearch(arr []int, x int) int {
	n := len(arr)
	if n == 0 {
		return -1 // Return immediately if the array is empty
	}
	step := int(math.Sqrt(float64(n)))
	prev := 0

	for arr[min(step, n)-1] < x {
		prev = step
		step += int(math.Sqrt(float64(n)))
		if prev >= n {
			return -1
		}
	}

	for arr[prev] < x {
		prev++
		if prev == min(step, n) {
			return -1
		}
	}

	if arr[prev] == x {
		return prev
	}

	return -1
}

// JumpSearchGeneric performs a jump search on a sorted slice of any comparable type
func JumpSearchGeneric[T comparable](arr []T, x T, less func(T, T) bool) int {
	n := len(arr)
	if n == 0 {
		return -1 // Return -1 or appropriate value for empty array
	}
	step := int(math.Sqrt(float64(n)))
	prev := 0

	for less(arr[min(step, n)-1], x) {
		prev = step
		step += int(math.Sqrt(float64(n)))
		if prev >= n {
			return -1
		}
	}

	for less(arr[prev], x) {
		prev++
		if prev == min(step, n) {
			return -1
		}
	}

	if arr[prev] == x {
		return prev
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
