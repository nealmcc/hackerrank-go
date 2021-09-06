package handshake

// handshake calculates the number of distinct pairs of people there are in
// a group of n people.  (n choose 2).
func handshake(n int32) int32 {
	if n < 2 {
		return 0
	}

	// using Pascal's triangle:
	// Choose(n, k) = Choose(n, k-1) * (n + 1 - k) / k
	// Choose(n, 0) = 1
	//
	// so:
	// Choose(n, 1) = Choose(n, 0) * (n + 1 - 1) / 1 = n
	// Choose(n, 2) = n * (n + 1 -2) / 2             = n * (n -1)/2

	return n * (n - 1) / 2
}
