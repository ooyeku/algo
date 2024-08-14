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

func TestBubbleSortGeneric(t *testing.T) {
	tests := []struct {
		input []interface{}
		want  []interface{}
		less  func(i, j int) bool
	}{
		{
			input: []interface{}{5, 2, 1, 3},
			want:  []interface{}{1, 2, 3, 5},
			less:  func(i, j int) bool { return i < j },
		},
		{
			input: []interface{}{"d", "a", "c", "b"},
			want:  []interface{}{"a", "b", "c", "d"},
			less:  func(i, j int) bool { return i < j },
		},
		{
			input: []interface{}{},
			want:  []interface{}{},
			less:  func(i, j int) bool { return i < j },
		},
	}

	for _, tt := range tests {
		got := BubbleSortGeneric(tt.input, tt.less)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("BubbleSortGeneric() = %v, want %v", got, tt.want)
		}
	}
}
