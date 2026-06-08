// stats_test.go
package stats

import (
	"testing"
)

func TestAverage(t *testing.T) {
	// Empty slice
	if _, err := Average([]int{}); err == nil {
		t.Errorf("Average of empty slice should return an error")
	}

	// Int slice
	res, err := Average([]int{1, 2, 3, 4, 5})
	if err != nil || res != 3.0 {
		t.Errorf("Average([1,2,3,4,5]) = (%f, %v); want (3.0, nil)", res, err)
	}

	// Float slice
	resFloat, err := Average([]float64{1.5, 2.5, 3.5})
	if err != nil || resFloat != 2.5 {
		t.Errorf("Average([1.5,2.5,3.5]) = (%f, %v); want (2.5, nil)", resFloat, err)
	}
}

func TestMedian(t *testing.T) {
	// Empty slice
	if _, err := Median([]int{}); err == nil {
		t.Errorf("Median of empty slice should return error")
	}

	// Odd length
	res, err := Median([]int{5, 2, 9, 1, 7}) // sorted: 1, 2, 5, 7, 9
	if err != nil || res != 5.0 {
		t.Errorf("Median([5,2,9,1,7]) = (%f, %v); want (5.0, nil)", res, err)
	}

	// Even length
	res, err = Median([]int{5, 2, 9, 1, 7, 12}) // sorted: 1, 2, 5, 7, 9, 12 -> average of 5 and 7 = 6.0
	if err != nil || res != 6.0 {
		t.Errorf("Median([5,2,9,1,7,12]) = (%f, %v); want (6.0, nil)", res, err)
	}

	// Float slice
	resFloat, err := Median([]float64{1.5, 5.5, 2.5, 3.5}) // sorted: 1.5, 2.5, 3.5, 5.5 -> median = 3.0
	if err != nil || resFloat != 3.0 {
		t.Errorf("Median([1.5,5.5,2.5,3.5]) = (%f, %v); want (3.0, nil)", resFloat, err)
	}
}

func TestMode(t *testing.T) {
	// Empty slice
	if _, err := Mode([]int{}); err == nil {
		t.Errorf("Mode of empty slice should return error")
	}

	// Simple mode
	res, err := Mode([]int{1, 3, 3, 2, 3, 4})
	if err != nil || res != 3 {
		t.Errorf("Mode([1,3,3,2,3,4]) = (%d, %v); want (3, nil)", res, err)
	}

	// Multiple modes tie-breaker (returns the one that appeared first)
	res, err = Mode([]int{4, 4, 3, 3, 2}) // both 4 and 3 have count 2, 4 appeared first
	if err != nil || res != 4 {
		t.Errorf("Mode([4,4,3,3,2]) = (%d, %v); want (4, nil)", res, err)
	}

	// String slice (Ordered)
	resStr, err := Mode([]string{"apple", "banana", "banana", "cherry"})
	if err != nil || resStr != "banana" {
		t.Errorf("Mode([apple,banana,banana,cherry]) = (%s, %v); want (banana, nil)", resStr, err)
	}
}

