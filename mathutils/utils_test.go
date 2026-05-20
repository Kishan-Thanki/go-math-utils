// utils_test.go
package mathutils

import (
	"math"
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b   int
		expect int
	}{
		{56, 98, 14},
		{48, 180, 12},
		{101, 103, 1},
		{0, 10, 10},
		{10, 0, 10},
		{0, 0, 0},
		{-56, 98, 14},
		{56, -98, 14},
		{-56, -98, 14},
		{math.MinInt, 0, math.MinInt}, // Returns absolute value represented in int
		{math.MinInt, 2, 2},
	}

	for _, tt := range tests {
		t.Run("GCD", func(t *testing.T) {
			result := GCD(tt.a, tt.b)
			if result != tt.expect {
				t.Errorf("GCD(%d, %d): expected %d, got %d", tt.a, tt.b, tt.expect, result)
			}
		})
	}
}

func TestLCM(t *testing.T) {
	tests := []struct {
		a, b      int
		expect    int
		expectErr bool
	}{
		{56, 98, 392, false},
		{4, 5, 20, false},
		{7, 13, 91, false},
		{0, 10, 0, false},
		{10, 0, 0, false},
		{0, 0, 0, false}, // Should not panic anymore
		{-4, 5, 20, false},
		{4, -5, 20, false},
		{-4, -5, 20, false},
		{math.MaxInt, 2, 0, true}, // Should overflow
	}

	for _, tt := range tests {
		t.Run("LCM", func(t *testing.T) {
			result, err := LCM(tt.a, tt.b)
			if tt.expectErr {
				if err == nil {
					t.Errorf("LCM(%d, %d): expected error, got nil", tt.a, tt.b)
				}
			} else {
				if err != nil {
					t.Errorf("LCM(%d, %d): expected no error, got %v", tt.a, tt.b, err)
				} else if result != tt.expect {
					t.Errorf("LCM(%d, %d): expected %d, got %d", tt.a, tt.b, tt.expect, result)
				}
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	// Dynamically set overflow input based on system bitness
	var overflowN int
	if math.MaxInt == math.MaxInt64 {
		overflowN = 21
	} else {
		overflowN = 13
	}

	tests := []struct {
		n      int
		expect int
		err    bool
	}{
		{5, 120, false},
		{0, 1, false},
		{1, 1, false},
		{7, 5040, false},
		{-1, 0, true},
		{overflowN, 0, true},
	}

	for _, tt := range tests {
		t.Run("Factorial", func(t *testing.T) {
			result, err := Factorial(tt.n)
			if tt.err && err == nil {
				t.Errorf("Factorial(%d): expected an error, but got nil", tt.n)
			} else if !tt.err && err != nil {
				t.Errorf("Factorial(%d): expected no error, but got %v", tt.n, err)
			} else if !tt.err && result != tt.expect {
				t.Errorf("Factorial(%d): expected %d, got %d", tt.n, tt.expect, result)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n      int
		expect bool
	}{
		{29, true},
		{15, false},
		{1, false},
		{0, false},
		{-7, false},
		{2, true},
		{3, true},
		{4, false},
	}

	// Dynamically append large prime check for 64-bit systems to compile safely on 32-bit systems
	if math.MaxInt == math.MaxInt64 {
		tests = append(tests, struct {
			n      int
			expect bool
		}{9223372036854775783, true})
	}

	for _, tt := range tests {
		t.Run("IsPrime", func(t *testing.T) {
			result := IsPrime(tt.n)
			if result != tt.expect {
				t.Errorf("IsPrime(%d): expected %v, got %v", tt.n, tt.expect, result)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	// Dynamically set overflow input based on system bitness
	var overflowN int
	if math.MaxInt == math.MaxInt64 {
		overflowN = 93
	} else {
		overflowN = 47
	}

	tests := []struct {
		n      int
		expect int
		err    bool
	}{
		{6, 8, false},
		{0, 0, false},
		{1, 1, false},
		{10, 55, false},
		{-1, 0, true},
		{-5, 0, true},
		{overflowN, 0, true},
	}

	for _, tt := range tests {
		t.Run("Fibonacci", func(t *testing.T) {
			result, err := Fibonacci(tt.n)
			if tt.err && err == nil {
				t.Errorf("Fibonacci(%d): expected an error, but got nil", tt.n)
			} else if !tt.err && err != nil {
				t.Errorf("Fibonacci(%d): expected no error, but got %v", tt.n, err)
			} else if !tt.err && result != tt.expect {
				t.Errorf("Fibonacci(%d): expected %d, got %d", tt.n, tt.expect, result)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		a, b   int
		expect int
		err    bool
	}{
		{2, 3, 8, false},
		{5, 0, 1, false},
		{3, 2, 9, false},
		{10, 2, 100, false},
		{-2, 3, -8, false},
		{2, -3, 0, true},   // Negative exponent
		{10, 100, 0, true}, // Overflow
	}

	for _, tt := range tests {
		t.Run("Power", func(t *testing.T) {
			result, err := Power(tt.a, tt.b)
			if tt.err && err == nil {
				t.Errorf("Power(%d, %d): expected an error, but got nil", tt.a, tt.b)
			} else if !tt.err && err != nil {
				t.Errorf("Power(%d, %d): expected no error, but got %v", tt.a, tt.b, err)
			} else if !tt.err && result != tt.expect {
				t.Errorf("Power(%d, %d): expected %d, got %d", tt.a, tt.b, tt.expect, result)
			}
		})
	}
}

func TestPrimeFactors(t *testing.T) {
	tests := []struct {
		n      int
		expect []int
		err    bool
	}{
		{56, []int{2, 2, 2, 7}, false},
		{29, []int{29}, false},
		{1, nil, true},
		{0, nil, true},
		{-56, nil, true},
	}

	for _, tt := range tests {
		t.Run("PrimeFactors", func(t *testing.T) {
			result, err := PrimeFactors(tt.n)
			if tt.err && err == nil {
				t.Errorf("PrimeFactors(%d): expected an error, but got nil", tt.n)
			} else if !tt.err && err != nil {
				t.Errorf("PrimeFactors(%d): expected no error, but got %v", tt.n, err)
			} else if !tt.err {
				if len(result) != len(tt.expect) {
					t.Errorf("PrimeFactors(%d): expected %v, got %v", tt.n, tt.expect, result)
					return
				}
				for i := range result {
					if result[i] != tt.expect[i] {
						t.Errorf("PrimeFactors(%d): expected %v, got %v", tt.n, tt.expect, result)
						break
					}
				}
			}
		})
	}
}

func TestSumNaturalNumbers(t *testing.T) {
	tests := []struct {
		n      int
		expect int
		err    bool
	}{
		{5, 15, false},
		{1, 1, false},
		{10, 55, false},
		{0, 0, false},
		{-1, 0, true},
		{math.MaxInt, 0, true}, // Will overflow
	}

	for _, tt := range tests {
		t.Run("SumNaturalNumbers", func(t *testing.T) {
			result, err := SumNaturalNumbers(tt.n)
			if tt.err && err == nil {
				t.Errorf("SumNaturalNumbers(%d): expected an error, but got nil", tt.n)
			} else if !tt.err && err != nil {
				t.Errorf("SumNaturalNumbers(%d): expected no error, but got %v", tt.n, err)
			} else if !tt.err && result != tt.expect {
				t.Errorf("SumNaturalNumbers(%d): expected %d, got %d", tt.n, tt.expect, result)
			}
		})
	}
}

func BenchmarkGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GCD(123456789, 987654321)
	}
}

func BenchmarkLCM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = LCM(12345, 67890)
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Factorial(12)
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsPrime(997)
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Fibonacci(40)
	}
}

func BenchmarkPower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Power(3, 15)
	}
}

func BenchmarkPrimeFactors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = PrimeFactors(12345678)
	}
}

func BenchmarkSumNaturalNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = SumNaturalNumbers(1000)
	}
}
