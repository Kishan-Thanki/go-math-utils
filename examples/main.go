package main

import (
	"fmt"

	"github.com/Kishan-Thanki/go-math-utils/mathutils"
)

func main() {
	// GCD (Greatest Common Divisor)
	fmt.Println("GCD of 56 and 98:", mathutils.GCD(56, 98)) // Should print 14

	// LCM (Least Common Multiple)
	fmt.Println("LCM of 56 and 98:", mathutils.LCM(56, 98)) // Should print 392

	// Factorial
	result, err := mathutils.Factorial(5)
	if err != nil {
		fmt.Println("Error calculating factorial:", err)
	} else {
		fmt.Println("Factorial of 5:", result) // Should print 120
	}

	// Prime Number Check
	fmt.Println("Is 29 prime?", mathutils.IsPrime(29)) // Should print true

	// Fibonacci Sequence
	fmt.Println("Fibonacci of 6:", mathutils.Fibonacci(6)) // Should print 8

	// Power function
	fmt.Println("2 raised to the power of 3:", mathutils.Power(2, 3)) // Should print 8

	// Prime Factors
	fmt.Println("Prime factors of 56:", mathutils.PrimeFactors(56)) // Should print [2 2 2 7]

	// Sum of Natural Numbers
	fmt.Println("Sum of first 5 natural numbers:", mathutils.SumNaturalNumbers(5)) // Should print 15
}
