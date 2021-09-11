package bestdivisor

func bestDivisor(n int32) int32 {
	best := divisor{1, 1}

	for d := int32(1); d*d <= n; d++ {
		if n%d == 0 {
			a := divisor{d, sumDigits(d)}
			if a.isBetter(best) {
				best = a
			}
			b := divisor{n / d, sumDigits(n / d)}
			if b.isBetter(best) {
				best = b
			}
		}
	}

	return best.value
}

func sumDigits(n int32) int32 {
	var sum int32
	for n > 9 {
		sum += n % 10
		n /= 10
	}
	sum += n
	return sum
}

type divisor struct {
	value     int32
	sumDigits int32
}

// isBetter determines if d is better than d2
func (d divisor) isBetter(d2 divisor) bool {
	if d.sumDigits > d2.sumDigits {
		return true
	}

	if d.sumDigits < d2.sumDigits {
		return false
	}

	return d.value < d2.value
}
