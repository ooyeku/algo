package sorting

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []int
		expected []int
	}{
		{
			desc:     "Empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			desc:     "Single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			desc:     "Two elements",
			input:    []int{2, 1},
			expected: []int{1, 2},
		},
		{
			desc:     "Multiple elements",
			input:    []int{5, 3, 8, 4, 1, 9, 2},
			expected: []int{1, 2, 3, 4, 5, 8, 9},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := HeapSort(tC.input)
			if !reflect.DeepEqual(result, tC.expected) {
				t.Errorf("HeapSort(): got %v, want %v", result, tC.expected)
			}
		})
	}
}

func TestHeapSortString(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []string
		expected []string
	}{
		{
			desc:     "Empty string slice",
			input:    []string{},
			expected: []string{},
		},
		{
			desc:     "Single string",
			input:    []string{"A"},
			expected: []string{"A"},
		},
		{
			desc:     "Two strings",
			input:    []string{"B", "A"},
			expected: []string{"A", "B"},
		},
		{
			desc:     "Multiple strings",
			input:    []string{"Dog", "Cat", "Cow", "Horse"},
			expected: []string{"Cat", "Cow", "Dog", "Horse"},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := HeapSortString(tC.input)
			if !reflect.DeepEqual(result, tC.expected) {
				t.Errorf("HeapSortString(): got %v, want %v", result, tC.expected)
			}
		})
	}
}
