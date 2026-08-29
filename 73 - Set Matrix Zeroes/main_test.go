package main

import (
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "example 1",
			matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			want:   [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			name:   "example 2",
			matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			want:   [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
		{
			name:   "single zero at 0,0",
			matrix: [][]int{{0, 1}, {1, 1}},
			want:   [][]int{{0, 0}, {0, 1}},
		},
		{
			name:   "no zeros",
			matrix: [][]int{{1, 2}, {3, 4}},
			want:   [][]int{{1, 2}, {3, 4}},
		},
		{
			name:   "all zeros",
			matrix: [][]int{{0, 0}, {0, 0}},
			want:   [][]int{{0, 0}, {0, 0}},
		},
		{
			name:   "single element zero",
			matrix: [][]int{{0}},
			want:   [][]int{{0}},
		},
		{
			name:   "single element non-zero",
			matrix: [][]int{{5}},
			want:   [][]int{{5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.matrix)
			if !reflect.DeepEqual(tt.matrix, tt.want) {
				t.Errorf("got %v, want %v", tt.matrix, tt.want)
			}
		})
	}
}
