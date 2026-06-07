// fibonacci.go
package sequence

import (
	"fmt"
	"math"
)

// maxFibonacciN holds the maximum input value for Fibonacci that will not cause integer overflow.
var maxFibonacciN int

func init() {
	if math.MaxInt == math.MaxInt64 {
		maxFibonacciN = 92
	} else {
		maxFibonacciN = 46
	}
}

// Fibonacci returns the nth Fibonacci number using an iterative approach.
// If n is negative or if the calculation would overflow the integer limit, it returns an error.
func Fibonacci(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("fibonacci index %d is negative", n)
	}
	if n > maxFibonacciN {
		return 0, fmt.Errorf("fibonacci of %d overflows int", n)
	}
	if n <= 1 {
		return n, nil
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b, nil
}
