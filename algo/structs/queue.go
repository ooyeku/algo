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

// Queue represents a thread-safe queue data structure.
// It uses a mutex to ensure exclusive access to the underlying slice.
// The queue supports the operations Enqueue, Dequeue, Peek, IsEmpty, and Size.
// Enqueue adds an item to the end of the queue.
// Dequeue removes and returns the item from the front of the queue.
// Peek returns the item from the front of the queue without removing it.
// IsEmpty checks if the queue is empty and returns a boolean value.
// Size returns the number of items in the queue.
type Queue struct {
	mu    sync.Mutex
	items []interface{}
}

// NewQueue creates a new instance of the Queue data structure with an empty list of items.
func NewQueue() *Queue {
	return &Queue{
		items: make([]interface{}, 0),
	}
}

// Enqueue adds an item to the queue.
// It acquires a lock on the queue, appends the item to the underlying slice,
// and then releases the lock.
func (q *Queue) Enqueue(item interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

// Dequeue removes and returns the first item from the queue. If the queue is empty,
// it returns nil. The method is thread-safe, and it uses a mutex to synchronize access
// to the queue.
func (q *Queue) Dequeue() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Peek returns the first element in the queue without removing it. If the queue is empty, it returns nil.
func (q *Queue) Peek() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return nil
	}
	return q.items[0]
}

// IsEmpty returns true if the queue is empty, false otherwise. It locks the queue for thread safety.
func (q *Queue) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items) == 0
}

// Size returns the number of elements in the queue.
func (q *Queue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items)
}
