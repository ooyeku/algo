package sorting

import (
	"math/rand"
	"testing"
)

func TestIntroSort(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want []int
	}{
		{"empty", []int{}, []int{}},
		{"single", []int{5}, []int{5}},
		{"unordered", []int{5, 2, 7, 3, 4}, []int{2, 3, 4, 5, 7}},
		{"reversed", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"duplicates", []int{5, 2, 2, 5, 5}, []int{2, 2, 5, 5, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntroSort(tt.data)
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("expected %v, got %v", tt.want, got)
				}
			}
		})
	}
}

func TestIntroSortGeneric(t *testing.T) {
	tests := []struct {
		name string
		data []interface{}
		want []interface{}
	}{
		{"empty", []interface{}{}, []interface{}{}},
		{"single", []interface{}{5}, []interface{}{5}},
		{"unordered", []interface{}{5, 2, 7, 3, 4}, []interface{}{2, 3, 4, 5, 7}},
		{"reversed", []interface{}{5, 4, 3, 2, 1}, []interface{}{1, 2, 3, 4, 5}},
		{"duplicates", []interface{}{5, 2, 2, 5, 5}, []interface{}{2, 2, 5, 5, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntroSortGeneric(tt.data, func(i, j int) bool { return tt.data[i].(int) < tt.data[j].(int) })
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("expected %v, got %v", tt.want, got)
				}
			}
		})
	}
}

func TestIntroSortString(t *testing.T) {
	tests := []struct {
		name string
		data []string
		want []string
	}{
		{"empty", []string{}, []string{}},
		{"single", []string{"a"}, []string{"a"}},
		{"unordered", []string{"e", "b", "g", "c", "d"}, []string{"b", "c", "d", "e", "g"}},
		{"reversed", []string{"e", "d", "c", "b", "a"}, []string{"a", "b", "c", "d", "e"}},
		{"duplicates", []string{"e", "b", "b", "e", "e"}, []string{"b", "b", "e", "e", "e"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntroSortString(tt.data)
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("expected %v, got %v", tt.want, got)
				}
			}
		})
	}
}

func BenchmarkIntroSort(b *testing.B) {
	data := make([]int, b.N)
	for i := range data {
		data[i] = rand.Int()
	}
	b.ResetTimer()
	IntroSort(data)
}
