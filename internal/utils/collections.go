package utils

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

