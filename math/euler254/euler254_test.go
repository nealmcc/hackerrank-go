package euler254

import "testing"

func TestSolve(t *testing.T) {
	tt := []struct {
		name string
		n, m int
		want int
	}{}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := solve(tc.n, tc.m); got != tc.want {
				t.Fatalf("solve(%d, %d = %d ; want %d",
					tc.n, tc.m, got, tc.want)
			}
		})
	}
}
