package structs

import "sync"

// Node represents a node in a linked list.
// Each node contains a value of type interface{} and a reference to the next node.
type Node struct {
	Value interface{}
	Next  *Node
}

// LinkedList represents a thread-safe singly linked list.
//
// It provides methods to append, prepend, and remove values from the list.
// The list can also be queried for its size and whether it is empty.
type LinkedList struct {
	mu   sync.Mutex
	head *Node
	size int
}

// NewLinkedList creates a new instance of the LinkedList data structure.
// It returns a pointer to the newly created LinkedList.
//
// Example usage:
//
//	ll := NewLinkedList()
//	ll.Append(1)
//	ll.Prepend(2)
//	ll.Remove(1)
//	size := ll.Size()
//	isEmpty := ll.IsEmpty()
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Append adds a new node with the given value to the end of the linked list. If the
// linked list is empty, the new node becomes the head. This method is thread-safe and
// uses a mutex to ensure concurrent access is properly synchronized. The size of the
// linked list is increased by 1 after appending the node.
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

// Prepend adds a new node containing the specified value to the beginning of the linked list.
// The function acquires a lock on the linked list to ensure thread safety using a mutex.
// It creates a new node with the specified value and sets the next pointer of the new node to the current head of the linked list.
// It then updates the head pointer to point to the new node, effectively making it the new head.
// Finally, it increments the size of the linked list by one.
// After the function completes, it releases the lock on the linked list.
func (ll *LinkedList) Prepend(value interface{}) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	newNode := &Node{Value: value, Next: ll.head}
	ll.head = newNode
	ll.size++
}

// Remove removes the first occurrence of the specified value from the linked list.
// If the linked list is empty, the method does nothing.
// If the specified value is the value of the head node, the head node is updated to the next node.
// Otherwise, the method iterates through the linked list until it finds the node with the specified value.
// If such a node is found, it is removed by updating the previous node's Next pointer to skip over it.
// The size of the linked list is decremented by 1 if a node is removed.
// This method is safe to use concurrently by multiple goroutines.
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

// Size returns the current size of the linked list.
// It acquires a lock to ensure thread safety and releases it before returning the size.
func (ll *LinkedList) Size() int {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	return ll.size
}

// IsEmpty returns a boolean value indicating whether the linked list is empty or not.
// It acquires a lock on the linked list, checks the size of the linked list, and returns true if the size is 0.
// It releases the lock before returning.
func (ll *LinkedList) IsEmpty() bool {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	return ll.size == 0
}
