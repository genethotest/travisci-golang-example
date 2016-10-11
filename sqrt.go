// This package is a Travis-CI example package.
package newmath

// Sqrt returns an approximation to the square root of x.
func Sqrt(x float64) float64 {
	// This is a terrible implementation.
	// Real code should import "math" and use math.Sqrt.
	z := 0.0
	for m := 0; m < 1000; m++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}
