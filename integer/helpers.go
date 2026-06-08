// helpers.go
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

// getLimits returns the minimum and maximum representable values for any generic Integer type.
func getLimits[T Integer]() (min T, max T) {
	var zero T
	minusOne := zero - 1
	isSigned := minusOne < 0
	if !isSigned {
		return 0, ^T(0)
	}
	temp := T(1)
	for temp > 0 {
		temp <<= 1
	}
	min = temp
	max = ^min
	return min, max
}
