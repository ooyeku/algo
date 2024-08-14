package structs

import "sync"

// Stack is a thread-safe data structure that implements the stack concept.
// It allows adding, removing and retrieving items in a Last-In-First-Out (LIFO) manner.
type Stack struct {
	mu    sync.Mutex
	items []interface{}
}

// NewStack returns a new Stack instance.
// The Stack is initialized with an empty slice of interface{}.
func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

// Push adds an item to the top of the stack. After pushing an item,
// the stack size is increased by one.
//
// It is safe to call this method concurrently from multiple goroutines,
// as it uses a mutex to provide synchronization.
func (s *Stack) Push(item interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack. If the stack is empty, it returns nil.
// It acquires a lock to ensure synchronization and releases it before returning the item.
func (s *Stack) Pop() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Peek returns the element at the top of the stack without removing it. If the stack is empty, it returns nil.
func (s *Stack) Peek() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

// IsEmpty returns true if the stack is empty, false otherwise.
// This method acquires a lock on the stack, checks the length of the items slice,
// and returns true if it is 0, indicating an empty stack.
// It releases the lock before returning the result.
func (s *Stack) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.items) == 0
}

// Size returns the number of items in the stack. It acquires a lock to ensure
// thread safety and releases it before returning the result.
func (s *Stack) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.items)
}
