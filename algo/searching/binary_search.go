// Copyright 2024 olayeku
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package searching

// BinarySearch performs a binary search on a sorted slice of integers.
// It returns the index of the target if found, or -1 if not found.
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// BinarySearchGeneric performs a binary search on a sorted slice of any comparable type.
// It returns the index of the target if found, or -1 if not found.
func BinarySearchGeneric[T any](arr []T, target T, less func(T, T) bool) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if !less(arr[mid], target) && !less(target, arr[mid]) {
			return mid
		} else if less(arr[mid], target) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
