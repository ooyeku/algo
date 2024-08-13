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
