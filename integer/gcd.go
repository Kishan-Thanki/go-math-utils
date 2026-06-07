// gcd.go
package integer

import "math"

// absUint returns the absolute value of an int as a uint to prevent overflow.
func absUint(val int) uint {
	if val == math.MinInt {
		return uint(math.MaxInt) + 1
	}
	if val < 0 {
		return uint(-val)
	}
	return uint(val)
}

// GCD returns the Greatest Common Divisor of two integers using the Euclidean algorithm.
// The result is always non-negative. If the result overflows the maximum signed integer
// limit (e.g. GCD(math.MinInt, 0)), it panics.
func GCD(a, b int) int {
	ua := absUint(a)
	ub := absUint(b)

	for ub != 0 {
		ua, ub = ub, ua%ub
	}
	if ua > math.MaxInt {
		panic("gcd result overflows int")
	}
	return int(ua)
}
