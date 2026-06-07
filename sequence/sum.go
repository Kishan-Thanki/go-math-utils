// sum.go
package sequence

import (
	"errors"
)

// SumNaturalNumbers returns the sum of the first n natural numbers (1 + 2 + ... + n).
// If n is negative or if the calculation would overflow the limits of type T, it returns an error.
func SumNaturalNumbers[T Integer](n T) (T, error) {
	if n < 0 {
		var zero T
		return zero, errors.New("sum of natural numbers is only defined for non-negative integers")
	}
	if n == 0 {
		var zero T
		return zero, nil
	}

	_, max := getLimits[T]()
	var term1, term2 T
	if n%2 == 0 {
		term1 = n / 2
		if n > max-1 {
			var zero T
			return zero, errors.New("sum of natural numbers overflows type")
		}
		term2 = n + 1
	} else {
		term1 = n
		term2 = (n-1)/2 + 1
	}

	// Check if term1 * term2 overflows max
	if term1 > 0 && term2 > max/term1 {
		var zero T
		return zero, errors.New("sum of natural numbers overflows type")
	}
	return term1 * term2, nil
}
