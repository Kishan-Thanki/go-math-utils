package numeric_test

import (
	"fmt"

	"github.com/Kishan-Thanki/go-math-utils/numeric"
)

func ExampleMin() {
	fmt.Println(numeric.Min(42, 13))
	// Output: 13
}

func ExampleMax() {
	fmt.Println(numeric.Max(3.14, 2.71))
	// Output: 3.14
}

func ExampleClamp() {
	fmt.Println(numeric.Clamp(15, 10, 20))
	fmt.Println(numeric.Clamp(5, 10, 20))
	// Output:
	// 15
	// 10
}

func ExampleAbs() {
	fmt.Println(numeric.Abs(-42))
	// Output: 42
}

func ExampleSign() {
	fmt.Println(numeric.Sign(-5))
	fmt.Println(numeric.Sign(0))
	fmt.Println(numeric.Sign(5))
	// Output:
	// -1
	// 0
	// 1
}

func ExampleIsEven() {
	fmt.Println(numeric.IsEven(4))
	// Output: true
}

func ExampleIsOdd() {
	fmt.Println(numeric.IsOdd(4))
	// Output: false
}
