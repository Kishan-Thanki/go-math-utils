// factorial.go
package mathutils

import (
	"errors"
)

// Factorial returns the factorial of a non-negative integer n.
// If n is negative, it returns an error.
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial of a negative number is undefined")
	}

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}
