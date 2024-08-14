package structs

import "sync"

// Stack represents a thread-safe stack data structure.
type Stack struct {
	mu    sync.Mutex
	items []interface{}
}

// NewStack creates a new stack.
func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

// Push adds an item to the top of the stack.
func (s *Stack) Push(item interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack.
// It returns nil if the stack is empty.
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

// Peek returns the item at the top of the stack without removing it.
// It returns nil if the stack is empty.
func (s *Stack) Peek() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.items) == 0
}

// Size returns the number of items in the stack.
func (s *Stack) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.items)
}
