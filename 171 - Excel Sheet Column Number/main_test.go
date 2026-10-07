package main

import "testing"

func TestTitleToNumber(t *testing.T) {
	tests := []struct {
		name   string
		column string
		want   int
	}{
		{"A", "A", 1},
		{"B", "B", 2},
		{"Z", "Z", 26},
		{"AA", "AA", 27},
		{"AB", "AB", 28},
		{"ZY", "ZY", 701},
		{"FXSHRXW", "FXSHRXW", 2147483647},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := titleToNumber(tt.column)
			if got != tt.want {
				t.Errorf("titleToNumber(%q) = %d, want %d", tt.column, got, tt.want)
			}
		})
	}
}
