// numeric_test.go
package numeric

import (
	"math"
	"testing"
)

func TestMin(t *testing.T) {
	// Int test
	if res := Min(10, 20); res != 10 {
		t.Errorf("Min(10, 20) = %d; want 10", res)
	}
	if res := Min(-5, -10); res != -10 {
		t.Errorf("Min(-5, -10) = %d; want -10", res)
	}

	// Float test
	if res := Min(3.14, 2.71); res != 2.71 {
		t.Errorf("Min(3.14, 2.71) = %f; want 2.71", res)
	}
}

func TestMax(t *testing.T) {
	// Int test
	if res := Max(10, 20); res != 20 {
		t.Errorf("Max(10, 20) = %d; want 20", res)
	}
	if res := Max(-5, -10); res != -5 {
		t.Errorf("Max(-5, -10) = %d; want -5", res)
	}

	// Float test
	if res := Max(3.14, 2.71); res != 3.14 {
		t.Errorf("Max(3.14, 2.71) = %f; want 3.14", res)
	}
}

func TestClamp(t *testing.T) {
	// Clamp within range
	if res := Clamp(15, 10, 20); res != 15 {
		t.Errorf("Clamp(15, 10, 20) = %d; want 15", res)
	}

	// Clamp below range
	if res := Clamp(5, 10, 20); res != 10 {
		t.Errorf("Clamp(5, 10, 20) = %d; want 10", res)
	}

	// Clamp above range
	if res := Clamp(25, 10, 20); res != 20 {
		t.Errorf("Clamp(25, 10, 20) = %d; want 20", res)
	}

	// Test float clamp
	if res := Clamp(1.5, 2.0, 3.0); res != 2.0 {
		t.Errorf("Clamp(1.5, 2.0, 3.0) = %f; want 2.0", res)
	}

	// Test panic when min > max
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Clamp(10, 20, 5) did not panic as expected")
		}
	}()
	_ = Clamp(10, 20, 5)
}

func TestAbs(t *testing.T) {
	// Positive int
	if res := Abs(10); res != 10 {
		t.Errorf("Abs(10) = %d; want 10", res)
	}

	// Negative int
	if res := Abs(-10); res != 10 {
		t.Errorf("Abs(-10) = %d; want 10", res)
	}

	// Float
	if res := Abs(-3.14); res != 3.14 {
		t.Errorf("Abs(-3.14) = %f; want 3.14", res)
	}

	// math.MinInt wrapping behavior
	if res := Abs(math.MinInt); res != math.MinInt {
		t.Errorf("Abs(math.MinInt) = %d; want math.MinInt due to wrapping", res)
	}
}

func TestSign(t *testing.T) {
	// Negative
	if res := Sign(-15); res != -1 {
		t.Errorf("Sign(-15) = %d; want -1", res)
	}
	if res := Sign(-3.14); res != -1 {
		t.Errorf("Sign(-3.14) = %d; want -1", res)
	}

	// Zero
	if res := Sign(0); res != 0 {
		t.Errorf("Sign(0) = %d; want 0", res)
	}
	if res := Sign(0.0); res != 0 {
		t.Errorf("Sign(0.0) = %d; want 0", res)
	}

	// Positive
	if res := Sign(42); res != 1 {
		t.Errorf("Sign(42) = %d; want 1", res)
	}
	if res := Sign(1.23); res != 1 {
		t.Errorf("Sign(1.23) = %d; want 1", res)
	}
}

func TestIsEvenIsOdd(t *testing.T) {
	// Signed ints
	if !IsEven(4) {
		t.Errorf("IsEven(4) should be true")
	}
	if IsEven(5) {
		t.Errorf("IsEven(5) should be false")
	}
	if !IsOdd(5) {
		t.Errorf("IsOdd(5) should be true")
	}
	if IsOdd(4) {
		t.Errorf("IsOdd(4) should be false")
	}

	// Unsigned ints
	var uval uint = 10
	if !IsEven(uval) {
		t.Errorf("IsEven(uint(10)) should be true")
	}

	// Negative values
	if !IsEven(-4) {
		t.Errorf("IsEven(-4) should be true")
	}
	if !IsOdd(-3) {
		t.Errorf("IsOdd(-3) should be true")
	}
}
