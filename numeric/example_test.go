package numeric_test

import (
	"os"
	"strconv"

	"github.com/kishan-thanki/go-math-utils/numeric"
)

func ExampleMin() {
	printInt(int64(numeric.Min(42, 13)))
	// Output: 13
}

func ExampleMax() {
	printFloat(numeric.Max(3.14, 2.71))
	// Output: 3.14
}

func ExampleClamp() {
	printInt(int64(numeric.Clamp(15, 10, 20)))
	printInt(int64(numeric.Clamp(5, 10, 20)))
	// Output:
	// 15
	// 10
}

func ExampleAbs() {
	printInt(int64(numeric.Abs(-42)))
	// Output: 42
}

func ExampleSign() {
	printInt(int64(numeric.Sign(-5)))
	printInt(int64(numeric.Sign(0)))
	printInt(int64(numeric.Sign(5)))
	// Output:
	// -1
	// 0
	// 1
}

func ExampleIsEven() {
	printBool(numeric.IsEven(4))
	// Output: true
}

func ExampleIsOdd() {
	printBool(numeric.IsOdd(4))
	// Output: false
}

// Helpers to avoid using the "fmt" package for testable examples.

func printInt(v int64) {
	os.Stdout.WriteString(strconv.FormatInt(v, 10) + "\n")
}

func printFloat(v float64) {
	os.Stdout.WriteString(strconv.FormatFloat(v, 'g', -1, 64) + "\n")
}

func printBool(v bool) {
	os.Stdout.WriteString(strconv.FormatBool(v) + "\n")
}
