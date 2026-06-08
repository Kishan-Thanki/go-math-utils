package integer_test

import (
	"os"
	"strconv"

	"github.com/kishan-thanki/go-math-utils/integer"
)

func ExampleGCD() {
	printInt(int64(integer.GCD(56, 98)))
	// Output: 14
}

func ExampleLCM() {
	val, err := integer.LCM(56, 98)
	if err == nil {
		printInt(int64(val))
	}
	// Output: 392
}

func ExampleFactorial() {
	val, err := integer.Factorial(5)
	if err == nil {
		printInt(int64(val))
	}
	// Output: 120
}

func ExampleIsPrime() {
	printBool(integer.IsPrime(29))
	// Output: true
}

func ExamplePower() {
	val, err := integer.Power(2, 3)
	if err == nil {
		printInt(int64(val))
	}
	// Output: 8
}

func ExamplePrimeFactors() {
	factors, err := integer.PrimeFactors(56)
	if err == nil {
		printSlice(factors)
	}
	// Output: [2 2 2 7]
}

func ExampleEulerTotient() {
	phi, err := integer.EulerTotient(9)
	if err == nil {
		printInt(int64(phi))
	}
	// Output: 6
}

func ExampleDivisors() {
	divs, err := integer.Divisors(12)
	if err == nil {
		printSlice(divs)
	}
	// Output: [1 2 3 4 6 12]
}

func ExampleNextPrime() {
	printInt(int64(integer.NextPrime(14)))
	// Output: 17
}

func ExampleNthPrime() {
	p, err := integer.NthPrime(5)
	if err == nil {
		printInt(int64(p))
	}
	// Output: 11
}

func ExamplePerfectNumber() {
	printBool(integer.PerfectNumber(28))
	// Output: true
}

// Helpers to avoid using the "fmt" package for testable examples.

func printInt(v int64) {
	os.Stdout.WriteString(strconv.FormatInt(v, 10) + "\n")
}

func printBool(v bool) {
	os.Stdout.WriteString(strconv.FormatBool(v) + "\n")
}

func printSlice[T integer.Integer](slice []T) {
	var buf []byte
	buf = append(buf, '[')
	for idx, v := range slice {
		if idx > 0 {
			buf = append(buf, ' ')
		}
		buf = append(buf, strconv.FormatInt(int64(v), 10)...)
	}
	buf = append(buf, ']', '\n')
	os.Stdout.Write(buf)
}