func TestSumSlice(t *testing.T) {
	// Empty slice
	if _, err := SumSlice([]int{}); err == nil {
		t.Errorf("SumSlice of empty slice should return an error")
	}

	// Int slice sum
	res, err := SumSlice([]int{1, 2, 3, 4, 5})
	if err != nil || res != 15 {
		t.Errorf("SumSlice([1,2,3,4,5]) = (%v, %v); want (15, nil)", res, err)
	}

	// Float slice sum
	resFloat, err := SumSlice([]float64{1.5, 2.5, 3.5})
	if err != nil || resFloat != 7.5 {
		t.Errorf("SumSlice([1.5,2.5,3.5]) = (%v, %v); want (7.5, nil)", resFloat, err)
	}

	// Integer overflow check using int8
	var maxVal int8 = 120
	_, err = SumSlice([]int8{maxVal, 10})
	if err == nil {
		t.Errorf("SumSlice(int8(120), 10) should return overflow error")
	}

	// Integer underflow check using int8
	var minVal int8 = -120
	_, err = SumSlice([]int8{minVal, -10})
	if err == nil {
		t.Errorf("SumSlice(int8(-120), -10) should return underflow error")
	}

	// Test float32
	if res, err := SumSlice([]float32{1.5, 2.5}); err != nil || res != 4.0 {
		t.Errorf("SumSlice([]float32) failed: %v", err)
	}

	// Test int16
	if res, err := SumSlice([]int16{1, 2}); err != nil || res != 3 {
		t.Errorf("SumSlice([]int16) failed: %v", err)
	}
	if _, err := SumSlice([]int16{30000, 10000}); err == nil {
		t.Errorf("SumSlice([]int16) overflow not detected")
	}
	if _, err := SumSlice([]int16{-30000, -10000}); err == nil {
		t.Errorf("SumSlice([]int16) underflow not detected")
	}

	// Test int32
	if res, err := SumSlice([]int32{1, 2}); err != nil || res != 3 {
		t.Errorf("SumSlice([]int32) failed: %v", err)
	}
	if _, err := SumSlice([]int32{2000000000, 2000000000}); err == nil {
		t.Errorf("SumSlice([]int32) overflow not detected")
	}
	if _, err := SumSlice([]int32{-2000000000, -2000000000}); err == nil {
		t.Errorf("SumSlice([]int32) underflow not detected")
	}

	// Test int64
	if res, err := SumSlice([]int64{1, 2}); err != nil || res != 3 {
		t.Errorf("SumSlice([]int64) failed: %v", err)
	}
	if _, err := SumSlice([]int64{9000000000000000000, 1000000000000000000}); err == nil {
		t.Errorf("SumSlice([]int64) overflow not detected")
	}
	if _, err := SumSlice([]int64{-9000000000000000000, -1000000000000000000}); err == nil {
		t.Errorf("SumSlice([]int64) underflow not detected")
	}

	// Test uint
	if res, err := SumSlice([]uint{1, 2}); err != nil || res != 3 {
		t.Errorf("SumSlice([]uint) failed: %v", err)
	}
	if _, err := SumSlice([]uint{^uint(0), 1}); err == nil {
		t.Errorf("SumSlice([]uint) overflow not detected")
	}

	// Test uint8
	if res, err := SumSlice([]uint8{1, 2}); err != nil || res != 3 {
		t.Errorf("SumSlice([]uint8) failed: %v", err)
	}
	if _, err := SumSlice([]uint8{250, 10}); err == nil {
		t.Errorf("SumSlice([]uint8) overflow not detected")
	}

	// Test uint16
	if res, err := SumSlice([]uint16{1, 2}); err != nil || res != 3 {
		t.Errorf("SumSlice([]uint16) failed: %v", err)
	}
	if _, err := SumSlice([]uint16{65000, 1000}); err == nil {
		t.Errorf("SumSlice([]uint16) overflow not detected")
	}

	// Test uint32
	if res, err := SumSlice([]uint32{1, 2}); err != nil || res != 3 {
		t.Errorf("SumSlice([]uint32) failed: %v", err)
	}
	if _, err := SumSlice([]uint32{4000000000, 3000000000}); err == nil {
		t.Errorf("SumSlice([]uint32) overflow not detected")
	}

	// Test uint64
	if res, err := SumSlice([]uint64{1, 2}); err != nil || res != 3 {
		t.Errorf("SumSlice([]uint64) failed: %v", err)
	}
	if _, err := SumSlice([]uint64{18000000000000000000, 1000000000000000000}); err == nil {
		t.Errorf("SumSlice([]uint64) overflow not detected")
	}

	// Test uintptr
	if res, err := SumSlice([]uintptr{1, 2}); err != nil || res != 3 {
		t.Errorf("SumSlice([]uintptr) failed: %v", err)
	}
	if _, err := SumSlice([]uintptr{^uintptr(0), 1}); err == nil {
		t.Errorf("SumSlice([]uintptr) overflow not detected")
	}
}

func TestMinSlice(t *testing.T) {
	// Empty slice
	if _, err := MinSlice([]int{}); err == nil {
		t.Errorf("MinSlice of empty slice should return an error")
	}

	// Int slice
	res, err := MinSlice([]int{5, 2, 9, 1, 7})
	if err != nil || res != 1 {
		t.Errorf("MinSlice([5,2,9,1,7]) = (%v, %v); want (1, nil)", res, err)
	}

	// String slice
	resStr, err := MinSlice([]string{"banana", "apple", "cherry"})
	if err != nil || resStr != "apple" {
		t.Errorf("MinSlice([banana,apple,cherry]) = (%v, %v); want (apple, nil)", resStr, err)
	}
}

func TestMaxSlice(t *testing.T) {
	// Empty slice
	if _, err := MaxSlice([]int{}); err == nil {
		t.Errorf("MaxSlice of empty slice should return an error")
	}

	// Int slice
	res, err := MaxSlice([]int{5, 2, 9, 1, 7})
	if err != nil || res != 9 {
		t.Errorf("MaxSlice([5,2,9,1,7]) = (%v, %v); want (9, nil)", res, err)
	}

	// String slice
	resStr, err := MaxSlice([]string{"banana", "apple", "cherry"})
	if err != nil || resStr != "cherry" {
		t.Errorf("MaxSlice([banana,apple,cherry]) = (%v, %v); want (cherry, nil)", resStr, err)
	}
}
