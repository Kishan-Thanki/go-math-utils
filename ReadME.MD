# Go Math Utils

A collection of basic math utilities built in **Go**. This package provides common mathematical functions such as GCD, LCM, factorial, prime number checks, Fibonacci sequence, power, prime factors, and more.

## Features

- **GCD**: Calculate the Greatest Common Divisor of two integers.
- **LCM**: Calculate the Least Common Multiple of two integers.
- **Factorial**: Calculate the factorial of a non-negative integer.
- **IsPrime**: Check if a number is prime.
- **Fibonacci**: Generate the nth Fibonacci number.
- **Power**: Calculate `a` raised to the power of `b`.
- **PrimeFactors**: Find the prime factors of a number.
- **SumNaturalNumbers**: Calculate the sum of the first `n` natural numbers.

## Technologies Used

- **Go**: The programming language used to build the package.
- **No external dependencies**: This package is lightweight and relies solely on Go's standard libraries.

## Installation

To use **Go Math Utils** in your Go project, follow the steps below:

1. Initialize your Go module (if you haven't already):
   ```bash
   go mod init <your-module-name>

2. Get the **Go Math Utils** package:
   ```bash
   go get github.com/Kishan-Thanki/go-math-utils

3. Import the package into your project:
   import "github.com/Kishan-Thanki/go-math-utils/mathutils"

## Usage
Here’s how you can use the package in your Go project:

### Example Code
    
    package main

    import (
        "fmt"
        "github.com/your-username/go-math-utils/mathutils"
    )

    func main() {
        fmt.Println("GCD of 56 and 98:", mathutils.GCD(56, 98))             // Should print 14
        fmt.Println("LCM of 56 and 98:", mathutils.LCM(56, 98))             // Should print 392
        fmt.Println("Factorial of 5:", mathutils.Factorial(5))              // Should print 120
        fmt.Println("Is 29 prime?", mathutils.IsPrime(29))                  // Should print true
        fmt.Println("Fibonacci of 6:", mathutils.Fibonacci(6))              // Should print 8
        fmt.Println("2 raised to the power of 3:", mathutils.Power(2, 3))   // Should print 8
        fmt.Println("Prime factors of 56:", mathutils.PrimeFactors(56))     // Should print [2 2 2 7]
        fmt.Println("Sum of first 5 natural numbers:", mathutils.SumNaturalNumbers(5)) // Should print 15
    }
    


### Functions
GCD(a, b int) int
- Returns the Greatest Common Divisor of a and b using the Euclidean algorithm.

LCM(a, b int) int
- Returns the Least Common Multiple of a and b.

Factorial(n int) int
- Returns the factorial of n (n!).

IsPrime(n int) bool
- Returns true if n is a prime number, otherwise false.

Fibonacci(n int) int
- Returns the nth Fibonacci number using an iterative approach.

Power(a, b int) int
- Returns a raised to the power of b.

PrimeFactors(n int) []int
- Returns a slice containing the prime factors of n.

SumNaturalNumbers(n int) int
- Returns the sum of the first n natural numbers.


### Tests
To ensure the correctness of the math utilities, unit tests are provided in the utils_test.go file. You can run the tests using the following command:
   ```bash
   go test
   ```
