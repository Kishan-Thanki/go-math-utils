// lcm.go
package mathutils

import (
	"errors"
	"math"
)

// LCM returns the Least Common Multiple of two integers.
// If both numbers are 0, it returns 0. If the calculation would overflow the integer limit, it returns an error.
func LCM(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, nil
	}

	gcd := GCD(a, b)
	ua := absUint(a)
	ub := absUint(b)
	ugcd := uint(gcd)

	// Avoid overflow: (ua / ugcd) * ub
	temp := ua / ugcd
	if ub > 0 && temp > math.MaxUint/ub {
		return 0, errors.New("lcm overflows int")
	}
	result := temp * ub
	if result > uint(math.MaxInt) {
		return 0, errors.New("lcm overflows int")
	}

	return int(result), nil
}
