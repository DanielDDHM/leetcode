package main

import "testing"

func TestIsNumber(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"single digit", "0", true},
		{"simple integer", "2", true},
		{"leading zeros", "0089", true},
		{"decimal with space", " 0.1 ", true},
		{"negative decimal", "-0.1", true},
		{"positive with plus", "+3.14", true},
		{"trailing dot", "4.", true},
		{"leading dot", "-.9", true},
		{"with exponent", "2e10", true},
		{"negative with exponent", "-90E3", true},
		{"positive exponent", "3e+7", true},
		{"negative exponent", "+6e-1", true},
		{"complex scientific", "53.5e93", true},
		{"with spaces", " -90e3   ", true},
		{"leading space", " 6e-1", true},
		{"integer dot", "123.45e-6", true},

		{"letters", "abc", false},
		{"invalid char in middle", "1 a", false},
		{"exponent no digits", "1e", false},
		{"exponent space no digits", " 1e", false},
		{"starts with exponent", "e3", false},
		{"decimal in exponent", " 99e2.5 ", false},
		{"double sign", "--6", false},
		{"mixed signs", "-+3", false},
		{"invalid letter", "95a54e53", false},

		{"only spaces", "   ", false},
		{"empty string", "", false},
		{"only dot", ".", false},
		{"only sign", "+", false},
		{"only exponent", "e", false},
		{"sign then dot", "+.", false},
		{"dot then exponent", ".e5", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.s); got != tt.want {
				t.Errorf("isNumber(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
