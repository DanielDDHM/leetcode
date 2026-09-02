package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		wantK    int
		wantNums []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 1, 2, 2, 3},
			wantK:    5,
			wantNums: []int{1, 1, 2, 2, 3},
		},
		{
			name:     "Example 2",
			nums:     []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			wantK:    7,
			wantNums: []int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			name:     "Single element",
			nums:     []int{1},
			wantK:    1,
			wantNums: []int{1},
		},
		{
			name:     "All same elements",
			nums:     []int{1, 1, 1},
			wantK:    2,
			wantNums: []int{1, 1},
		},
		{
			name:     "No duplicates",
			nums:     []int{1, 2, 3, 4},
			wantK:    4,
			wantNums: []int{1, 2, 3, 4},
		},
		{
			name:     "Multiple groups with extras",
			nums:     []int{1, 1, 1, 2, 2, 2, 3, 3, 3},
			wantK:    6,
			wantNums: []int{1, 1, 2, 2, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK := removeDuplicates(tt.nums)
			if gotK != tt.wantK {
				t.Errorf("removeDuplicates() returned k=%d, want %d", gotK, tt.wantK)
			}
			if !reflect.DeepEqual(tt.nums[:gotK], tt.wantNums) {
				t.Errorf("removeDuplicates() modified nums to %v, want %v", tt.nums[:gotK], tt.wantNums)
			}
		})
	}
}
