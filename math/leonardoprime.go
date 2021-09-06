package math

import (
	"fmt"
)

var primes []int64

func init() {
	// the product of the first 20 primes overflows int64, so is more than enough.
	primes = sieve(20)
	fmt.Println("primes", primes)
}

// sieve returns the first n primes.
// Uses the sieve of Eratosthenes.
func sieve(n int) []int64 {
	primes := make([]int64, 0, n)
	primes = append(primes, 2)
	// isComposite is indexed by indexOf(odd) where odd is an odd integer > 2
	isComposite := make([]bool, n*n)
	p, done := nextPrime(isComposite, 2)
	for !done {
		primes = append(primes, p)
		markMultiples(isComposite, p)
		p, done = nextPrime(isComposite, p)
		done = done || len(primes) >= n
	}

	return primes
}

func nextPrime(isComposite []bool, after int64) (p int64, done bool) {
	for i := indexOf(after) + 1; i < len(isComposite); i++ {
		if !isComposite[i] {
			return primeOf(i), false
		}
	}
	return 0, true
}

func markMultiples(isComposite []bool, p int64) {
	var (
		mult int64 = p * 3
		i    int   = indexOf(mult)
	)
	for i < len(isComposite) {
		isComposite[i] = true
		mult += 2 * p
		i = indexOf(mult)
	}
}

// indexOf returns the index within isComposite for a potential prime number.
func indexOf(p int64) int {
	return int((p-1)/2 - 1)
}

// primeOf returns the prime number for the given index within isComposite.
func primeOf(i int) int64 {
	return int64(2*(i+1) + 1)
}

// primeCount returns the largest number of distinct prime factors for any
// integer in the range [1, n].
func primeCount(n int64) int32 {
	for i := 0; i < len(primes); i++ {
		prod := foldLeft(func(a, b int64) int64 { return a * b }, 1, primes[:i])
		switch {
		case prod == n:
			return int32(i)
		case prod < 0 || prod > n:
			return int32(i - 1)
		}
	}
	// impossible
	return -1
}

type binfunc func(int64, int64) int64

func foldLeft(fn binfunc, accum int64, slice []int64) int64 {
	for _, n := range slice {
		accum = fn(accum, n)
	}
	return accum
}
