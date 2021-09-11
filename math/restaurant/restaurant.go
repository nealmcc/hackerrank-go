package restaurant

// restaurant determines how many sections a rectangule of size a x b
// can be divided into, given the following constraints:
//
// Each piece must be the same size as the others.
// Each piece must be a square
// The width of each square must be as large as possible.
// None of the original rectangle can be left over.
func restaurant(a, b int32) int32 {
	width := gcd(a, b)
	return a / width * b / width
}

// gcd returns the greatest common divisor of a and b
func gcd(a, b int32) int32 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
