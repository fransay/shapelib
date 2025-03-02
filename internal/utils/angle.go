package utils

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

