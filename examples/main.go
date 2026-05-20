package main

import (
	"fmt"

	"github.com/Kishan-Thanki/go-math-utils/mathutils"
)

func main() {
	// GCD (Greatest Common Divisor)
	fmt.Println("GCD of 56 and 98:", mathutils.GCD(56, 98)) // Should print 14

	// LCM (Least Common Multiple)
	if lcmVal, err := mathutils.LCM(56, 98); err != nil {
		fmt.Println("Error calculating LCM:", err)
	} else {
		fmt.Println("LCM of 56 and 98:", lcmVal) // Should print 392
	}

	// Factorial
	if factVal, err := mathutils.Factorial(5); err != nil {
		fmt.Println("Error calculating factorial:", err)
	} else {
		fmt.Println("Factorial of 5:", factVal) // Should print 120
	}

	// Prime Number Check
	fmt.Println("Is 29 prime?", mathutils.IsPrime(29)) // Should print true

	// Fibonacci Sequence
	if fibVal, err := mathutils.Fibonacci(6); err != nil {
		fmt.Println("Error calculating Fibonacci:", err)
	} else {
		fmt.Println("Fibonacci of 6:", fibVal) // Should print 8
	}

	// Power function
	if powVal, err := mathutils.Power(2, 3); err != nil {
		fmt.Println("Error calculating Power:", err)
	} else {
		fmt.Println("2 raised to the power of 3:", powVal) // Should print 8
	}

	// Prime Factors
	if factors, err := mathutils.PrimeFactors(56); err != nil {
		fmt.Println("Error calculating Prime Factors:", err)
	} else {
		fmt.Println("Prime factors of 56:", factors) // Should print [2 2 2 7]
	}

	// Sum of Natural Numbers
	if sumVal, err := mathutils.SumNaturalNumbers(5); err != nil {
		fmt.Println("Error calculating Sum of Natural Numbers:", err)
	} else {
		fmt.Println("Sum of first 5 natural numbers:", sumVal) // Should print 15
	}
}
