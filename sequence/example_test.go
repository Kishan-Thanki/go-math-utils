package sequence_test

import (
	"os"
	"strconv"

	"github.com/kishan-thanki/go-math-utils/sequence"
)

func ExampleFibonacci() {
	val, err := sequence.Fibonacci(6)
	if err == nil {
		printInt(int64(val))
	}
	// Output: 8
}

func ExampleSumNaturalNumbers() {
	val, err := sequence.SumNaturalNumbers(5)
	if err == nil {
		printInt(int64(val))
	}
	// Output: 15
}

func printInt(v int64) {
	os.Stdout.WriteString(strconv.FormatInt(v, 10) + "\n")
}
