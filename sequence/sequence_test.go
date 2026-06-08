// sequence_test.go
package sequence

import (
	"math"
	"testing"
)

func TestFibonacci(t *testing.T) {
	// Dynamically set overflow input based on system bitness
	var overflowN int
	if math.MaxInt == math.MaxInt64 {
		overflowN = 93
	} else {
		overflowN = 47
	}

	tests := []struct {
		n      int
		expect int
		err    bool
	}{
		{6, 8, false},
		{0, 0, false},
		{1, 1, false},
		{10, 55, false},
		{-1, 0, true},
		{-5, 0, true},
		{overflowN, 0, true},
	}

	for _, tt := range tests {
		t.Run("Fibonacci", func(t *testing.T) {
			result, err := Fibonacci(tt.n)
			if tt.err && err == nil {
				t.Errorf("Fibonacci(%d): expected an error, but got nil", tt.n)
			} else if !tt.err && err != nil {
				t.Errorf("Fibonacci(%d): expected no error, but got %v", tt.n, err)
			} else if !tt.err && result != tt.expect {
				t.Errorf("Fibonacci(%d): expected %d, got %d", tt.n, tt.expect, result)
			}
		})
	}
}

func TestSumNaturalNumbers(t *testing.T) {
	tests := []struct {
		n      int
		expect int
		err    bool
	}{
		{5, 15, false},
		{1, 1, false},
		{10, 55, false},
		{0, 0, false},
		{-1, 0, true},
		{math.MaxInt, 0, true}, // Will overflow
	}

	for _, tt := range tests {
		t.Run("SumNaturalNumbers", func(t *testing.T) {
			result, err := SumNaturalNumbers(tt.n)
			if tt.err && err == nil {
				t.Errorf("SumNaturalNumbers(%d): expected an error, but got nil", tt.n)
			} else if !tt.err && err != nil {
				t.Errorf("SumNaturalNumbers(%d): expected no error, but got %v", tt.n, err)
			} else if !tt.err && result != tt.expect {
				t.Errorf("SumNaturalNumbers(%d): expected %d, got %d", tt.n, tt.expect, result)
			}
		})
	}
}

func TestGenericTypes(t *testing.T) {
	// Test Fibonacci with int16
	var n16 int16 = 10
	if res, err := Fibonacci(n16); err != nil || res != 55 {
		t.Errorf("Fibonacci(int16(10)) = (%v, %v); want (55, nil)", res, err)
	}

	// Test SumNaturalNumbers with uint32
	var nu32 uint32 = 10
	if res, err := SumNaturalNumbers(nu32); err != nil || res != 55 {
		t.Errorf("SumNaturalNumbers(uint32(10)) = (%v, %v); want (55, nil)", res, err)
	}
}
