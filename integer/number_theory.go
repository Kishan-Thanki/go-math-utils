// number_theory.go
package integer

import (
	"errors"
	"fmt"
	"slices"
)

// EulerTotient calculates Euler's totient function phi(n) of a positive integer.
// phi(n) counts the positive integers up to n that are relatively prime to n.
// Returns an error if n <= 0.
func EulerTotient[T Integer](n T) (T, error) {
	if n <= 0 {
		var zero T
		return zero, fmt.Errorf("euler totient is only defined for positive integers, got %v", n)
	}

	result := n
	temp := n
	for i := T(2); i <= temp/i; i++ {
		if temp%i == 0 {
			for temp%i == 0 {
				temp /= i
			}
			result -= result / i
		}
	}
	if temp > 1 {
		result -= result / temp
	}
	return result, nil
}

// Divisors returns a sorted slice containing all positive divisors of n.
// Returns an error if n == 0.
func Divisors[T Integer](n T) ([]T, error) {
	if n == 0 {
		return nil, errors.New("divisors are undefined for 0")
	}

	ua := absUint64(n)
	var divisors []T
	for i := uint64(1); i <= ua/i; i++ {
		if ua%i == 0 {
			// Check for overflow when casting back to T
			if uint64(T(i)) != i || T(i) < 0 {
				return nil, fmt.Errorf("divisor %d overflows type", i)
			}
			divisors = append(divisors, T(i))

			other := ua / i
			if other != i {
				if uint64(T(other)) != other || T(other) < 0 {
					return nil, fmt.Errorf("divisor %d overflows type", other)
				}
				divisors = append(divisors, T(other))
			}
		}
	}
	slices.Sort(divisors)
	return divisors, nil
}

// NextPrime returns the smallest prime number strictly greater than n.
// If the next prime overflows the maximum value representable by type T, it panics.
func NextPrime[T Integer](n T) T {
	_, max := getLimits[T]()
	if n >= max {
		panic("next prime overflows type")
	}

	start := uint64(2)
	if n >= 2 {
		start = uint64(n) + 1
	}

	for {
		if isPrimeUint64(start) {
			if uint64(T(start)) != start || T(start) < 0 {
				panic("next prime overflows type")
			}
			return T(start)
		}
		start++
	}
}

// NthPrime returns the n-th prime number (where 1st prime is 2, 2nd is 3, etc.).
// Returns an error if n <= 0, or if the n-th prime overflows type T.
func NthPrime[T Integer](n T) (T, error) {
	if n <= 0 {
		var zero T
		return zero, errors.New("n-th prime is only defined for positive integers")
	}

	count := T(0)
	candidate := uint64(2)
	for {
		if isPrimeUint64(candidate) {
			count++
			if count == n {
				if uint64(T(candidate)) != candidate || T(candidate) < 0 {
					var zero T
					return zero, fmt.Errorf("%d-th prime overflows type", n)
				}
				return T(candidate), nil
			}
		}
		candidate++
	}
}

// PerfectNumber returns true if n is a perfect number.
// A perfect number is a positive integer that is equal to the sum of its positive proper divisors.
// (Proper divisors are positive divisors excluding the number itself).
func PerfectNumber[T Integer](n T) bool {
	if n <= 1 {
		return false
	}

	divs, err := Divisors(n)
	if err != nil {
		return false
	}

	var sum T
	for _, d := range divs {
		if d == n {
			continue
		}
		_, max := getLimits[T]()
		if sum > max-d {
			return false
		}
		sum += d
	}
	return sum == n
}

