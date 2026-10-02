package main

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		values     []int
		expected   []interface{}
	}{
		{
			name:       "official example",
			operations: []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"},
			values:     []int{1, 2, 0},
			expected:   []interface{}{nil, nil, nil, nil, 0, nil, 2, 1},
		},
		{
			name:       "single element",
			operations: []string{"MinStack", "push", "getMin", "top", "pop"},
			values:     []int{5},
			expected:   []interface{}{nil, nil, 5, 5, nil},
		},
		{
			name:       "negative numbers",
			operations: []string{"MinStack", "push", "push", "push", "getMin", "pop", "getMin"},
			values:     []int{-1, -2, -3},
			expected:   []interface{}{nil, nil, nil, nil, -3, nil, -2},
		},
		{
			name:       "identical minimums",
			operations: []string{"MinStack", "push", "push", "push", "getMin", "pop", "getMin", "pop", "getMin"},
			values:     []int{1, 1, 1},
			expected:   []interface{}{nil, nil, nil, nil, 1, nil, 1, nil, 1},
		},
		{
			name:       "min at different positions",
			operations: []string{"MinStack", "push", "push", "getMin", "push", "getMin", "pop", "getMin"},
			values:     []int{2147483647, -2147483648, 0},
			expected:   []interface{}{nil, nil, nil, -2147483648, nil, -2147483648, nil, -2147483648},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minStack := Constructor()
			valueIdx := 0
			expectedIdx := 0

			for _, op := range tt.operations {
				switch op {
				case "MinStack":
					if tt.expected[expectedIdx] != nil {
						t.Errorf("expected nil for MinStack construction, got %v", tt.expected[expectedIdx])
					}
					expectedIdx++
				case "push":
					minStack.Push(tt.values[valueIdx])
					valueIdx++
					if tt.expected[expectedIdx] != nil {
						t.Errorf("expected nil for push, got %v", tt.expected[expectedIdx])
					}
					expectedIdx++
				case "pop":
					minStack.Pop()
					if tt.expected[expectedIdx] != nil {
						t.Errorf("expected nil for pop, got %v", tt.expected[expectedIdx])
					}
					expectedIdx++
				case "top":
					result := minStack.Top()
					expected := tt.expected[expectedIdx].(int)
					if result != expected {
						t.Errorf("top: expected %d, got %d", expected, result)
					}
					expectedIdx++
				case "getMin":
					result := minStack.GetMin()
					expected := tt.expected[expectedIdx].(int)
					if result != expected {
						t.Errorf("getMin: expected %d, got %d", expected, result)
					}
					expectedIdx++
				}
			}
		})
	}
}
