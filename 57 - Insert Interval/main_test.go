package main

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		expected    [][]int
	}{
		{
			name:        "example 1",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			expected:    [][]int{{1, 5}, {6, 9}},
		},
		{
			name:        "example 2",
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			expected:    [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name:        "empty intervals",
			intervals:   [][]int{},
			newInterval: []int{1, 5},
			expected:    [][]int{{1, 5}},
		},
		{
			name:        "new interval before all",
			intervals:   [][]int{{5, 7}, {8, 10}},
			newInterval: []int{1, 2},
			expected:    [][]int{{1, 2}, {5, 7}, {8, 10}},
		},
		{
			name:        "new interval after all",
			intervals:   [][]int{{1, 2}, {3, 5}},
			newInterval: []int{6, 8},
			expected:    [][]int{{1, 2}, {3, 5}, {6, 8}},
		},
		{
			name:        "new interval merges all",
			intervals:   [][]int{{1, 5}, {6, 9}},
			newInterval: []int{2, 10},
			expected:    [][]int{{1, 10}},
		},
		{
			name:        "new interval inside existing",
			intervals:   [][]int{{1, 10}},
			newInterval: []int{2, 5},
			expected:    [][]int{{1, 10}},
		},
		{
			name:        "single interval exact overlap",
			intervals:   [][]int{{1, 5}},
			newInterval: []int{1, 5},
			expected:    [][]int{{1, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := insert(tt.intervals, tt.newInterval)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
