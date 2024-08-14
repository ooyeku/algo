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

// HashMap represents a thread-safe hash map data structure.
type HashMap struct {
	mu    sync.Mutex
	items map[interface{}]interface{}
}

// NewHashMap creates a new hash map.
func NewHashMap() *HashMap {
	return &HashMap{
		items: make(map[interface{}]interface{}),
	}
}

// Put adds a key-value pair to the hash map.
func (hm *HashMap) Put(key, value interface{}) {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	hm.items[key] = value
}

// Get retrieves the value associated with the specified key.
// It returns nil if the key does not exist.
func (hm *HashMap) Get(key interface{}) interface{} {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	return hm.items[key]
}

// Remove deletes the key-value pair associated with the specified key.
func (hm *HashMap) Remove(key interface{}) {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	delete(hm.items, key)
}

// Size returns the number of items in the hash map.
func (hm *HashMap) Size() int {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	return len(hm.items)
}

// IsEmpty checks if the hash map is empty.
func (hm *HashMap) IsEmpty() bool {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	return len(hm.items) == 0
}
