// gcd.go
package integer

import "math"

// Integer represents all signed and unsigned integer types.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// absUint64 returns the absolute value of any generic integer as a uint64.
func absUint64[T Integer](val T) uint64 {
	if val < 0 {
		// Cast val to int64 first. This avoids negative wrap-around issues 
		// when converting smaller signed types (like int8(-128)) directly to uint64.
		v64 := int64(val)
		if v64 == math.MinInt64 {
			return uint64(math.MaxInt64) + 1
		}
		if v64 < 0 {
			return uint64(-v64)
		}
		return uint64(v64)
	}
	return uint64(val)
}

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
