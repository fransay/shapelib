package geom

type Pointer interface {
	Point2D | Point3D
}

// Point represents the fundamental point geometric type
type Point struct {
	geomt
}

// Area returns area of p which is always zero
// since a point is dimensionless.
func (p *Point) Area() float64 {
	return 0.0
}

// Length return the length of p which is alway zero
// since a point is dimensionless.
func (p *Point) Length() float64 {
	return 0.0
}

// IsEmpty checks if p has a nil SRID or dimension
func (p *Point) IsEmpty() bool {
	if p.geomt.Srid.ID == 0 {
		return true
	}
	return false
}

// todos: implement SRID of geometries. especially point in this case
