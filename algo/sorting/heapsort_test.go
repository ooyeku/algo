package sorting

import (
	"fmt"
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

func TestHeapSortGeneric(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []interface{}
		less     func(i, j int) bool
		expected []interface{}
	}{
		{
			desc:     "Empty interface slice",
			input:    []interface{}{},
			less:     func(i, j int) bool { return i < j },
			expected: []interface{}{},
		},
		{
			desc:     "Single interface",
			input:    []interface{}{1},
			less:     func(i, j int) bool { return i < j },
			expected: []interface{}{1},
		},
		{
			desc:     "Two interfaces",
			input:    []interface{}{2, 1},
			less:     func(i, j int) bool { return fmt.Sprintf("%v", i) < fmt.Sprintf("%v", j) },
			expected: []interface{}{1, 2},
		},
		{
			desc:     "Multiple interfaces",
			input:    []interface{}{5, 3, "A", 4, "B", 9, 2},
			less:     func(i, j int) bool { return fmt.Sprintf("%v", i) < fmt.Sprintf("%v", j) },
			expected: []interface{}{"A", "B", 2, 3, 4, 5, 9},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := HeapSortGeneric(tC.input, tC.less)
			if !reflect.DeepEqual(result, tC.expected) {
				t.Errorf("HeapSortGeneric(): got %v, want %v", result, tC.expected)
			}
		})
	}
}
