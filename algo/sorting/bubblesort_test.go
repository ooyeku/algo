package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{5, 2, 3, 1}, []int{1, 2, 3, 5}},
		{[]int{5, 1, 1, 2, 0, 0}, []int{0, 0, 1, 1, 2, 5}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		got := BubbleSort(tt.input)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
		}
	}
}

func TestBubbleSortString(t *testing.T) {
	tests := []struct {
		input []string
		want  []string
	}{
		{[]string{"d", "a", "c", "b"}, []string{"a", "b", "c", "d"}},
		{[]string{"a", "c", "", "b"}, []string{"", "a", "b", "c"}},
		{[]string{}, []string{}},
	}

	for _, tt := range tests {
		got := BubbleSortString(tt.input)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("BubbleSortString() = %v, want %v", got, tt.want)
		}
	}
}
