package main

import "testing"

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		s3   string
		want bool
	}{
		{
			name: "example1",
			s1:   "aabcc",
			s2:   "dbbca",
			s3:   "aadbbcbcac",
			want: true,
		},
		{
			name: "example2",
			s1:   "aabcc",
			s2:   "dbbca",
			s3:   "aadbbbaccc",
			want: false,
		},
		{
			name: "example3",
			s1:   "",
			s2:   "",
			s3:   "",
			want: true,
		},
		{
			name: "s1_empty",
			s1:   "",
			s2:   "ab",
			s3:   "ab",
			want: true,
		},
		{
			name: "s2_empty",
			s1:   "ab",
			s2:   "",
			s3:   "ab",
			want: true,
		},
		{
			name: "s3_wrong_length",
			s1:   "a",
			s2:   "b",
			s3:   "abc",
			want: false,
		},
		{
			name: "single_char_no_interleave",
			s1:   "a",
			s2:   "b",
			s3:   "ba",
			want: true,
		},
		{
			name: "impossible_order",
			s1:   "ab",
			s2:   "cd",
			s3:   "adcb",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isInterleave(tt.s1, tt.s2, tt.s3)
			if got != tt.want {
				t.Errorf("isInterleave(%q, %q, %q) = %v, want %v",
					tt.s1, tt.s2, tt.s3, got, tt.want)
			}
		})
	}
}
