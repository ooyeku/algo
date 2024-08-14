package searching

import "testing"

func TestBinarySearch(t *testing.T) {
	tt := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"One", []int{1, 2, 3, 4, 5}, 1, 0},
		{"Two", []int{1, 2, 3, 4, 5}, 2, 1},
		{"Middle", []int{1, 2, 3, 4, 5}, 3, 2},
		{"Four", []int{1, 2, 3, 4, 5}, 4, 3},
		{"End", []int{1, 2, 3, 4, 5}, 5, 4},
		{"Empty", []int{}, 5, -1},
		{"NotFound", []int{1, 2, 3, 4, 5}, 6, -1},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := BinarySearch(tc.arr, tc.target); got != tc.expected {
				t.Errorf("Expected %d but got %d", tc.expected, got)
			}
		})
	}
}

func TestBinarySearchGeneric(t *testing.T) {
	tt := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"One", []int{1, 2, 3, 4, 5}, 1, 0},
		{"Two", []int{1, 2, 3, 4, 5}, 2, 1},
		{"Middle", []int{1, 2, 3, 4, 5}, 3, 2},
		{"Four", []int{1, 2, 3, 4, 5}, 4, 3},
		{"End", []int{1, 2, 3, 4, 5}, 5, 4},
		{"Empty", []int{}, 5, -1},
		{"NotFound", []int{1, 2, 3, 4, 5}, 6, -1},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			less := func(a, b int) bool {
				return a < b
			}
			if got := BinarySearchGeneric(tc.arr, tc.target, less); got != tc.expected {
				t.Errorf("Expected %d but got %d", tc.expected, got)
			}
		})
	}
}
