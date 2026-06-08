// gcd.go
package integer

// GCD returns the Greatest Common Divisor of two integers using the Euclidean algorithm.
// The result is always non-negative. If the result overflows the maximum positive value
// of the type T (e.g. GCD(math.MinInt, 0)), it panics.
func GCD[T Integer](a, b T) T {
	ua := absUint64(a)
	ub := absUint64(b)

	for ub != 0 {
		ua, ub = ub, ua%ub
	}
	// Check for overflow: if ua overflows the storage of T, casting it to T and back to uint64
	// yields a different value, or T(ua) becomes negative (for signed types).
	if uint64(T(ua)) != ua || T(ua) < 0 {
		panic("gcd result overflows type")
	}
	return T(ua)
}
