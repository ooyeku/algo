package searching

import (
	"testing"
)

func TestJumpSearch(t *testing.T) {
	tt := []struct {
		name     string
		arr      []int
		x        int
		expected int
	}{
		{"Empty", []int{}, 1, -1},
		{"Single", []int{1}, 1, 0},
		{"NotFound", []int{1, 2, 3, 4, 5}, 6, -1},
		{"Multiple", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6, 5},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := JumpSearch(tc.arr, tc.x)
			if got != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, got)
			}
		})
	}
}

func TestJumpSearchGeneric(t *testing.T) {
	tt := []struct {
		name     string
		arr      []int
		x        int
		expected int
	}{
		{"Empty", []int{}, 1, -1},
		{"Single", []int{1}, 1, 0},
		{"NotFound", []int{1, 2, 3, 4, 5}, 6, -1},
		{"Multiple", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6, 5},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			less := func(a, b int) bool { return a < b }
			got := JumpSearchGeneric(tc.arr, tc.x, less)
			if got != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, got)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tt := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Equal", 1, 1, 1},
		{"FirstLess", -2, 3, -2},
		{"SecondLess", 4, 3, 3},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := min(tc.a, tc.b)
			if got != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, got)
			}
		})
	}
}
