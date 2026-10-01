package main

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "example1",
			s:    "the sky is blue",
			want: "blue is sky the",
		},
		{
			name: "example2",
			s:    "  hello world  ",
			want: "world hello",
		},
		{
			name: "example3",
			s:    "a good   example",
			want: "example good a",
		},
		{
			name: "single_word",
			s:    "hello",
			want: "hello",
		},
		{
			name: "multiple_spaces",
			s:    "  a  b  c  ",
			want: "c b a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseWords(tt.s)
			if got != tt.want {
				t.Errorf("reverseWords(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
