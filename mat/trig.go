package mat

import "math"

// Cot return the cotangent of an angle in degree decimals
func Cot(angle float64) float64 { return 1 / (math.Tan(angle)) }

// SumOfInteriorAngle return the sum of interior angle of a regular polygon
func SumOfInteriorAngle(numberOfSides int) (sumIntAngles int) {
	sumIntAngles = (numberOfSides - 2) * 180
	return sumIntAngles
}

// InteriorAngle return the interior angle of a regular polygon given the number of sides
func InteriorAngle(numberOfSides int) (intAngle float64) {
	sumIntAngle := SumOfInteriorAngle(numberOfSides)
	intAngle = float64(sumIntAngle) / float64(numberOfSides)
	return intAngle
}

// ExteriorAngle return the exterior angle of a regular polygon given the number of sides
func ExteriorAngle(numberOfSides float64) float64 {
	return 360.0 / numberOfSides
}

// SumOfExteriorAngle return the sum of exterior angle of every polygon is 360 degree
const SumOfExteriorAngle = 360
