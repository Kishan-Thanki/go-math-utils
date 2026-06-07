// factorial.go
package integer

import (
	"errors"
	"fmt"
	"math"
)

// Factorial returns the factorial of a non-negative integer n.
// If n is negative, or if the calculation would overflow the limits of type T, it returns an error.
func Factorial[T Integer](n T) (T, error) {
	if n < 0 {
		var zero T
		return zero, errors.New("factorial of a negative number is undefined")
	}
	var result T = 1
	for i := T(1); i <= n; i++ {
		uResult := uint64(result)
		uI := uint64(i)
		if uResult > math.MaxUint64/uI {
			var zero T
			return zero, fmt.Errorf("factorial of %v overflows type", n)
		}
		next := uResult * uI
		if uint64(T(next)) != next || T(next) < 0 {
			var zero T
			return zero, fmt.Errorf("factorial of %v overflows type", n)
		}
		result = T(next)
	}
	return result, nil
}
