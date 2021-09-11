package reversegame

import "testing"

func TestSolve(t *testing.T) {
	tt := []struct {
		name string
		n    int
		k    int
		want int
	}{
		{
			name: "example 1",
			n:    3,
			k:    1,
			want: 2,
		},
		{
			name: "example 2",
			n:    5,
			k:    2,
			want: 4,
		},
		{
			name: "even n, even k",
			n:    6,
			k:    2,
			want: 5,
		},
		{
			name: "even n, odd k",
			n:    6,
			k:    3,
			want: 4,
		},
		{
			name: "odd n, even k",
			n:    5,
			k:    2,
			want: 4,
		},
		{
			name: "odd n, odd k",
			n:    7,
			k:    5,
			want: 2,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := solve(tc.n, tc.k); got != tc.want {
				t.Fatalf("solve(%d, %d) = %d ; want %d",
					tc.n, tc.k, got, tc.want)
			}
		})
	}
}
