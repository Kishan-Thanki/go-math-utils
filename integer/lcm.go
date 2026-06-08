// lcm.go
package integer

import (
	"errors"
	"math"
)

// LCM returns the Least Common Multiple of two generic integers.
// If both numbers are 0, it returns 0. If the calculation would overflow the limits of type T, it returns an error.
func LCM[T Integer](a, b T) (T, error) {
	if a == 0 || b == 0 {
		var zero T
		return zero, nil
	}

	gcd := GCD(a, b)
	ua := absUint64(a)
	ub := absUint64(b)
	ugcd := uint64(gcd)

	// Avoid overflow: (ua / ugcd) * ub
	temp := ua / ugcd
	if ub > 0 && temp > math.MaxUint64/ub {
		var zero T
		return zero, errors.New("lcm overflows type")
	}
	result := temp * ub
	// Check for overflow
	if uint64(T(result)) != result || T(result) < 0 {
		var zero T
		return zero, errors.New("lcm overflows type")
	}

	return T(result), nil
}
