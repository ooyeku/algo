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

package sorting

// HeapSort sorts the given slice using the heapsort algorithm.
func HeapSort(slice []int) []int {
	n := len(slice)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(slice, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		slice[0], slice[i] = slice[i], slice[0]
		heapify(slice, i, 0)
	}
	return slice
}

func heapify(slice []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && slice[left] > slice[largest] {
		largest = left
	}
	if right < n && slice[right] > slice[largest] {
		largest = right
	}
	if largest != i {
		slice[i], slice[largest] = slice[largest], slice[i]
		heapify(slice, n, largest)
	}
}

// HeapSortString sorts the given slice of strings using the heapsort algorithm.
func HeapSortString(slice []string) []string {
	n := len(slice)
	for i := n/2 - 1; i >= 0; i-- {
		heapifyString(slice, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		slice[0], slice[i] = slice[i], slice[0]
		heapifyString(slice, i, 0)
	}
	return slice
}

func heapifyString(slice []string, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && slice[left] > slice[largest] {
		largest = left
	}
	if right < n && slice[right] > slice[largest] {
		largest = right
	}
	if largest != i {
		slice[i], slice[largest] = slice[largest], slice[i]
		heapifyString(slice, n, largest)
	}
}

// HeapSortGeneric sorts the given slice using the heapsort algorithm.
func HeapSortGeneric(slice []interface{}, less func(i, j int) bool) []interface{} {
	n := len(slice)
	for i := n/2 - 1; i >= 0; i-- {
		heapifyGeneric(slice, n, i, less)
	}
	for i := n - 1; i >= 0; i-- {
		slice[0], slice[i] = slice[i], slice[0]
		heapifyGeneric(slice, i, 0, less)
	}
	return slice
}

func heapifyGeneric(slice []interface{}, n, i int, less func(i, j int) bool) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && less(largest, left) {
		largest = left
	}
	if right < n && less(largest, right) {
		largest = right
	}
	if largest != i {
		slice[i], slice[largest] = slice[largest], slice[i]
		heapifyGeneric(slice, n, largest, less)
	}
}
