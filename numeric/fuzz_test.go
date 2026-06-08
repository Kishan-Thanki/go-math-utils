package numeric

import (
	"math"
	"testing"
)

func FuzzClamp(f *testing.F) {
	f.Add(15, 10, 20)
	f.Add(5, 10, 20)
	f.Add(25, 10, 20)
	f.Fuzz(func(t *testing.T, val, min, max int) {
		if min > max {
			defer func() {
				if r := recover(); r == nil {
					t.Fatalf("Clamp(%d, %d, %d) did not panic with min > max", val, min, max)
				}
			}()
			_ = Clamp(val, min, max)
			return
		}
		res := Clamp(val, min, max)
		if res < min {
			t.Fatalf("Clamp(%d, %d, %d) = %d is less than min", val, min, max, res)
		}
		if res > max {
			t.Fatalf("Clamp(%d, %d, %d) = %d is greater than max", val, min, max, res)
		}
		if val >= min && val <= max {
			if res != val {
				t.Fatalf("Clamp(%d, %d, %d) = %d; want val", val, min, max, res)
			}
		}
	})
}

func FuzzAbs(f *testing.F) {
	f.Add(42)
	f.Add(-42)
	f.Add(0)
	f.Fuzz(func(t *testing.T, x int) {
		res := Abs(x)
		if x == math.MinInt {
			if res != math.MinInt {
				t.Fatalf("Abs(math.MinInt) = %d; want math.MinInt", res)
			}
			return
		}
		if res < 0 {
			t.Fatalf("Abs(%d) = %d is negative", x, res)
		}
		if x < 0 {
			if res != -x {
				t.Fatalf("Abs(%d) = %d; want %d", x, res, -x)
			}
		} else {
			if res != x {
				t.Fatalf("Abs(%d) = %d; want %d", x, res, x)
			}
		}
	})
}

func FuzzSign(f *testing.F) {
	f.Add(10)
	f.Add(-10)
	f.Add(0)
	f.Fuzz(func(t *testing.T, x int) {
		res := Sign(x)
		if x < 0 {
			if res != -1 {
				t.Fatalf("Sign(%d) = %d; want -1", x, res)
			}
		} else if x > 0 {
			if res != 1 {
				t.Fatalf("Sign(%d) = %d; want 1", x, res)
			}
		} else {
			if res != 0 {
				t.Fatalf("Sign(%d) = %d; want 0", x, res)
			}
		}
	})
}
