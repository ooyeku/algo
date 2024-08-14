package structs

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("NewStack", func(t *testing.T) {
		stack := NewStack()
		if stack.Size() != 0 {
			t.Errorf("New stack should be empty")
		}
	})

	t.Run("Push and Pop", func(t *testing.T) {
		stack := NewStack()
		stack.Push(1)

		if stack.Size() != 1 {
			t.Errorf("Stack size should be 1 after push")
		}

		item := stack.Pop()
		if item != 1 || stack.Size() != 0 {
			t.Errorf("Pop did not return the right item or changed the stack size")
		}
	})

	t.Run("Pop empty", func(t *testing.T) {
		stack := NewStack()
		item := stack.Pop()
		if item != nil {
			t.Errorf("Pop on empty stack should return nil")
		}
	})

	t.Run("Peek", func(t *testing.T) {
		stack := NewStack()
		stack.Push(1)
		item := stack.Peek()

		if item != 1 || stack.Size() != 1 {
			t.Errorf("Peek should return the top item without changing stack size")
		}
	})

	t.Run("Peek empty", func(t *testing.T) {
		stack := NewStack()
		item := stack.Peek()
		if item != nil {
			t.Errorf("Peek on empty stack should return nil")
		}
	})

	t.Run("IsEmpty", func(t *testing.T) {
		stack := NewStack()
		if !stack.IsEmpty() {
			t.Errorf("New stack should be empty")
		}

		stack.Push(1)
		if stack.IsEmpty() {
			t.Errorf("Stack should not be empty after push")
		}

		stack.Pop()
		if !stack.IsEmpty() {
			t.Errorf("Stack should be empty after pop")
		}
	})

	t.Run("Size", func(t *testing.T) {
		stack := NewStack()
		if stack.Size() != 0 {
			t.Errorf("New stack should have size 0")
		}

		stack.Push(1)
		if stack.Size() != 1 {
			t.Errorf("Push should increase size")
		}

		stack.Pop()
		if stack.Size() != 0 {
			t.Errorf("Pop should decrease size")
		}
	})
}
