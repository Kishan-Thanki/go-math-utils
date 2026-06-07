package integer

import (
	"testing"
)

func FuzzGCD(f *testing.F) {
	f.Add(48, 18)
	f.Add(0, 5)
	f.Add(-10, 20)
	f.Fuzz(func(t *testing.T, a, b int) {
		defer func() {
			if r := recover(); r != nil {
				if r != "gcd result overflows int" {
					t.Fatalf("unexpected panic: %v", r)
				}
			}
		}()
		res := GCD(a, b)
		if res < 0 {
			t.Fatalf("GCD(%d, %d) returned negative: %d", a, b, res)
		}
		if res > 0 {
			if a%res != 0 {
				t.Fatalf("GCD(%d, %d) = %d does not divide a", a, b, res)
			}
			if b%res != 0 {
				t.Fatalf("GCD(%d, %d) = %d does not divide b", a, b, res)
			}
		}
	})
}

func FuzzLCM(f *testing.F) {
	f.Add(56, 98)
	f.Add(0, 0)
	f.Add(-4, 5)
	f.Fuzz(func(t *testing.T, a, b int) {
		res, err := LCM(a, b)
		if err != nil {
			return
		}
		if a == 0 || b == 0 {
			if res != 0 {
				t.Fatalf("LCM(%d, %d) = %d; want 0", a, b, res)
			}
			return
		}
		if res%a != 0 {
			t.Fatalf("LCM(%d, %d) = %d is not divisible by a", a, b, res)
		}
		if res%b != 0 {
			t.Fatalf("LCM(%d, %d) = %d is not divisible by b", a, b, res)
		}
	})
}

func FuzzPower(f *testing.F) {
	f.Add(2, 3)
	f.Add(-2, 3)
	f.Add(5, 0)
	f.Fuzz(func(t *testing.T, a, b int) {
		res, err := Power(a, b)
		if err != nil {
			return
		}
		if b == 0 {
			if res != 1 {
				t.Fatalf("Power(%d, 0) = %d; want 1", a, res)
			}
		}
	})
}

func FuzzFactorial(f *testing.F) {
	f.Add(5)
	f.Add(-1)
	f.Fuzz(func(t *testing.T, n int) {
		res, err := Factorial(n)
		if err != nil {
			return
		}
		if n == 0 || n == 1 {
			if res != 1 {
				t.Fatalf("Factorial(%d) = %d; want 1", n, res)
			}
		}
	})
}
