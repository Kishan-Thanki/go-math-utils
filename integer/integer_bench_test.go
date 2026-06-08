// integer_bench_test.go
package integer

import "testing"

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
