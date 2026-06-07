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

