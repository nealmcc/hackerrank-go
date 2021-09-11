package bestdivisor

import (
	"testing"
)

func TestBestDivisor_FromInput(t *testing.T) {
	tt := []struct {
		name string
		in   int32
		want int32
	}{
		{
			name: "12 -> 1 * 6 = 6",
			in:   12,
			want: 6,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := bestDivisor(tc.in); got != tc.want {
				t.Fatalf("bestDivisor(%d) = %d ; want %d",
					tc.in, got, tc.want)
			}
		})
	}
}
