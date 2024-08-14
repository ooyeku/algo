// Copyright 2024 olayeku
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package structs

import "sync"

// Node represents a node in a linked list.
type Node struct {
	Value interface{}
	Next  *Node
}

// LinkedList represents a thread-safe linked list data structure.
type LinkedList struct {
	mu   sync.Mutex
	head *Node
	size int
}

// NewLinkedList creates a new linked list.
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Append adds an item to the end of the linked list.
func (ll *LinkedList) Append(value interface{}) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	newNode := &Node{Value: value}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	ll.size++
}

// Prepend adds an item to the beginning of the linked list.
func (ll *LinkedList) Prepend(value interface{}) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	newNode := &Node{Value: value, Next: ll.head}
	ll.head = newNode
	ll.size++
}

// Remove removes the first occurrence of the specified value from the linked list.
func (ll *LinkedList) Remove(value interface{}) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	if ll.head == nil {
		return
	}
	if ll.head.Value == value {
		ll.head = ll.head.Next
		ll.size--
		return
	}
	current := ll.head
	for current.Next != nil && current.Next.Value != value {
		current = current.Next
	}
	if current.Next != nil {
		current.Next = current.Next.Next
		ll.size--
	}
}

// Size returns the number of items in the linked list.
func (ll *LinkedList) Size() int {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	return ll.size
}

// IsEmpty checks if the linked list is empty.
func (ll *LinkedList) IsEmpty() bool {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	return ll.size == 0
}
