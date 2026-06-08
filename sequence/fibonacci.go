// fibonacci.go
package sequence

import (
	"errors"
	"strconv"
)

// Integer represents all signed and unsigned integer types.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
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

// Fibonacci returns the nth Fibonacci number using an iterative approach.
// If n is negative or if the calculation would overflow the limits of type T, it returns an error.
func Fibonacci[T Integer](n T) (T, error) {
	if n < 0 {
		var zero T
		nStr := strconv.FormatInt(int64(n), 10)
		return zero, errors.New("fibonacci index " + nStr + " is negative")
	}
	if n <= 1 {
		return n, nil
	}
	_, max := getLimits[T]()
	a, b := T(0), T(1)
	for i := T(2); i <= n; i++ {
		if a > max-b {
			var zero T
			nStr := strconv.FormatUint(uint64(n), 10)
			return zero, errors.New("fibonacci of " + nStr + " overflows type")
		}
		a, b = b, a+b
	}
	return b, nil
}
