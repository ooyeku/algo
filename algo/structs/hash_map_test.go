package structs

import (
	"testing"
)

func TestHashMap(t *testing.T) {
	tests := []struct {
		name     string
		ops      []func(*HashMap) interface{}
		expected []interface{}
	}{
		{
			name: "AddItemGetItem",
			ops: []func(*HashMap) interface{}{
				func(hm *HashMap) interface{} {
					hm.Put("key1", "value1")
					return nil
				},
				func(hm *HashMap) interface{} {
					return hm.Get("key1")
				},
			},
			expected: []interface{}{nil, "value1"},
		},
		{
			name: "AddMultipleGetMultiple",
			ops: []func(*HashMap) interface{}{
				func(hm *HashMap) interface{} {
					hm.Put("key1", "value1")
					return nil
				},
				func(hm *HashMap) interface{} {
					hm.Put("key2", "value2")
					return nil
				},
				func(hm *HashMap) interface{} {
					return hm.Get("key1")
				},
				func(hm *HashMap) interface{} {
					return hm.Get("key2")
				},
			},
			expected: []interface{}{nil, nil, "value1", "value2"},
		},
		{
			name: "RemoveItem",
			ops: []func(*HashMap) interface{}{
				func(hm *HashMap) interface{} {
					hm.Put("key1", "value1")
					return nil
				},
				func(hm *HashMap) interface{} {
					hm.Remove("key1")
					return nil
				},
				func(hm *HashMap) interface{} {
					return hm.Get("key1")
				},
			},
			expected: []interface{}{nil, nil, nil},
		},
		{
			name: "CheckSize",
			ops: []func(*HashMap) interface{}{
				func(hm *HashMap) interface{} {
					hm.Put("key1", "value1")
					return nil
				},
				func(hm *HashMap) interface{} {
					hm.Put("key2", "value2")
					return nil
				},
				func(hm *HashMap) interface{} {
					return hm.Size()
				},
			},
			expected: []interface{}{nil, nil, 2},
		},
		{
			name: "CheckEmpty",
			ops: []func(*HashMap) interface{}{
				func(hm *HashMap) interface{} {
					return hm.IsEmpty()
				},
				func(hm *HashMap) interface{} {
					hm.Put("key1", "value1")
					return nil
				},
				func(hm *HashMap) interface{} {
					return hm.IsEmpty()
				},
				func(hm *HashMap) interface{} {
					hm.Remove("key1")
					return nil
				},
				func(hm *HashMap) interface{} {
					return hm.IsEmpty()
				},
			},
			expected: []interface{}{true, nil, false, nil, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := NewHashMap()
			for i, op := range tt.ops {
				result := op(hm)
				if result != tt.expected[i] {
					t.Errorf("Expected %v, got %v", tt.expected[i], result)
				}
			}
		})
	}
}
