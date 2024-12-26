// utils_test.go
package mathutils

import (
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b   int
		expect int
	}{
		{56, 98, 14},  // Example where GCD is 14
		{48, 180, 12}, // Another example
		{101, 103, 1}, // Case with prime numbers
		{0, 10, 10},   // Edge case where one number is 0
		{10, 0, 10},   // Edge case where other number is 0
		{0, 0, 0},     // Edge case where both numbers are 0
	}

	for _, tt := range tests {
		t.Run("GCD", func(t *testing.T) {
			result := GCD(tt.a, tt.b)
			if result != tt.expect {
				t.Errorf("expected %d, got %d", tt.expect, result)
			}
		})
	}
}

func TestLCM(t *testing.T) {
	tests := []struct {
		a, b   int
		expect int
	}{
		{56, 98, 392}, // Example where LCM is 392
		{4, 5, 20},    // Another example
		{7, 13, 91},   // Case with two prime numbers
		{0, 10, 0},    // Edge case where one number is 0
		{10, 0, 0},    // Edge case where other number is 0
	}

	for _, tt := range tests {
		t.Run("LCM", func(t *testing.T) {
			result := LCM(tt.a, tt.b)
			if result != tt.expect {
				t.Errorf("expected %d, got %d", tt.expect, result)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		n      int
		expect int
		err    bool
	}{
		{5, 120, false},  // Factorial of 5 is 120
		{0, 1, false},    // Factorial of 0 is 1
		{1, 1, false},    // Factorial of 1 is 1
		{7, 5040, false}, // Factorial of 7 is 5040
		{-1, 0, true},    // Factorial of -1 should return an error
	}

	for _, tt := range tests {
		t.Run("Factorial", func(t *testing.T) {
			result, err := Factorial(tt.n)
			if tt.err && err == nil {
				t.Errorf("expected an error for input %d, but got nil", tt.n)
			} else if !tt.err && err != nil {
				t.Errorf("expected no error for input %d, but got %v", tt.n, err)
			} else if result != tt.expect {
				t.Errorf("expected %d, got %d", tt.expect, result)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n      int
		expect bool
	}{
		{29, true},  // 29 is prime
		{15, false}, // 15 is not prime
		{1, false},  // 1 is not prime
		{2, true},   // 2 is prime
		{4, false},  // 4 is not prime
	}

	for _, tt := range tests {
		t.Run("IsPrime", func(t *testing.T) {
			result := IsPrime(tt.n)
			if result != tt.expect {
				t.Errorf("expected %v, got %v", tt.expect, result)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n      int
		expect int
	}{
		{6, 8},   // 6th Fibonacci number is 8
		{0, 0},   // 0th Fibonacci number is 0
		{1, 1},   // 1st Fibonacci number is 1
		{10, 55}, // 10th Fibonacci number is 55
		{-1, -1}, // Invalid case: negative index
	}

	for _, tt := range tests {
		t.Run("Fibonacci", func(t *testing.T) {
			result := Fibonacci(tt.n)
			if result != tt.expect {
				t.Errorf("expected %d, got %d", tt.expect, result)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		a, b   int
		expect int
	}{
		{2, 3, 8},    // 2^3 = 8
		{5, 0, 1},    // 5^0 = 1
		{3, 2, 9},    // 3^2 = 9
		{10, 2, 100}, // 10^2 = 100
		{-2, 3, -8},  // Negative base with odd exponent
	}

	for _, tt := range tests {
		t.Run("Power", func(t *testing.T) {
			result := Power(tt.a, tt.b)
			if result != tt.expect {
				t.Errorf("expected %d, got %d", tt.expect, result)
			}
		})
	}
}

func TestSumNaturalNumbers(t *testing.T) {
	tests := []struct {
		n      int
		expect int
	}{
		{5, 15},  // Sum of first 5 natural numbers is 15
		{1, 1},   // Sum of first 1 natural number is 1
		{10, 55}, // Sum of first 10 natural numbers is 55
		{0, 0},   // Sum of first 0 natural numbers is 0
	}

	for _, tt := range tests {
		t.Run("SumNaturalNumbers", func(t *testing.T) {
			result := SumNaturalNumbers(tt.n)
			if result != tt.expect {
				t.Errorf("expected %d, got %d", tt.expect, result)
			}
		})
	}
}
