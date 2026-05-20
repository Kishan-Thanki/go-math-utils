// power.go
package mathutils

import (
	"errors"
	"math"
)

// overflowMul checks if multiplying x and y would result in integer overflow/underflow.
func overflowMul(x, y int) bool {
	if x == 0 || y == 0 {
		return false
	}
	if x > 0 {
		if y > 0 {
			return x > math.MaxInt/y
		}
		return y < math.MinInt/x
	}
	if y > 0 {
		return x < math.MinInt/y
	}
	return y < math.MaxInt/x
}

// Power calculates a raised to the power of b (a^b) using binary exponentiation.
// If b is negative, or if the calculation would overflow the integer limit, it returns an error.
func Power(a, b int) (int, error) {
	if b < 0 {
		return 0, errors.New("power with negative exponent is undefined for integers")
	}
	if b == 0 {
		return 1, nil
	}
	if a == 0 {
		return 0, nil
	}

	result := 1
	base := a
	for b > 0 {
		if b%2 == 1 {
			if overflowMul(result, base) {
				return 0, errors.New("power result overflows int")
			}
			result *= base
		}
		b /= 2
		if b > 0 {
			if overflowMul(base, base) {
				return 0, errors.New("power result overflows int")
			}
			base *= base
		}
	}
	return result, nil
}
