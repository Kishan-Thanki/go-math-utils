// sequence_bench_test.go
package sequence

import "testing"

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Fibonacci(40)
	}
}

func BenchmarkSumNaturalNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = SumNaturalNumbers(1000)
	}
}
