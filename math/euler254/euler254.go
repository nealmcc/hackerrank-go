// Package euler254 solves Problem 254 from projecteuler.net as described on
// hackerrank.com:
//
// https://www.hackerrank.com/contests/projecteuler/challenges/euler254/problem
// https://projecteuler.net/problem=254
package euler254

// f is the sum of the factorials of the digits of n.
//
// Ex:
//  f(342) = 3! + 4! + 2! = 32
func f(n int) int {
	return 0
}

// sf is the sum of the digits of f(n).
//
// Ex:
//  sf(342) = 3 + 2 + 5
func sf(n int) int {
	return 0
}

// g is the smallest positive integer n such that sf(n) = i
//
// Though sf(342) is 5, sf(25) is also 5.
// ex:
//  g(5) = 25
//  g(20) = 267
func g(i int) (n int) {
	return 0
}

// sg is the sum of the digits of g(i).
//
// ex:
//  sg(5) = 2 + 5 = 7
//  sg(20) = 2 + 6 + 7 = 15
func sg(i int) int {
	return 0
}

// solve finds the sum from i = 1 to i = n of sg(i), module m.
func solve(n, m int) int {
	return 0
}
