package geog

// geomm defined a set of fundamental charateristics every geometry must have, e.g a spatial reference
// identification, dimensionality, etc.
type geomm struct {
	// spatial reference ID
	Srid SRID
	// dimension of geometry, i.e point = 0, line = 1, etc.
	Dim DIM
}

// DIM defines a dimension type
type DIM int

// New Dimensional type
func New(dm int) DIM {
	return DIM(dm)
}

// Check if dimension is a standard dimension
func (d *DIM) checkStandDIM() (isStand bool) {
	if *d < 0 || *d > 3 {
		isStand = false
	} else {
		isStand = true
	}
	return isStand
}
