package utils

import "math"

// Cot returns the cotangent of the
func Cot(angle float64) float64 { return 1 / (math.Tan(angle)) }

// SumOfInteriorAngle return the sum of interior angle of a regular polygon
func SumOfInteriorAngle(numberOfSide int) (intAngles int) {
	intAngles = (numberOfSide - 2) * 180
	return intAngles
}

// ExteriorAngle return the exterior angle regular polgyon given the number of sides
func ExteriorAngle(numberOfSides float64) float64 {
	return 360.0 / numberOfSides
}
