package math

import (
	"testing"
)

func TestGameWithCells(t *testing.T) {
	tt := []struct {
		name string
		n, m int32
		want int32
	}{
		{
			name: "1x1 (base case)",
			n:    1,
			m:    1,
			want: 1,
		},
		{
			name: "1x2 (base case)",
			n:    1,
			m:    2,
			want: 1,
		},
		{
			name: "2x1 (base case)",
			n:    2,
			m:    1,
			want: 1,
		},
		{
			name: "2x2 (base case)",
			n:    2,
			m:    2,
			want: 1,
		},
		{
			name: "4x4 (both even)",
			n:    4,
			m:    4,
			want: 4,
		},
		{
			name: "2x3 (m odd)",
			n:    2,
			m:    3,
			want: 2,
		},
		{
			name: "3x2 (n odd)",
			n:    3,
			m:    2,
			want: 2,
		},
		{
			name: "5x5 (both odd)",
			n:    5,
			m:    5,
			want: 9,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := gameWithCells(tc.n, tc.m); got != tc.want {
				t.Fatalf("gameWithCells(%d, %d) = %d ; want %d",
					tc.n, tc.m, got, tc.want)
			}
		})
	}
}
