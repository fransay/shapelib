package utils

import "math"

// EqualSlice compare two slice of types float64
func EqualSlice(sliceOne, sliceTwo []float64) (equal bool) {
	if len(sliceOne) != len(sliceTwo) {
		equal = false
		return
	}

	for index, _ := range sliceOne {
		if sliceOne[index] != sliceTwo[index] {
			equal = false
			return
		}
	}
	equal = true
	return equal
}

// Rad2Deg convert radians to degree
func Rad2Deg(radians float64) (deg float64) {
	deg = radians * (180 / math.Pi)
	return deg
}

// Deg2Rad convert degrees to radians
func Deg2Rad(deg float64) (rad float64) {
	rad = deg * (math.Pi / 180)
	return rad
}

// Distance returns distance between two stations with coordinates in an array form
func Distance(stn1, stn2 []float64) (dist float64) {
	dist = math.Sqrt(math.Abs(stn2[0]-stn1[0]) + math.Abs(stn2[1]-stn1[1]))
	return dist
}

// Diff returns the absolute difference between two floating point numbers
func Diff(a, b float64) (diff float64) {
	diff = math.Abs(b - a)
	return diff
}

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

// IsArrayElementFloat64 checks if elements of type float64 in array are the same.
func IsArrayElementFloat64(arr []float64) bool {
	currentElement := arr[0]
	tolerance := 0.005
	for _, element := range arr {
		if !IsClose(element, currentElement, tolerance) {
			return false
		}
	}
	return false
}

func AddTwoInt(a, b int) int {
	return a + b
}
