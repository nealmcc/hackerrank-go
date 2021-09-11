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

func TestSequence(t *testing.T) {
	tt := []struct {
		name string
		in   int
		want []int
	}{
		{"size 1", 1, seqSlow(1)},
		{"size 2", 2, seqSlow(2)},
		{"size 3", 3, seqSlow(3)},
		{"size 4", 4, seqSlow(4)},
		{"size 5", 5, seqSlow(5)},
		{"size 6", 6, seqSlow(6)},
		{"size 7", 7, seqSlow(7)},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := sequence(tc.in)
			if len(got) != len(tc.want) {
				t.Fatalf("sequence(%d) = length %d ; want %d",
					tc.in, len(got), len(tc.want))
			}

			for i, want := range tc.want {
				if got[i] != want {
					t.Fatalf("\n\t\tsequence(%d) = %v\n\t\t\t want %v", tc.in, got, tc.want)
				}
			}
		})
	}
}
