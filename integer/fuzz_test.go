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
				if r != "gcd result overflows type" {
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

func FuzzEulerTotient(f *testing.F) {
	f.Add(9)
	f.Add(-1)
	f.Fuzz(func(t *testing.T, n int) {
		res, err := EulerTotient(n)
		if err != nil {
			return
		}
		if res <= 0 {
			t.Fatalf("EulerTotient(%d) = %d is not positive", n, res)
		}
	})
}

func FuzzDivisors(f *testing.F) {
	f.Add(12)
	f.Add(0)
	f.Fuzz(func(t *testing.T, n int) {
		res, err := Divisors(n)
		if err != nil {
			return
		}
		for _, d := range res {
			if d <= 0 {
				t.Fatalf("Divisors(%d) contains non-positive divisor: %d", n, d)
			}
			ua := absUint64(n)
			if ua%uint64(d) != 0 {
				t.Fatalf("divisor %d does not divide absolute value of %d", d, n)
			}
		}
	})
}

func FuzzNextPrime(f *testing.F) {
	f.Add(14)
	f.Add(-5)
	f.Fuzz(func(t *testing.T, n int) {
		defer func() {
			if r := recover(); r != nil {
				if r != "next prime overflows type" {
					t.Fatalf("unexpected panic in NextPrime: %v", r)
				}
			}
		}()
		res := NextPrime(n)
		if res <= n {
			t.Fatalf("NextPrime(%d) = %d is not strictly greater than n", n, res)
		}
		if !isPrimeUint64(uint64(res)) {
			t.Fatalf("NextPrime(%d) = %d is not prime", n, res)
		}
	})
}

func FuzzNthPrime(f *testing.F) {
	f.Add(5)
	f.Add(-1)
	f.Fuzz(func(t *testing.T, n int) {
		res, err := NthPrime(n)
		if err != nil {
			return
		}
		if res < 2 {
			t.Fatalf("NthPrime(%d) = %d is less than 2", n, res)
		}
		if !isPrimeUint64(uint64(res)) {
			t.Fatalf("NthPrime(%d) = %d is not prime", n, res)
		}
	})
}

func FuzzPerfectNumber(f *testing.F) {
	f.Add(6)
	f.Add(12)
	f.Add(-6)
	f.Fuzz(func(t *testing.T, n int) {
		res := PerfectNumber(n)
		if n <= 1 && res {
			t.Fatalf("PerfectNumber(%d) returned true, but n <= 1", n)
		}
		if res {
			divs, err := Divisors(n)
			if err != nil {
				t.Fatalf("PerfectNumber is true but Divisors failed for %d: %v", n, err)
			}
			var sum int
			for _, d := range divs {
				if d != n {
					sum += d
				}
			}
			if sum != n {
				t.Fatalf("PerfectNumber is true but sum of proper divisors of %d is %d", n, sum)
			}
		}
	})
}
