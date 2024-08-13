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

import "math"

// JumpSearch performs a jump search on a sorted integer slice
func JumpSearch(arr []int, x int) int {
	n := len(arr)
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
