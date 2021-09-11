// Package reversegame solves this problem:
//
// Akash and Akhil are playing a game. They have balls numbered 0 from to N-1.
// Akhil asks Akash to reverse the position of the balls, i.e., to change the
// order from say, 0,1,2,3 to 3,2,1,0. He further asks Akash to reverse the
// position of the balls N times, each time starting from one position further
// to the right, till he reaches the last ball. So, Akash has to reverse the
// positions of the ball starting from 0th position, then from 1st position,
// then from 2nd position and so on.
//
// At the end of the game, Akhil will ask Akash the final position of any ball
// numbered K. Akash will win the game, if he can answer. Help Akash.
package reversegame

func solve(size, needle int) int {
	seq := sequence(size)
	return indexOf(seq, needle)
}

func sequence(n int) []int {
	seq := make([]int, n)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			seq[i] = n - 1 - (i / 2)
		} else {
			seq[i] = i / 2
		}
	}

	return seq
}

func indexOf(haystack []int, needle int) int {
	for i, n := range haystack {
		if n == needle {
			return i
		}
	}
	return -1
}

func seqSlow(n int) []int {
	seq := make([]int, n)
	for i := 0; i < n; i++ {
		seq[i] = i
	}

	for i := 0; i < n; i++ {
		reverse(seq[i:])
	}
	return seq
}

func reverse(ints []int) {
	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
		ints[i], ints[j] = ints[j], ints[i]
	}
}
