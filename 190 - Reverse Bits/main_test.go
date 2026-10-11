package main

import "testing"

func TestReverseBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "example 1",
			n:    43261596,
			want: 964176192,
		},
		{
			name: "example 2",
			n:    2147483644,
			want: 1073741822,
		},
		{
			name: "zero",
			n:    0,
			want: 0,
		},
		{
			name: "one",
			n:    1,
			want: 2147483648,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBits(tt.n)
			if got != tt.want {
				t.Errorf("reverseBits(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
