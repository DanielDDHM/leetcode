package main

import "testing"

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{
			name:   "example1",
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			name:   "example2",
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
		{
			name:   "example3",
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
		{
			name:   "single_number",
			tokens: []string{"42"},
			want:   42,
		},
		{
			name:   "negative_result",
			tokens: []string{"5", "3", "-"},
			want:   2,
		},
		{
			name:   "division_with_negatives",
			tokens: []string{"-13", "5", "/"},
			want:   -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := evalRPN(tt.tokens)
			if got != tt.want {
				t.Errorf("evalRPN(%v) = %d, want %d", tt.tokens, got, tt.want)
			}
		})
	}
}
