package stats_test

import (
	"fmt"

	"github.com/Kishan-Thanki/go-math-utils/stats"
)

func ExampleAverage() {
	avg, err := stats.Average([]int{1, 2, 3, 4, 5})
	if err == nil {
		fmt.Println(avg)
	}
	// Output: 3
}

func ExampleMedian() {
	med, err := stats.Median([]float64{1.5, 5.5, 2.5, 3.5})
	if err == nil {
		fmt.Println(med)
	}
	// Output: 3
}

func ExampleMode() {
	mode, err := stats.Mode([]string{"apple", "banana", "banana", "cherry"})
	if err == nil {
		fmt.Println(mode)
	}
	// Output: banana
}
