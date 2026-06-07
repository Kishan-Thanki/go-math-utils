package stats

import (
	"math/rand"
	"testing"
)

func generateSlice(size int) []int {
	rng := rand.New(rand.NewSource(42))
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = rng.Intn(100)
	}
	return s
}

func BenchmarkAverage_Small(b *testing.B) {
	slice := generateSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Average(slice)
	}
}

func BenchmarkAverage_Large(b *testing.B) {
	slice := generateSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Average(slice)
	}
}

func BenchmarkMedian_Small(b *testing.B) {
	slice := generateSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Median(slice)
	}
}

func BenchmarkMedian_Large(b *testing.B) {
	slice := generateSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Median(slice)
	}
}

func BenchmarkMode_Small(b *testing.B) {
	slice := generateSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Mode(slice)
	}
}

func BenchmarkMode_Large(b *testing.B) {
	slice := generateSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Mode(slice)
	}
}

func BenchmarkSumSlice_Small(b *testing.B) {
	slice := generateSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = SumSlice(slice)
	}
}

func BenchmarkSumSlice_Large(b *testing.B) {
	slice := generateSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = SumSlice(slice)
	}
}

func BenchmarkMinSlice_Small(b *testing.B) {
	slice := generateSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = MinSlice(slice)
	}
}

func BenchmarkMinSlice_Large(b *testing.B) {
	slice := generateSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = MinSlice(slice)
	}
}

func BenchmarkMaxSlice_Small(b *testing.B) {
	slice := generateSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = MaxSlice(slice)
	}
}

func BenchmarkMaxSlice_Large(b *testing.B) {
	slice := generateSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = MaxSlice(slice)
	}
}
