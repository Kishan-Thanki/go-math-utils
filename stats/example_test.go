package stats_test

import (
	"os"
	"strconv"

	"github.com/kishan-thanki/go-math-utils/stats"
)

func ExampleAverage() {
	avg, err := stats.Average([]int{1, 2, 3, 4, 5})
	if err == nil {
		printFloat(avg)
	}
	// Output: 3
}

func ExampleMedian() {
	med, err := stats.Median([]float64{1.5, 5.5, 2.5, 3.5})
	if err == nil {
		printFloat(med)
	}
	// Output: 3
}

func ExampleMode() {
	mode, err := stats.Mode([]string{"apple", "banana", "banana", "cherry"})
	if err == nil {
		printString(mode)
	}
	// Output: banana
}

func ExampleSumSlice() {
	sum, err := stats.SumSlice([]int{10, 20, 30})
	if err == nil {
		printInt(int64(sum))
	}
	// Output: 60
}

func ExampleMinSlice() {
	min, err := stats.MinSlice([]int{5, 2, 9, 1, 7})
	if err == nil {
		printInt(int64(min))
	}
	// Output: 1
}

func ExampleMaxSlice() {
	max, err := stats.MaxSlice([]int{5, 2, 9, 1, 7})
	if err == nil {
		printInt(int64(max))
	}
	// Output: 9
}

// Helpers to avoid using the "fmt" package for testable examples.

func printFloat(v float64) {
	os.Stdout.WriteString(strconv.FormatFloat(v, 'g', -1, 64) + "\n")
}

func printInt(v int64) {
	os.Stdout.WriteString(strconv.FormatInt(v, 10) + "\n")
}

func printString(v string) {
	os.Stdout.WriteString(v + "\n")
}
