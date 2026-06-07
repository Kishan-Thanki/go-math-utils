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

func ExampleEulerTotient() {
	phi, err := integer.EulerTotient(9)
	if err == nil {
		fmt.Println(phi)
	}
	// Output: 6
}

func ExampleDivisors() {
	divs, err := integer.Divisors(12)
	if err == nil {
		fmt.Println(divs)
	}
	// Output: [1 2 3 4 6 12]
}

func ExampleNextPrime() {
	fmt.Println(integer.NextPrime(14))
	// Output: 17
}

func ExampleNthPrime() {
	p, err := integer.NthPrime(5)
	if err == nil {
		fmt.Println(p)
	}
	// Output: 11
}
