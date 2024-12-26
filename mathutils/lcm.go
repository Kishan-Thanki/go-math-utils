// lcm.go
package mathutils

// LCM returns the Least Common Multiple of two integers.
func LCM(a, b int) int {
	return a * b / GCD(a, b) // Uses the GCD function to calculate the LCM
}
