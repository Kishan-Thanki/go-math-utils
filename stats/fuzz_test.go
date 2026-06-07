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
