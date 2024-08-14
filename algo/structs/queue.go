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
type Queue struct {
	mu    sync.Mutex
	items []interface{}
}

// NewQueue creates a new queue.
func NewQueue() *Queue {
	return &Queue{
		items: make([]interface{}, 0),
	}
}

// Enqueue adds an item to the end of the queue.
func (q *Queue) Enqueue(item interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item from the front of the queue.
// It returns nil if the queue is empty.
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

// Peek returns the item at the front of the queue without removing it.
// It returns nil if the queue is empty.
func (q *Queue) Peek() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return nil
	}
	return q.items[0]
}

// IsEmpty checks if the queue is empty.
func (q *Queue) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items) == 0
}

// Size returns the number of items in the queue.
func (q *Queue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items)
}
