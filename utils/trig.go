package utils

import "math"

// Cot returns the cotangent of the
func Cot(angle float64) float64 {
	return 1 / (math.Tan(angle))
}
