package sorting

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "ascending",
			input:    []int{5, 1, 3, 8, 10, 2},
			expected: []int{1, 2, 3, 5, 8, 10},
		},
		{
			name:     "descending",
			input:    []int{10, 8, 6, 4, 2},
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "single",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "empty",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.input); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("MergeSort() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestMergeSortString(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "alphabetical",
			input:    []string{"banana", "apple", "cherry", "date"},
			expected: []string{"apple", "banana", "cherry", "date"},
		},
		{
			name:     "single",
			input:    []string{"only"},
			expected: []string{"only"},
		},
		{
			name:     "empty",
			input:    []string{},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSortString(tt.input); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("MergeSortString() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
