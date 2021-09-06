package findpoint

/*
 * Complete the 'findPoint' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER px
 *  2. INTEGER py
 *  3. INTEGER qx
 *  4. INTEGER qy
 */

// findPoint returns the reflection point (x, y) that is the result
// of rotating point p about point q by 180 degrees.
func findPoint(px int32, py int32, qx int32, qy int32) []int32 {
    x := 2*qx - px
    y := 2*qy - py
    return []int32{x, y}
}
