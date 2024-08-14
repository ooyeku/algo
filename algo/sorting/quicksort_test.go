package sorting

import (
	"fmt"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{3, 2, 1},
			want:  []int{1, 2, 3},
		},
		{
			input: []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6},
			want:  []int{2, 3, 5, 6, 7, 9, 10, 11, 12, 14},
		},
		{
			input: []int{},
			want:  []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {
			got := QuickSort(tc.input)
			if !sort.IntsAreSorted(got) || len(got) != len(tc.want) {
				t.Errorf("QuickSort() = %v; want %v", got, tc.want)
			}
		})
	}
}

func TestQuickSortString(t *testing.T) {
	testCases := []struct {
		input []string
		want  []string
	}{
		{
			input: []string{"apple", "banana", "cherry"},
			want:  []string{"apple", "banana", "cherry"},
		},
		{
			input: []string{"violet", "indigo", "blue", "green", "yellow", "orange", "red"},
			want:  []string{"blue", "green", "indigo", "orange", "red", "violet", "yellow"},
		},
		{
			input: []string{},
			want:  []string{},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {
			got := QuickSortString(tc.input)
			if !sort.StringsAreSorted(got) || len(got) != len(tc.want) {
				t.Errorf("QuickSortString() = %v; want %v", got, tc.want)
			}
		})
	}
}

func TestQuickSortGeneric(t *testing.T) {
	testCases := []struct {
		input []interface{}
		less  func(i, j int) bool
		want  []interface{}
	}{
		{
			input: []interface{}{3.1, 2.1, 1.1},
			less:  func(i, j int) bool { return i < j },
			want:  []interface{}{1.1, 2.1, 3.1},
		},
		{
			input: []interface{}{9.1, 7.1, 5.1, 11.1, 12.1, 2.1, 14.1, 3.1, 10.1, 6.1},
			less:  func(i, j int) bool { return i < j },
			want:  []interface{}{2.1, 3.1, 5.1, 6.1, 7.1, 9.1, 10.1, 11.1, 12.1, 14.1},
		},
		{
			input: []interface{}{},
			less:  func(i, j int) bool { return i < j },
			want:  []interface{}{},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {
			got := QuickSortGeneric(tc.input, tc.less)
			if !sort.SliceIsSorted(got, tc.less) || len(got) != len(tc.want) {
				t.Errorf("QuickSortGeneric() = %v; want %v", got, tc.want)
			}
		})
	}
}
