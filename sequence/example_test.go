package sequence_test

import (
	"fmt"

	"github.com/Kishan-Thanki/go-math-utils/sequence"
)

func ExampleFibonacci() {
	val, err := sequence.Fibonacci(6)
	if err == nil {
		fmt.Println(val)
	}
	// Output: 8
}

func ExampleSumNaturalNumbers() {
	val, err := sequence.SumNaturalNumbers(5)
	if err == nil {
		fmt.Println(val)
	}
	// Output: 15
}
