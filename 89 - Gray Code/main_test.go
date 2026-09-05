package main

import (
	"testing"
)

func TestGrayCode(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "n=1",
			n:    1,
			want: []int{0, 1},
		},
		{
			name: "n=2",
			n:    2,
			want: []int{0, 1, 3, 2},
		},
		{
			name: "n=3",
			n:    3,
			want: []int{0, 1, 3, 2, 6, 7, 5, 4},
		},
		{
			name: "n=4",
			n:    4,
			want: []int{0, 1, 3, 2, 6, 7, 5, 4, 12, 13, 15, 14, 10, 11, 9, 8},
		},
		{
			name: "n=16 max",
			n:    16,
			want: generateExpectedGrayCode(16),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := grayCode(tt.n)
			if len(got) != len(tt.want) {
				t.Errorf("grayCode(%d) length = %d, want %d", tt.n, len(got), len(tt.want))
				return
			}
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("grayCode(%d)[%d] = %d, want %d", tt.n, i, v, tt.want[i])
				}
			}
		})
	}
}

func generateExpectedGrayCode(n int) []int {
	result := make([]int, 0, (1 << n))
	for i := 0; i < (1 << n); i++ {
		result = append(result, i^(i>>1))
	}
	return result
}
