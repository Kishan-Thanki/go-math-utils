package sequence

import (
	"testing"
)

func FuzzFibonacci(f *testing.F) {
	f.Add(6)
	f.Add(-1)
	f.Fuzz(func(t *testing.T, n int) {
		res, err := Fibonacci(n)
		if err != nil {
			return
		}
		if n == 0 {
			if res != 0 {
				t.Fatalf("Fibonacci(0) = %d; want 0", res)
			}
		}
		if n == 1 {
			if res != 1 {
				t.Fatalf("Fibonacci(1) = %d; want 1", res)
			}
		}
	})
}

func FuzzSumNaturalNumbers(f *testing.F) {
	f.Add(5)
	f.Add(-1)
	f.Fuzz(func(t *testing.T, n int) {
		res, err := SumNaturalNumbers(n)
		if err != nil {
			return
		}
		if n == 0 {
			if res != 0 {
				t.Fatalf("SumNaturalNumbers(0) = %d; want 0", res)
			}
			return
		}
		if res < 0 {
			t.Fatalf("SumNaturalNumbers(%d) = %d cannot be negative", n, res)
		}
	})
}
