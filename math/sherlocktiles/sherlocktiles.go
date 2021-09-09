package sherlocktiles

import (
	"math"
)

/*
 * Complete the 'movingTiles' function below.
 *
 * The function is expected to return a DOUBLE_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER l - the length of the sides of each square
 *  2. INTEGER s1 - the velocity of square 1, along the line y=x
 *  3. INTEGER s2 - the velocity of square 2, along the line y=x
 *  4. INTEGER_ARRAY queries - a list of
 */

// movingTiles calculates, for each area q in queries:
// the time t in seconds, when the area of the overlapping portions ofc
// s1 and s2 will equal q.
func movingTiles(l, v1, v2 int, queries []int64) []float64 {
	// t calculates the time at which the area of the overlapping section is q
	t := func(q int64) float64 {
		width := math.Sqrt(float64(q))
		dist := float64(l) - width
		seconds := dist * math.Sqrt(2) / math.Abs(float64(v1-v2))
		// runtime.Breakpoint()
		return seconds
	}

	times := make([]float64, 0, 8)
	for _, q := range queries {
		times = append(times, t(q))
	}
	return times
}
