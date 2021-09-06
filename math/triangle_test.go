package math

import "testing"

func TestLowestTriangle(t *testing.T) {
	tt := []struct {
		name string
		area int32
		base int32
		want int32
	}{
		{
			name: "example 1",
			area: 6,
			base: 4,
			want: 3,
		},
		{
			name: "example 2",
			area: 2,
			base: 2,
			want: 2,
		},
		{
			name: "example 3",
			area: 100,
			base: 17,
			want: 12,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := lowestTriangle(tc.base, tc.area); got != tc.want {
				t.Fatalf("lowestTriangle(%d, %d) => %d ; want %d",
					tc.base, tc.area, got, tc.want)
			}
		})
	}
}
