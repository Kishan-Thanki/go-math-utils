// prime.go
package integer

import (
	"fmt"
)

// IsPrime checks if a number n is prime.
// It returns true if n is prime, otherwise false.
// It is safe against integer overflow in the loop condition.
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i <= n/i; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// PrimeFactors returns a slice containing the prime factors of n.
// If n <= 1, it returns an error. It is safe against integer overflow in the loop condition.
func PrimeFactors(n int) ([]int, error) {
	if n <= 1 {
		return nil, fmt.Errorf("prime factorization is only defined for integers greater than 1, got %d", n)
	}

	factors := []int{}
	// Trial division starting at 2
	for i := 2; i <= n/i; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors, nil
}
