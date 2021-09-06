package main

/*
 * Complete the 'gameWithCells' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER m
 */

func gameWithCells(n, m int32) int32 {
	// let nq = n - n%2  // nq is even
	// let nr = n%2      // nr is 0 or 1
	// let mq = m - m%2  // mq is even
	// let mr = m%2      // mr is 0 or 1

	// therefore:
	// n = nq + nr
	// m = mq + mr
	//
	// area = n * m
	//      = (nq + nr)(mq + mr)
	//      = nq*mq + nq*mr + mq*nr + nr*mr
	//      = area1 + area2 + area3 + area4
	//
	// in area1, each drop covers 4 cells
	// in areas 2 and 3, each drop covers 2 cells
	// in area 4, each drop covers 1 cell

	var (
		nq, nr int32 = n - n%2, n % 2
		mq, mr       = m - m%2, m % 2
		area1        = nq * mq
		area2        = mr * nq
		area3        = nr * mq
		area4        = nr * mr
	)

	return area1/4 + area2/2 + area3/2 + area4
}
