package structs

import (
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("NewQueue", func(t *testing.T) {
		q := NewQueue()
		if q == nil {
			t.Errorf("NewQueue should never return nil")
		}
		if !q.IsEmpty() {
			t.Errorf("NewQueue should return an empty Queue")
		}
	})

	t.Run("Enqueue", func(t *testing.T) {
		q := NewQueue()
		q.Enqueue("test")
		if q.IsEmpty() {
			t.Errorf("Queue should not be empty after an Enqueue operation")
		}
		if q.Size() != 1 {
			t.Errorf("Queue Size should be 1 after one Enqueue operation, got: %d", q.Size())
		}
	})

	t.Run("Dequeue", func(t *testing.T) {
		q := NewQueue()
		q.Enqueue("test")
		item := q.Dequeue()
		if item != "test" {
			t.Errorf("Dequeue should return the first element Enqueued, got: %v", item)
		}
		if !q.IsEmpty() {
			t.Errorf("Queue should be empty after a Dequeue operation on a single-item Queue")
		}
	})

	t.Run("Peek", func(t *testing.T) {
		q := NewQueue()
		q.Enqueue("test")
		item := q.Peek()
		if item != "test" {
			t.Errorf("Peek should return the first element Enqueued without removing it, got: %v", item)
		}
		if q.IsEmpty() {
			t.Errorf("Queue should not be empty after a Peek operation")
		}
	})

	t.Run("IsEmpty", func(t *testing.T) {
		q := NewQueue()
		if !q.IsEmpty() {
			t.Errorf("Queue should be empty after initialization, got: %v", q.IsEmpty())
		}
		q.Enqueue("test")
		if q.IsEmpty() {
			t.Errorf("Queue should not be empty after Enqueue operation")
		}
	})

	t.Run("Size", func(t *testing.T) {
		q := NewQueue()
		if q.Size() != 0 {
			t.Errorf("Size should be 0 after initialization, got: %d", q.Size())
		}
		q.Enqueue("test")
		if q.Size() != 1 {
			t.Errorf("Size should be 1 after adding one item, got: %d", q.Size())
		}
	})
}
