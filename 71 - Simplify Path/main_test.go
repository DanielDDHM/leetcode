package main

import "testing"

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "example 1",
			path: "/home/",
			want: "/home",
		},
		{
			name: "example 2",
			path: "/home//foo/",
			want: "/home/foo",
		},
		{
			name: "example 3",
			path: "/home/user/Documents/../Pictures",
			want: "/home/user/Pictures",
		},
		{
			name: "example 4",
			path: "/../",
			want: "/",
		},
		{
			name: "example 5",
			path: "/.../a/../b/c/../d/./",
			want: "/.../b/d",
		},
		{
			name: "root only",
			path: "/",
			want: "/",
		},
		{
			name: "single dot",
			path: "/.",
			want: "/",
		},
		{
			name: "multiple parent refs at root",
			path: "/../../../",
			want: "/",
		},
		{
			name: "dots as directory name",
			path: "/a/...",
			want: "/a/...",
		},
		{
			name: "multiple slashes",
			path: "////home////user",
			want: "/home/user",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := simplifyPath(tt.path)
			if got != tt.want {
				t.Errorf("simplifyPath(%q) = %q, want %q", tt.path, got, tt.want)
			}
		})
	}
}
