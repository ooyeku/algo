package structs

import (
	"sort"
	"testing"
)

func lessInt(a, b interface{}) bool {
	return a.(int) < b.(int)
}

func TestRedBlackTree(t *testing.T) {
	cases := []struct {
		name      string
		insertInt []int
		searchInt []int
		found     []bool
		inOrder   []int
	}{
		{
			name:      "EmptyTree",
			insertInt: []int{},
			searchInt: []int{1, 2, 3},
			found:     []bool{false, false, false},
			inOrder:   []int{},
		},
		{
			name:      "SingleElement",
			insertInt: []int{10},
			searchInt: []int{10, 20, 30},
			found:     []bool{true, false, false},
			inOrder:   []int{10},
		},
		{
			name:      "MultipleElements",
			insertInt: []int{50, 30, 70, 20, 40, 25},
			searchInt: []int{20, 25, 70, 100},
			found:     []bool{true, true, true, false},
			inOrder:   []int{20, 25, 30, 40, 50, 70},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rbt := NewRedBlackTree(lessInt)
			for _, val := range c.insertInt {
				rbt.Insert(val)
			}
			for idx, val := range c.searchInt {
				res := rbt.Search(val)
				if res != c.found[idx] {
					t.Errorf("Expected search result %v, but got %v", c.found[idx], res)
				}
			}
			resultInOrder := make([]int, 0)
			rbt.InOrder(func(value interface{}) {
				resultInOrder = append(resultInOrder, value.(int))
			})
			if !sort.IntsAreSorted(resultInOrder) {
				t.Errorf("Values are not sorted in InOrder traversal")
			}
			for idx, val := range c.inOrder {
				if resultInOrder[idx] != val {
					t.Errorf("Expected value on InOrder traversal %v, but got %v", val, resultInOrder[idx])
				}
			}
		})
	}
}
