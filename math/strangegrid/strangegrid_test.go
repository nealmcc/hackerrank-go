package strangegrid

import "testing"

func TestSolve(t *testing.T) {
	tt := []struct {
		name string
		r    int
		c    int
		want int
	}{
		{
			name: "example 1",
			r:    6,
			c:    3,
			want: 25,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := solve(tc.r, tc.c); got != tc.want {
				t.Fatalf("solve(%d, %d) = %d ; want %d",
					tc.r, tc.c, got, tc.want)
			}
		})
	}
}
