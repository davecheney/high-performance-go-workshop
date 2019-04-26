package asm

import "testing"

func TestAdd(t *testing.T) {
	tests := map[string]struct {
		a, b int64
		want int64
	}{
		"simple": {
			a: 1, b: 3,
			want: 4,
		},
		"subtraction": {
			a: 1, b: -2,
			want: -1,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Add(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("want: %d, got: %d", tc.want, got)
			}
		})
	}
}
