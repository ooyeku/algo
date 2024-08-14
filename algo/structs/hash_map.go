package structs

import "sync"

// HashMap is a thread-safe map implementation in Go. It provides methods for adding, getting, removing, and checking the size and emptiness of items in the map.
// Declaration:
type HashMap struct {
	mu    sync.Mutex
	items map[interface{}]interface{}
}

// NewHashMap returns a new instance of HashMap with an empty item map.
func NewHashMap() *HashMap {
	return &HashMap{
		items: make(map[interface{}]interface{}),
	}
}

// Put adds or updates an item in the HashMap with the specified key and value.
// It locks the HashMap, adds or updates the item, and then unlocks the HashMap.
func (hm *HashMap) Put(key, value interface{}) {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	hm.items[key] = value
}

// Get returns the value associated with the given key in the HashMap.
// If the key is not found, it returns nil.
// The function acquires a lock on the HashMap before accessing the items.
// The lock is released after retrieving the value.
//
// Parameters:
//   - key: The key of the item to retrieve.
//
// Returns:
//   - interface{}: The value associated with the key, or nil if the key is not found.
func (hm *HashMap) Get(key interface{}) interface{} {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	return hm.items[key]
}

// Remove removes the key-value pair with the specified key from the HashMap.
// If the key does not exist in the HashMap, no action is taken.
// This method acquires and releases a lock to ensure thread-safety.
// To avoid potential race conditions, it is recommended to call this method
// within a lock or in a thread-safe manner.
func (hm *HashMap) Remove(key interface{}) {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	delete(hm.items, key)
}

// Size returns the number of items in the HashMap.
// It acquires the lock, retrieves the length of the items map,
// and releases the lock before returning the result.
func (hm *HashMap) Size() int {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	return len(hm.items)
}

// IsEmpty returns a boolean value indicating whether the HashMap is empty or not.
// It acquires a lock on the HashMap, checks if the length of the items in the
// HashMap is zero, and releases the lock. It returns true if the items are empty,
// otherwise, it returns false.
func (hm *HashMap) IsEmpty() bool {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	return len(hm.items) == 0
}
