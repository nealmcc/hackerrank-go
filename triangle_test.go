package main

import "testing"

/*
 * Complete the 'lowestTriangle' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER trianglebase
 *  2. INTEGER area
 */

func lowestTriangle(base, area int32) int32 {
	height := int32(2 * area / base)
	if base*height == 2*area {
		return height
	}
	return height + 1
}

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
