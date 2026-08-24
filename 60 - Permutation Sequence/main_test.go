package main

import "testing"

func TestGetPermutation(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want string
	}{
		{"Example 1", 3, 3, "213"},
		{"Example 2", 4, 9, "2314"},
		{"Example 3", 3, 1, "123"},
		{"First permutation n=1", 1, 1, "1"},
		{"Single element", 2, 1, "12"},
		{"Single element last", 2, 2, "21"},
		{"n=4 k=1", 4, 1, "1234"},
		{"n=4 k=24", 4, 24, "4321"},
		{"n=3 k=6", 3, 6, "321"},
		{"n=5 k=20", 5, 20, "15243"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getPermutation(tt.n, tt.k)
			if got != tt.want {
				t.Errorf("getPermutation(%d, %d) = %q, want %q", tt.n, tt.k, got, tt.want)
			}
		})
	}
}
