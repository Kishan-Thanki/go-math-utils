// power.go
package integer

import (
	"errors"
)

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

// overflowMul checks if multiplying x and y would result in integer overflow/underflow.
func overflowMul[T Integer](x, y T) bool {
	if x == 0 || y == 0 {
		return false
	}
	min, max := getLimits[T]()
	if x > 0 {
		if y > 0 {
			return x > max/y
		}
		return y < min/x
	}
	if y > 0 {
		return x < min/y
	}
	return y < max/x
}

// Power calculates a raised to the power of b (a^b) using binary exponentiation.
// If b is negative, or if the calculation would overflow the limits of type T, it returns an error.
func Power[T Integer](a, b T) (T, error) {
	if b < 0 {
		var zero T
		return zero, errors.New("power with negative exponent is undefined for integers")
	}
	if b == 0 {
		return 1, nil
	}
	if a == 0 {
		return 0, nil
	}

	var result T = 1
	base := a
	for b > 0 {
		if b%2 == 1 {
			if overflowMul(result, base) {
				var zero T
				return zero, errors.New("power result overflows type")
			}
			result *= base
		}
		b /= 2
		if b > 0 {
			if overflowMul(base, base) {
				var zero T
				return zero, errors.New("power result overflows type")
			}
			base *= base
		}
	}
	return result, nil
}
