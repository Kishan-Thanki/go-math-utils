package main

import (
	"fmt"

	"github.com/Kishan-Thanki/go-math-utils/integer"
	"github.com/Kishan-Thanki/go-math-utils/numeric"
	"github.com/Kishan-Thanki/go-math-utils/sequence"
	"github.com/Kishan-Thanki/go-math-utils/stats"
)

func main() {
	// GCD (Greatest Common Divisor)
	fmt.Println("GCD of 56 and 98:", integer.GCD(56, 98)) // Should print 14
	fmt.Println("GCD of int8(12) and int8(18):", integer.GCD(int8(12), int8(18))) // Should print 6

	// LCM (Least Common Multiple)
	if lcmVal, err := integer.LCM(56, 98); err != nil {
		fmt.Println("Error calculating LCM:", err)
	} else {
		fmt.Println("LCM of 56 and 98:", lcmVal) // Should print 392
	}

	// Factorial
	if factVal, err := integer.Factorial(int16(5)); err != nil {
		fmt.Println("Error calculating factorial:", err)
	} else {
		fmt.Println("Factorial of int16(5):", factVal) // Should print 120
	}

	// Prime Number Check
	fmt.Println("Is 29 prime?", integer.IsPrime(29)) // Should print true

	// Fibonacci Sequence
	if fibVal, err := sequence.Fibonacci(uint32(6)); err != nil {
		fmt.Println("Error calculating Fibonacci:", err)
	} else {
		fmt.Println("Fibonacci of uint32(6):", fibVal) // Should print 8
	}

	// Power function
	if powVal, err := integer.Power(2, 3); err != nil {
		fmt.Println("Error calculating Power:", err)
	} else {
		fmt.Println("2 raised to the power of 3:", powVal) // Should print 8
	}

	// Prime Factors
	if factors, err := integer.PrimeFactors(56); err != nil {
		fmt.Println("Error calculating Prime Factors:", err)
	} else {
		fmt.Println("Prime factors of 56:", factors) // Should print [2 2 2 7]
	}

	// Sum of Natural Numbers
	if sumVal, err := sequence.SumNaturalNumbers(5); err != nil {
		fmt.Println("Error calculating Sum of Natural Numbers:", err)
	} else {
		fmt.Println("Sum of first 5 natural numbers:", sumVal) // Should print 15
	}

	// Numeric Helpers
	fmt.Println("Min of 10 and 20:", numeric.Min(10, 20))           // Should print 10
	fmt.Println("Max of 3.14 and 2.71:", numeric.Max(3.14, 2.71))   // Should print 3.14
	fmt.Println("Clamp 25 to [10, 20]:", numeric.Clamp(25, 10, 20)) // Should print 20
	fmt.Println("Abs of -42:", numeric.Abs(-42))                    // Should print 42
	fmt.Println("Sign of -15:", numeric.Sign(-15))                  // Should print -1
	fmt.Println("Is 4 even?", numeric.IsEven(4))                    // Should print true
	fmt.Println("Is 4 odd?", numeric.IsOdd(4))                      // Should print false

	// Stats Package
	intSlice := []int{1, 3, 3, 2, 3, 4, 5}
	avg, _ := stats.Average(intSlice)
	med, _ := stats.Median(intSlice)
	mode, _ := stats.Mode(intSlice)
	fmt.Printf("Stats on [1, 3, 3, 2, 3, 4, 5]: Avg: %.2f, Median: %.1f, Mode: %d\n", avg, med, mode)

	floatSlice := []float64{1.5, 3.5, 2.5}
	avgFloat, _ := stats.Average(floatSlice)
	fmt.Printf("Avg of [1.5, 3.5, 2.5]: %.2f\n", avgFloat)

	// Number Theory
	phi, _ := integer.EulerTotient(9)
	divs, _ := integer.Divisors(12)
	nextP := integer.NextPrime(14)
	nthP, _ := integer.NthPrime(5)
	fmt.Printf("EulerTotient(9): %d\n", phi)
	fmt.Printf("Divisors(12): %v\n", divs)
	fmt.Printf("NextPrime(14): %d\n", nextP)
	fmt.Printf("5th Prime: %d\n", nthP)
}
