package main

import (
	"testing"
)

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "Example 1",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			name: "Example 2",
			s:    "a",
			t:    "a",
			want: "a",
		},
		{
			name: "Example 3",
			s:    "a",
			t:    "aa",
			want: "",
		},
		{
			name: "No match",
			s:    "abc",
			t:    "xyz",
			want: "",
		},
		{
			name: "T longer than S",
			s:    "a",
			t:    "aaa",
			want: "",
		},
		{
			name: "Entire string is window",
			s:    "abc",
			t:    "abc",
			want: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minWindow(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("minWindow(%q, %q) = %q, want %q", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
