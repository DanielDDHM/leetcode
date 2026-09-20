package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "example 1",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "example 2",
			s:    "race a car",
			want: false,
		},
		{
			name: "example 3",
			s:    " ",
			want: true,
		},
		{
			name: "empty string",
			s:    "",
			want: true,
		},
		{
			name: "single character",
			s:    "a",
			want: true,
		},
		{
			name: "numeric palindrome",
			s:    "12321",
			want: true,
		},
		{
			name: "mixed case",
			s:    "AaBbAa",
			want: true,
		},
		{
			name: "not numeric palindrome",
			s:    "12345",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
