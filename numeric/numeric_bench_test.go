package numeric

import (
	"testing"
)

func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Min(12345, 67890)
	}
}

func BenchmarkMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Max(12345, 67890)
	}
}

func BenchmarkClamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Clamp(25, 10, 20)
	}
}

func BenchmarkAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Abs(-42)
	}
}

func BenchmarkSign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Sign(-15)
	}
}

func BenchmarkIsEven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsEven(12345678)
	}
}

func BenchmarkIsOdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsOdd(12345678)
	}
}
