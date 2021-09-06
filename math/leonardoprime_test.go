package math

import (
	"math"
	"testing"
)

func TestPrimeCount(t *testing.T) {
	tt := []struct {
		name string
		n    int64
		want int32
	}{
		{"1 has 0 prime factors", 1, 0},
		{"2 has 1 prime factor", 2, 1},
		{"3 has 1 prime factor", 3, 1},
		{"210 has 4 prime factors", 210, 4},
		{"3,569,485,920 has 9 prime factors", 3569485920, 9},
		{"55,327,031,760 has 10 prime factors", 55327031760, 10},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := primeCount(tc.n); got != tc.want {
				t.Fatalf("primeCount(%d) = %d ; want %d",
					tc.n, got, tc.want)
			}
		})
	}
}

func TestCountPrimeFactors(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		n    int64
		want int32
	}{
		{"1 has no prime factors", 1, 0},
		{"2 has one prime factor", 2, 1},
		{"3 has one prime factor", 3, 1},
		{"210 has one four prime factors", 210, 4},
		{"500 has a number less than it with 4 prime factors", 500, 4},
		{"5,000 -> 5", 5000, 5},
		{"614889782588491410 -> 15", 614889782588491410, 15},
		{"614889782588491409 -> 14", 614889782588491409, 14},
		{"614889782588491408 -> 14", 614889782588491408, 14},
		{"614889782588491407 -> 14", 614889782588491407, 14},
		{"6148897825884914 -> 13", 6148897825884914, 13},
		{"614889782588 -> 11", 614889782588, 11},
		{"max int64 -> 15", math.MaxInt64, 15},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := primeCount(tc.n); got != tc.want {
				t.Fatalf("primeCount(%d) = %d ; want %d",
					tc.n, got, tc.want)
			}
		})
	}
}

func TestIndexOf(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		in   int64
		want int
	}{
		{"3 => 0", 3, 0},
		{"5 => 1", 5, 1},
		{"7 => 2", 7, 2},
		{"9 => 3", 9, 3},
		{"11 => 4", 11, 4},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := indexOf(tc.in); got != tc.want {
				t.Fatalf("indexOf(%d) = %d ; want %d",
					tc.in, got, tc.want)
			}
		})
	}
}

func TestPrimeOf(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		want int64
		in   int
	}{
		{"3 => 0", 3, 0},
		{"5 => 1", 5, 1},
		{"7 => 2", 7, 2},
		{"9 => 3", 9, 3},
		{"11 => 4", 11, 4},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := primeOf(tc.in); got != tc.want {
				t.Fatalf("primeOf(%d) = %d ; want %d",
					tc.in, got, tc.want)
			}
		})
	}
}
