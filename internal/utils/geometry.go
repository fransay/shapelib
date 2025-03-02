package utils

// Distance returns distance between two stations with coordinates in an array form
func Distance(stn1, stn2 []float64) (dist float64) {
	dist = math.Sqrt(math.Abs(stn2[0]-stn1[0]) + math.Abs(stn2[1]-stn1[1]))
	return dist
}