package main

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	tests := []struct {
		name       string
		capacity   int
		operations []struct {
			op    string
			key   int
			value int
		}
		expected []int
	}{
		{
			name:     "Example 1",
			capacity: 2,
			operations: []struct {
				op    string
				key   int
				value int
			}{
				{op: "put", key: 1, value: 1},
				{op: "put", key: 2, value: 2},
				{op: "get", key: 1},
				{op: "put", key: 3, value: 3},
				{op: "get", key: 2},
				{op: "put", key: 4, value: 4},
				{op: "get", key: 1},
				{op: "get", key: 3},
				{op: "get", key: 4},
			},
			expected: []int{1, -1, -1, 3, 4},
		},
		{
			name:     "Example 2",
			capacity: 2,
			operations: []struct {
				op    string
				key   int
				value int
			}{
				{op: "put", key: 1, value: 1},
				{op: "put", key: 2, value: 2},
				{op: "get", key: 1},
				{op: "put", key: 3, value: 3},
				{op: "get", key: 2},
				{op: "put", key: 4, value: 4},
				{op: "get", key: 1},
				{op: "get", key: 3},
				{op: "get", key: 4},
			},
			expected: []int{1, -1, -1, 3, 4},
		},
		{
			name:     "Single capacity",
			capacity: 1,
			operations: []struct {
				op    string
				key   int
				value int
			}{
				{op: "put", key: 1, value: 1},
				{op: "get", key: 1},
				{op: "put", key: 2, value: 2},
				{op: "get", key: 1},
				{op: "get", key: 2},
			},
			expected: []int{1, -1, 2},
		},
		{
			name:     "Update existing key",
			capacity: 2,
			operations: []struct {
				op    string
				key   int
				value int
			}{
				{op: "put", key: 1, value: 1},
				{op: "put", key: 1, value: 2},
				{op: "get", key: 1},
				{op: "put", key: 2, value: 2},
				{op: "put", key: 1, value: 3},
				{op: "get", key: 2},
			},
			expected: []int{2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lru := Constructor(tt.capacity)
			resultIdx := 0

			for _, op := range tt.operations {
				if op.op == "get" {
					result := lru.Get(op.key)
					if resultIdx < len(tt.expected) && result != tt.expected[resultIdx] {
						t.Errorf("Get(%d) = %d, expected %d", op.key, result, tt.expected[resultIdx])
					}
					resultIdx++
				} else if op.op == "put" {
					lru.Put(op.key, op.value)
				}
			}
		})
	}
}
