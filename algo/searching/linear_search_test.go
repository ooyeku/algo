package searching

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		name   string
		slice  []int
		target int
		want   int
	}{
		{"Empty slice", []int{}, 5, -1},
		{"Target at beginning", []int{5, 3, 2, 1}, 5, 0},
		{"Target at end", []int{1, 2, 3, 5}, 5, 3},
		{"Target in middle", []int{1, 5, 2}, 5, 1},
		{"Target not present", []int{1, 2, 3}, 5, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinearSearch(tt.slice, tt.target); got != tt.want {
				t.Errorf("LinearSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinearSearchString(t *testing.T) {
	tests := []struct {
		name   string
		slice  []string
		target string
		want   int
	}{
		{"Empty slice", []string{}, "5", -1},
		{"Target at beginning", []string{"5", "3", "2", "1"}, "5", 0},
		{"Target at end", []string{"1", "2", "3", "5"}, "5", 3},
		{"Target in middle", []string{"1", "5", "2"}, "5", 1},
		{"Target not present", []string{"1", "2", "3"}, "5", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinearSearchString(tt.slice, tt.target); got != tt.want {
				t.Errorf("LinearSearchString() = %v, want %v", got, tt.want)
			}
		})
	}
}
