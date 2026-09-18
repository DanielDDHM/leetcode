package main

import (
	"slices"
	"testing"
)

func TestGetRow(t *testing.T) {
	tests := []struct {
		name     string
		rowIndex int
		want     []int
	}{
		{
			name:     "Row 0",
			rowIndex: 0,
			want:     []int{1},
		},
		{
			name:     "Row 1",
			rowIndex: 1,
			want:     []int{1, 1},
		},
		{
			name:     "Row 2",
			rowIndex: 2,
			want:     []int{1, 2, 1},
		},
		{
			name:     "Row 3",
			rowIndex: 3,
			want:     []int{1, 3, 3, 1},
		},
		{
			name:     "Row 4",
			rowIndex: 4,
			want:     []int{1, 4, 6, 4, 1},
		},
		{
			name:     "Row 5",
			rowIndex: 5,
			want:     []int{1, 5, 10, 10, 5, 1},
		},
		{
			name:     "Row 33 (max constraint)",
			rowIndex: 33,
			want:     []int{1, 33, 528, 5456, 40920, 237336, 1107568, 4272048, 13884156, 38567100, 92561040, 193536720, 354817320, 573166440, 818809200, 1037158320, 1166803110, 1166803110, 1037158320, 818809200, 573166440, 354817320, 193536720, 92561040, 38567100, 13884156, 4272048, 1107568, 237336, 40920, 5456, 528, 33, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getRow(tt.rowIndex)
			if !slices.Equal(got, tt.want) {
				t.Errorf("getRow(%d) = %v, want %v", tt.rowIndex, got, tt.want)
			}
		})
	}
}
