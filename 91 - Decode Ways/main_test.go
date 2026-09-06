package main

import "testing"

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "12",
			want: 2,
		},
		{
			name: "Example 2",
			s:    "226",
			want: 3,
		},
		{
			name: "Example 3",
			s:    "06",
			want: 0,
		},
		{
			name: "Single digit 1-9",
			s:    "5",
			want: 1,
		},
		{
			name: "Leading zero",
			s:    "0",
			want: 0,
		},
		{
			name: "String with zero in middle",
			s:    "101",
			want: 1,
		},
		{
			name: "All ones",
			s:    "111",
			want: 3,
		},
		{
			name: "Longer valid string",
			s:    "1123",
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.s); got != tt.want {
				t.Errorf("numDecodings(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
