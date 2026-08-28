package main

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "example 1",
			n:    2,
			want: 2,
		},
		{
			name: "example 2",
			n:    3,
			want: 3,
		},
		{
			name: "edge case: n=1",
			n:    1,
			want: 1,
		},
		{
			name: "edge case: n=4",
			n:    4,
			want: 5,
		},
		{
			name: "larger input",
			n:    10,
			want: 89,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.n)
			if got != tt.want {
				t.Errorf("climbStairs(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
