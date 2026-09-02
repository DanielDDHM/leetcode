package main

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{"example1", []int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{"example2", []int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{"single_element_found", []int{1}, 1, true},
		{"single_element_not_found", []int{1}, 2, false},
		{"all_same_found", []int{1, 1, 1, 1, 1, 1, 1}, 1, true},
		{"all_same_not_found", []int{1, 1, 1, 1, 1, 1, 1}, 2, false},
		{"two_elements_found_first", []int{1, 3}, 1, true},
		{"two_elements_found_second", []int{1, 3}, 3, true},
		{"duplicates_at_pivot", []int{3, 1, 1}, 3, true},
		{"target_at_start", []int{4, 5, 6, 7, 0, 1, 2}, 4, true},
		{"target_at_end", []int{4, 5, 6, 7, 0, 1, 2}, 2, true},
		{"not_rotated", []int{1, 2, 3, 4, 5}, 3, true},
		{"empty_like_duplicates", []int{1, 3, 1, 1, 1}, 3, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("search(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
