// integer_test.go
package integer

import (
	"math"
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b        int
		expect      int
		expectPanic bool
	}{
		{56, 98, 14, false},
		{48, 180, 12, false},
		{101, 103, 1, false},
		{0, 10, 10, false},
		{10, 0, 10, false},
		{0, 0, 0, false},
		{-56, 98, 14, false},
		{56, -98, 14, false},
		{-56, -98, 14, false},
		{math.MinInt, 0, 0, true},
		{math.MinInt, math.MinInt, 0, true},
		{math.MinInt, 2, 2, false},
	}

	for _, tt := range tests {
		t.Run("GCD", func(t *testing.T) {
			if tt.expectPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("GCD(%d, %d): expected panic, but did not panic", tt.a, tt.b)
					}
				}()
			}
			result := GCD(tt.a, tt.b)
			if !tt.expectPanic && result != tt.expect {
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

func TestGenericTypes(t *testing.T) {
	// Test GCD with int8
	var a8, b8 int8 = 12, 18
	if res := GCD(a8, b8); res != 6 {
		t.Errorf("GCD(int8(12), int8(18)) = %d; want 6", res)
	}

	// Test GCD with uint32
	var au32, bu32 uint32 = 40, 60
	if res := GCD(au32, bu32); res != 20 {
		t.Errorf("GCD(uint32(40), uint32(60)) = %d; want 20", res)
	}

	// Test Factorial with int16
	var n16 int16 = 5
	if res, err := Factorial(n16); err != nil || res != 120 {
		t.Errorf("Factorial(int16(5)) = (%v, %v); want (120, nil)", res, err)
	}

	// Test Power with uint64
	var base, exp uint64 = 2, 10
	if res, err := Power(base, exp); err != nil || res != 1024 {
		t.Errorf("Power(uint64(2), uint64(10)) = (%v, %v); want (1024, nil)", res, err)
	}
}

func TestEulerTotient(t *testing.T) {
	tests := []struct {
		n      int
		expect int
		err    bool
	}{
		{9, 6, false},
		{10, 4, false},
		{1, 1, false},
		{0, 0, true},
		{-5, 0, true},
	}

	for _, tt := range tests {
		t.Run("EulerTotient", func(t *testing.T) {
			res, err := EulerTotient(tt.n)
			if tt.err && err == nil {
				t.Errorf("EulerTotient(%d): expected error, got nil", tt.n)
			} else if !tt.err && err != nil {
				t.Errorf("EulerTotient(%d): expected no error, got %v", tt.n, err)
			} else if !tt.err && res != tt.expect {
				t.Errorf("EulerTotient(%d) = %d; want %d", tt.n, res, tt.expect)
			}
		})
	}
}

func TestDivisors(t *testing.T) {
	tests := []struct {
		n      int
		expect []int
		err    bool
	}{
		{12, []int{1, 2, 3, 4, 6, 12}, false},
		{7, []int{1, 7}, false},
		{-12, []int{1, 2, 3, 4, 6, 12}, false},
		{0, nil, true},
	}

	for _, tt := range tests {
		t.Run("Divisors", func(t *testing.T) {
			res, err := Divisors(tt.n)
			if tt.err && err == nil {
				t.Errorf("Divisors(%d): expected error, got nil", tt.n)
			} else if !tt.err && err != nil {
				t.Errorf("Divisors(%d): expected no error, got %v", tt.n, err)
			} else if !tt.err {
				if len(res) != len(tt.expect) {
					t.Errorf("Divisors(%d) = %v; want %v", tt.n, res, tt.expect)
					return
				}
				for i := range res {
					if res[i] != tt.expect[i] {
						t.Errorf("Divisors(%d) = %v; want %v", tt.n, res, tt.expect)
						break
					}
				}
			}
		})
	}

	// Test boundary overflow check using int8
	t.Run("DivisorsInt8Overflow", func(t *testing.T) {
		var n int8 = -128
		_, err := Divisors(n)
		if err == nil {
			t.Errorf("Divisors(int8(-128)): expected overflow error because divisor 128 overflows int8, got nil")
		}
	})
}

func TestNextPrime(t *testing.T) {
	tests := []struct {
		n      int
		expect int
	}{
		{14, 17},
		{3, 5},
		{2, 3},
		{-5, 2},
		{0, 2},
	}

	for _, tt := range tests {
		t.Run("NextPrime", func(t *testing.T) {
			res := NextPrime(tt.n)
			if res != tt.expect {
				t.Errorf("NextPrime(%d) = %d; want %d", tt.n, res, tt.expect)
			}
		})
	}

	// Test panic on overflow using int8
	t.Run("NextPrimeInt8Panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("NextPrime(int8(127)) did not panic on overflow")
			}
		}()
		var n int8 = 127
		_ = NextPrime(n)
	})
}

func TestNthPrime(t *testing.T) {
	tests := []struct {
		n      int
		expect int
		err    bool
	}{
		{1, 2, false},
		{5, 11, false},
		{10, 29, false},
		{31, 127, false},
		{0, 0, true},
		{-5, 0, true},
	}

	for _, tt := range tests {
		t.Run("NthPrime", func(t *testing.T) {
			res, err := NthPrime(tt.n)
			if tt.err && err == nil {
				t.Errorf("NthPrime(%d): expected error, got nil", tt.n)
			} else if !tt.err && err != nil {
				t.Errorf("NthPrime(%d): expected no error, got %v", tt.n, err)
			} else if !tt.err && res != tt.expect {
				t.Errorf("NthPrime(%d) = %d; want %d", tt.n, res, tt.expect)
			}
		})
	}

	// Test boundary overflow using int8
	t.Run("NthPrimeInt8Overflow", func(t *testing.T) {
		var n int8 = 32 // 32nd prime is 131, which overflows int8
		_, err := NthPrime(n)
		if err == nil {
			t.Errorf("NthPrime(int8(32)): expected error, got nil")
		}
	})
}

func TestPerfectNumber(t *testing.T) {
	tests := []struct {
		n      int
		expect bool
	}{
		{6, true},   // 1 + 2 + 3 = 6
		{28, true},  // 1 + 2 + 4 + 7 + 14 = 28
		{496, true},
		{12, false},
		{1, false},
		{0, false},
		{-6, false},
	}

	for _, tt := range tests {
		t.Run("PerfectNumber", func(t *testing.T) {
			res := PerfectNumber(tt.n)
			if res != tt.expect {
				t.Errorf("PerfectNumber(%d) = %v; want %v", tt.n, res, tt.expect)
			}
		})
	}
}


func BenchmarkEulerTotient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = EulerTotient(1000)
	}
}

func BenchmarkDivisors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Divisors(1000)
	}
}

func BenchmarkNextPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NextPrime(1000)
	}
}

func BenchmarkNthPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NthPrime(100)
	}
}

func BenchmarkPerfectNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PerfectNumber(28)
	}
}


