package geom

// geomt defined a set of fundamental charateristics
// every geometry must have, e.g a spatial reference
// identification, dimensionality, etc.
type geomt struct {
	// spatial reference ID
	Srid SRID
	// dimension of geometry, i.e point = 0, line = 1, etc.
	Dim int
}
