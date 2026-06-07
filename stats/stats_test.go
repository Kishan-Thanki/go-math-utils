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
