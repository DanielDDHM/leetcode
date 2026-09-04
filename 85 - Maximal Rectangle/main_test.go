package main

import "testing"

func TestMaximalRectangle(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]byte
		expected int
	}{
		{
			name: "Example 1",
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			expected: 6,
		},
		{
			name:     "Example 2 - Single zero",
			matrix:   [][]byte{{'0'}},
			expected: 0,
		},
		{
			name:     "Example 3 - Single one",
			matrix:   [][]byte{{'1'}},
			expected: 1,
		},
		{
			name: "All zeros",
			matrix: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			expected: 0,
		},
		{
			name: "All ones",
			matrix: [][]byte{
				{'1', '1', '1'},
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
			expected: 9,
		},
		{
			name: "Single row",
			matrix: [][]byte{
				{'1', '1', '0', '1', '1'},
			},
			expected: 2,
		},
		{
			name: "Single column",
			matrix: [][]byte{
				{'1'},
				{'1'},
				{'0'},
				{'1'},
			},
			expected: 2,
		},
		{
			name: "Horizontal rectangle",
			matrix: [][]byte{
				{'1', '1', '1', '1'},
				{'0', '0', '0', '0'},
			},
			expected: 4,
		},
		{
			name: "Vertical rectangle",
			matrix: [][]byte{
				{'1', '0'},
				{'1', '0'},
				{'1', '0'},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximalRectangle(tt.matrix)
			if result != tt.expected {
				t.Errorf("maximalRectangle() = %d, want %d", result, tt.expected)
			}
		})
	}
}
