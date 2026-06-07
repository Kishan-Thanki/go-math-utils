// numeric.go
package numeric

import "cmp"

// Signed represents all signed integer types.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Float represents all floating-point types.
type Float interface {
	~float32 | ~float64
}

// SignedOrFloat represents types that are either signed integers or floats.
type SignedOrFloat interface {
	Signed | Float
}

// Integer represents all signed and unsigned integer types.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Min returns the smaller of a and b.
func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the larger of a and b.
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Clamp restricts a value to the range [min, max].
// It panics if min > max.
func Clamp[T cmp.Ordered](val, min, max T) T {
	if min > max {
		panic("clamp: min cannot be greater than max")
	}
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}

// Abs returns the absolute value of x.
// Note: For signed integers, Abs(math.MinInt) will overflow and return math.MinInt due to two's complement representation.
func Abs[T SignedOrFloat](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Sign returns -1 if x is negative, 0 if x is zero, and 1 if x is positive.
func Sign[T SignedOrFloat](x T) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

// IsEven returns true if x is even.
func IsEven[T Integer](x T) bool {
	return x%2 == 0
}

// IsOdd returns true if x is odd.
func IsOdd[T Integer](x T) bool {
	return x%2 != 0
}
