package structs

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	t.Run("LinkedList operations", func(t *testing.T) {
		tests := []struct {
			name       string
			operations func(ll *LinkedList)
			isEmpty    bool
			size       int
		}{
			{
				name:       "Initialization",
				operations: func(ll *LinkedList) {},
				isEmpty:    true,
				size:       0,
			},
			{
				name: "Single append operation",
				operations: func(ll *LinkedList) {
					ll.Append(1)
				},
				isEmpty: false,
				size:    1,
			},
			{
				name: "Multiple append operations",
				operations: func(ll *LinkedList) {
					ll.Append(1)
					ll.Append(2)
					ll.Append(3)
				},
				isEmpty: false,
				size:    3,
			},
			{
				name: "Single prepend operation",
				operations: func(ll *LinkedList) {
					ll.Prepend(1)
				},
				isEmpty: false,
				size:    1,
			},
			{
				name: "Multiple prepend operations",
				operations: func(ll *LinkedList) {
					ll.Prepend(1)
					ll.Prepend(2)
					ll.Prepend(3)
				},
				isEmpty: false,
				size:    3,
			},
			{
				name: "Append and prepend operations",
				operations: func(ll *LinkedList) {
					ll.Prepend(1)
					ll.Append(2)
					ll.Prepend(3)
				},
				isEmpty: false,
				size:    3,
			},
			{
				name: "Single remove operation",
				operations: func(ll *LinkedList) {
					ll.Append(1)
					ll.Remove(1)
				},
				isEmpty: true,
				size:    0,
			},
			{
				name: "Multiple remove operations",
				operations: func(ll *LinkedList) {
					ll.Append(1)
					ll.Append(2)
					ll.Append(3)
					ll.Remove(1)
					ll.Remove(2)
					ll.Remove(3)
				},
				isEmpty: true,
				size:    0,
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				ll := NewLinkedList()
				tc.operations(ll)
				if ll.IsEmpty() != tc.isEmpty {
					t.Errorf("Expected isEmpty to be %v, but got %v", tc.isEmpty, ll.IsEmpty())
				}
				if ll.Size() != tc.size {
					t.Errorf("Expected size to be %v, but got %v", tc.size, ll.Size())
				}
			})

		}
	})
}
