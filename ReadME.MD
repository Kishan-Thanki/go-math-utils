# go-math-utils

**A collection of safe, overflow-proof mathematical utilities built in Go.**

[![Go CI](https://github.com/Kishan-Thanki/go-math-utils/actions/workflows/go.yml/badge.svg)](https://github.com/Kishan-Thanki/go-math-utils/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/Kishan-Thanki/go-math-utils.svg)](https://pkg.go.dev/github.com/Kishan-Thanki/go-math-utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/Kishan-Thanki/go-math-utils)](https://goreportcard.com/report/github.com/Kishan-Thanki/go-math-utils)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/Kishan-Thanki/go-math-utils/blob/main/LICENSE)

This package provides common mathematical functions — GCD, LCM, Factorial, Fibonacci, Power, IsPrime, PrimeFactors, and SumNaturalNumbers — all hardened against **integer overflows**, **division-by-zero**, and **infinite loops** that plague naïve implementations.

## Features

- **GCD**: Greatest Common Divisor via Euclidean algorithm. Handles negatives and `math.MinInt` safely.
- **LCM**: Least Common Multiple with division-before-multiply to prevent premature overflow.
- **Factorial**: Overflow-protected with platform-aware limits (n≤20 on 64-bit, n≤12 on 32-bit).
- **IsPrime**: Optimized step-by-6 trial division using `i <= n/i` guard to avoid `i*i` overflow on large primes.
- **Fibonacci**: Safe up to the 92nd term on 64-bit platforms.
- **Power**: Logarithmic binary exponentiation — O(log b) instead of O(b).
- **PrimeFactors**: Overflow-safe factorization returning `([]int, error)`.
- **SumNaturalNumbers**: Uses the Gauss formula with overflow-safe half-term multiplication.

## Performance

All core functions run with **zero allocations** and sub-10ns latency on Apple M2:

| Function             | Time/op   | Allocs/op |
|----------------------|-----------|-----------|
| `GCD`                | ~2 ns     | 0         |
| `LCM`                | ~6 ns     | 0         |
| `Factorial`          | ~5.5 ns   | 0         |
| `IsPrime` (997)      | ~9.5 ns   | 0         |
| `Fibonacci`          | ~14 ns    | 0         |
| `Power`              | ~8 ns     | 0         |
| `PrimeFactors`       | ~222 ns   | 4         |
| `SumNaturalNumbers`  | ~2 ns     | 0         |

## Technologies Used

- **Go**: The programming language used to build the package.
- **No external dependencies**: Relies solely on Go's standard library.

## Installation

1. Initialize your Go module (if you haven't already):
   ```bash
   go mod init <your-module-name>
   ```

2. Get the package:
   ```bash
   go get github.com/Kishan-Thanki/go-math-utils
   ```

3. Import it:
   ```go
   import "github.com/Kishan-Thanki/go-math-utils/mathutils"
   ```

## Usage

```go
package main

import (
	"fmt"

	"github.com/Kishan-Thanki/go-math-utils/mathutils"
)

func main() {
	// GCD — returns int
	fmt.Println("GCD of 56 and 98:", mathutils.GCD(56, 98)) // 14

	// LCM — returns (int, error)
	if lcmVal, err := mathutils.LCM(56, 98); err == nil {
		fmt.Println("LCM of 56 and 98:", lcmVal) // 392
	}

	// Factorial — returns (int, error)
	if factVal, err := mathutils.Factorial(5); err == nil {
		fmt.Println("Factorial of 5:", factVal) // 120
	}

	// IsPrime — returns bool
	fmt.Println("Is 29 prime?", mathutils.IsPrime(29)) // true

	// Fibonacci — returns (int, error)
	if fibVal, err := mathutils.Fibonacci(6); err == nil {
		fmt.Println("Fibonacci of 6:", fibVal) // 8
	}

	// Power — returns (int, error)
	if powVal, err := mathutils.Power(2, 3); err == nil {
		fmt.Println("2 raised to the power of 3:", powVal) // 8
	}

	// PrimeFactors — returns ([]int, error)
	if factors, err := mathutils.PrimeFactors(56); err == nil {
		fmt.Println("Prime factors of 56:", factors) // [2 2 2 7]
	}

	// SumNaturalNumbers — returns (int, error)
	if sumVal, err := mathutils.SumNaturalNumbers(5); err == nil {
		fmt.Println("Sum of first 5 natural numbers:", sumVal) // 15
	}
}
```

## API Reference

| Function | Signature | Notes |
|---|---|---|
| `GCD` | `GCD(a, b int) int` | Always non-negative; handles `math.MinInt` |
| `LCM` | `LCM(a, b int) (int, error)` | Returns `(0, nil)` if either input is 0 |
| `Factorial` | `Factorial(n int) (int, error)` | Error if n < 0 or result overflows |
| `IsPrime` | `IsPrime(n int) bool` | Overflow-safe for primes near `math.MaxInt` |
| `Fibonacci` | `Fibonacci(n int) (int, error)` | Error if n < 0 or result overflows |
| `Power` | `Power(a, b int) (int, error)` | Error if b < 0 or result overflows |
| `PrimeFactors` | `PrimeFactors(n int) ([]int, error)` | Error if n ≤ 1 |
| `SumNaturalNumbers` | `SumNaturalNumbers(n int) (int, error)` | Error if n < 0 or result overflows |

## Development

```bash
# Run all tests
make test

# Run benchmarks
make bench

# Check test coverage
make cover

# Format code
make fmt

# Vet code
make vet
```

## License

[MIT](./LICENSE)
