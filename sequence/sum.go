// sum.go
package sequence

import (
	"errors"
	"math"
)

// SumNaturalNumbers returns the sum of the first n natural numbers (1 + 2 + ... + n).
// If n is negative or if the calculation would overflow the integer limit, it returns an error.
func SumNaturalNumbers(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("sum of natural numbers is only defined for non-negative integers")
	}
	if n == 0 {
		return 0, nil
	}

	// Calculate n * (n + 1) / 2.
	// We can compute (n / 2) * (n + 1) if n is even, or n * ((n + 1) / 2) if n is odd.
	// This reduces the size of intermediate terms and helps prevent overflow.
	var term1, term2 int
	if n%2 == 0 {
		term1 = n / 2
		term2 = n + 1
	} else {
		term1 = n
		term2 = (n-1)/2 + 1
	}

	// Check if term1 * term2 overflows math.MaxInt
	if term1 > 0 && term2 > math.MaxInt/term1 {
		return 0, errors.New("sum of natural numbers overflows int")
	}
	return term1 * term2, nil
}
