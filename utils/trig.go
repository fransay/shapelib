package utils

import "math"

// Cot return the cotangent of an angle in degree decimals
func Cot(angle float64) float64 { return 1 / (math.Tan(angle)) }

// SumOfInteriorAngle return the sum of interior angle of a regular polygon
func SumOfInteriorAngle(numberOfSide int) (intAngles int) {
	intAngles = (numberOfSide - 2) * 180
	return intAngles
}

// ExteriorAngle return the exterior angle of a regular polygon given the number of sides
func ExteriorAngle(numberOfSides float64) float64 {
	return 360.0 / numberOfSides
}
