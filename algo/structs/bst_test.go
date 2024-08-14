package structs

import (
	"sync"
	"testing"
)

func TestNewBinarySearchTree(t *testing.T) {
	tests := []struct {
		name string
		less func(a, b interface{}) bool
		want *BinarySearchTree
	}{
		{
			name: "Should Create New BST",
			less: func(a, b interface{}) bool { return a.(int) < b.(int) },
			want: &BinarySearchTree{
				less: func(a, b interface{}) bool { return a.(int) < b.(int) },
				mu:   sync.Mutex{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinarySearchTree(tt.less); got.less(1, 2) != tt.want.less(1, 2) {
				t.Errorf("NewBinarySearchTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_Insert(t *testing.T) {
	bst := NewBinarySearchTree(func(a, b interface{}) bool { return a.(int) < b.(int) })
	tests := []struct {
		name  string
		value interface{}
	}{
		{
			name:  "Insert Value Into BST",
			value: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst.Insert(tt.value)
			if !bst.Search(tt.value) {
				t.Errorf("Expected %v to exist in BST", tt.value)
			}
		})
	}
}

func TestBinarySearchTree_Search(t *testing.T) {
	bst := NewBinarySearchTree(func(a, b interface{}) bool { return a.(int) < b.(int) })
	bst.Insert(5)
	tests := []struct {
		name  string
		value interface{}
		want  bool
	}{
		{
			name:  "Existing Value In BST",
			value: 5,
			want:  true,
		},
		{
			name:  "Non-Existing Value In BST",
			value: 3,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bst.Search(tt.value); got != tt.want {
				t.Errorf("BinarySearchTree.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_InOrder(t *testing.T) {
	bst := NewBinarySearchTree(func(a, b interface{}) bool { return a.(int) < b.(int) })
	for _, v := range []int{5, 3, 7, 2, 4, 6, 8} {
		bst.Insert(v)
	}
	tests := []struct {
		name string
		f    func(value interface{})
	}{
		{
			name: "InOrder Traversal Of BST",
			f: func(value interface{}) {
				_ = value // ignoring the function as implementation will vary based on use-cases
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst.InOrder(tt.f)
		})
	}
}
