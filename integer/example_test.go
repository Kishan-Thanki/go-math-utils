package integer_test

import (
	"fmt"

	"github.com/Kishan-Thanki/go-math-utils/integer"
)

func ExampleGCD() {
	fmt.Println(integer.GCD(56, 98))
	// Output: 14
}

func ExampleLCM() {
	val, err := integer.LCM(56, 98)
	if err == nil {
		fmt.Println(val)
	}
	// Output: 392
}

func ExampleFactorial() {
	val, err := integer.Factorial(5)
	if err == nil {
		fmt.Println(val)
	}
	// Output: 120
}

func ExampleIsPrime() {
	fmt.Println(integer.IsPrime(29))
	// Output: true
}

func ExamplePower() {
	val, err := integer.Power(2, 3)
	if err == nil {
		fmt.Println(val)
	}
	// Output: 8
}

func ExamplePrimeFactors() {
	factors, err := integer.PrimeFactors(56)
	if err == nil {
		fmt.Println(factors)
	}
	// Output: [2 2 2 7]
}
