package stats

import (
	"testing"
)

func FuzzAverage(f *testing.F) {
	f.Add([]byte{1, 2, 3})
	f.Fuzz(func(t *testing.T, data []byte) {
		slice := make([]int, len(data))
		for i, v := range data {
			slice[i] = int(v)
		}
		res, err := Average(slice)
		if len(slice) == 0 {
			if err == nil {
				t.Fatalf("Average of empty slice did not return error")
			}
			return
		}
		if err != nil {
			t.Fatalf("Average failed unexpectedly: %v", err)
		}
		if res < 0 || res > 255 {
			t.Fatalf("Average %f out of bounds for values [0, 255]", res)
		}
	})
}

func FuzzMedian(f *testing.F) {
	f.Add([]byte{1, 2, 3})
	f.Fuzz(func(t *testing.T, data []byte) {
		slice := make([]int, len(data))
		for i, v := range data {
			slice[i] = int(v)
		}
		res, err := Median(slice)
		if len(slice) == 0 {
			if err == nil {
				t.Fatalf("Median of empty slice did not return error")
			}
			return
		}
		if err != nil {
			t.Fatalf("Median failed unexpectedly: %v", err)
		}
		if res < 0 || res > 255 {
			t.Fatalf("Median %f out of bounds for values [0, 255]", res)
		}
	})
}

func FuzzMode(f *testing.F) {
	f.Add([]byte{1, 2, 3, 3})
	f.Fuzz(func(t *testing.T, data []byte) {
		slice := make([]int, len(data))
		for i, v := range data {
			slice[i] = int(v)
		}
		res, err := Mode(slice)
		if len(slice) == 0 {
			if err == nil {
				t.Fatalf("Mode of empty slice did not return error")
			}
			return
		}
		if err != nil {
			t.Fatalf("Mode failed unexpectedly: %v", err)
		}

		// Verify that the result is actually an element of the slice
		found := false
		for _, v := range slice {
			if v == res {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("Mode returned %d which is not present in slice %v", res, slice)
		}
	})
}

func FuzzSumSlice(f *testing.F) {
	f.Add([]byte{1, 2, 3})
	f.Fuzz(func(t *testing.T, data []byte) {
		slice := make([]int, len(data))
		for i, v := range data {
			slice[i] = int(v)
		}
		res, err := SumSlice(slice)
		if len(slice) == 0 {
			if err == nil {
				t.Fatalf("SumSlice of empty slice did not return error")
			}
			return
		}
		if err != nil {
			return
		}
		if res < 0 {
			t.Fatalf("SumSlice returned negative sum %d for positive inputs", res)
		}
	})
}

func FuzzMinSlice(f *testing.F) {
	f.Add([]byte{5, 2, 9})
	f.Fuzz(func(t *testing.T, data []byte) {
		slice := make([]int, len(data))
		for i, v := range data {
			slice[i] = int(v)
		}
		res, err := MinSlice(slice)
		if len(slice) == 0 {
			if err == nil {
				t.Fatalf("MinSlice of empty slice did not return error")
			}
			return
		}
		if err != nil {
			t.Fatalf("MinSlice failed unexpectedly: %v", err)
		}
		for _, v := range slice {
			if res > v {
				t.Fatalf("MinSlice returned %d, but found smaller value %d", res, v)
			}
		}
	})
}

func FuzzMaxSlice(f *testing.F) {
	f.Add([]byte{5, 2, 9})
	f.Fuzz(func(t *testing.T, data []byte) {
		slice := make([]int, len(data))
		for i, v := range data {
			slice[i] = int(v)
		}
		res, err := MaxSlice(slice)
		if len(slice) == 0 {
			if err == nil {
				t.Fatalf("MaxSlice of empty slice did not return error")
			}
			return
		}
		if err != nil {
			t.Fatalf("MaxSlice failed unexpectedly: %v", err)
		}
		for _, v := range slice {
			if res < v {
				t.Fatalf("MaxSlice returned %d, but found larger value %d", res, v)
			}
		}
	})
}
