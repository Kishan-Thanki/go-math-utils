// power.go
package mathutils

// Power calculates a raised to the power of b (a^b).
func Power(a, b int) int {
	result := 1
	for b > 0 {
		result *= a
		b--
	}
	return result
}
