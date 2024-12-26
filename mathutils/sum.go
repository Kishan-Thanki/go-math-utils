// sum.go
package mathutils

// SumNaturalNumbers returns the sum of the first n natural numbers (1 + 2 + ... + n).
func SumNaturalNumbers(n int) int {
	return n * (n + 1) / 2
}
