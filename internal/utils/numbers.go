package utils

import "math"

// AddTwoInt returns the addition of two integers
func AddTwoInt(a, b int) int { return a + b }

// AddTwoFloat returns the addition of two floats :: 64 bit
func AddTwoFloat(a, b float64) float64 { return a + b }

// IsClose checks how close two numbers using a tolerance
func IsClose(a float64, b float64, tolerance float64) (isClose bool) {
	var diff = Diff(a, b)
	isClose = diff <= tolerance
	return isClose
}

// CompareThreeFloat64 returns an equality boolean check of three float64 types
func CompareThreeFloat64(a, b, c float64) (comp bool) {
	comp = (a == b) && (a == c) && (b == c)
	return comp
}

// CompareTwoFloat64 returns an equality boolean check of three float64 types
func CompareTwoFloat64(a, b float64) bool { return a == b }

// Diff returns the absolute difference between two floating point numbers
func Diff(a, b float64) (diff float64) {
	diff = math.Abs(b - a)
	return diff
}
