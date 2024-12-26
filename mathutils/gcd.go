// gcd.go
package mathutils

// GCD returns the Greatest Common Divisor of two integers using the Euclidean algorithm.
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
