package searching

// LinearSearch performs a linear search on an integer slice
func LinearSearch(slice []int, target int) int {
	for i, value := range slice {
		if value == target {
			return i
		}
	}
	return -1 // Return -1 if the target is not found
}

// LinearSearchString performs a linear search on a string slice
func LinearSearchString(slice []string, target string) int {
	for i, value := range slice {
		if value == target {
			return i
		}
	}
	return -1 // Return -1 if the target is not found
}

// LinearSearchGeneric performs a linear search on a generic slice
func LinearSearchGeneric[T comparable](slice []T, target T) int {
	for i, value := range slice {
		if value == target {
			return i
		}
	}
	return -1 // Return -1 if the target is not found
}
