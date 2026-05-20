// factorial.go
package mathutils

import (
	"errors"
	"fmt"
	"math"
)

// maxFactorialN holds the maximum input value for Factorial that will not cause integer overflow.
var maxFactorialN int

func init() {
	if math.MaxInt == math.MaxInt64 {
		maxFactorialN = 20
	} else {
		maxFactorialN = 12
	}
}

// Factorial returns the factorial of a non-negative integer n.
// If n is negative, or if the calculation would overflow the integer limit, it returns an error.
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial of a negative number is undefined")
	}
	if n > maxFactorialN {
		return 0, fmt.Errorf("factorial of %d overflows int", n)
	}

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}
