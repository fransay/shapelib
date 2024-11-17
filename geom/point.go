package geom

type Pointer interface {
	Point2D | Point3D
}

// Point represents the fundamental point geometric type
type Point struct {
	geomm
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

// IsEmpty checks if p has a nil SRID, a negative SRID.
// or negative dimension
func (p *Point) IsEmpty() bool {
	if p.geomm.Srid.ID == 0 || p.geomm.Srid.ID < 0 {
		return true
	}
	return false
}

// IsPointDIMStand return a check bool if a point has standard dimension
func (p *Point) IsPointDIMStand() bool {
	if p.geomm.Dim.checkStandDIM() {
		return true
	}
	return false
}

// todos: implement SRID of geometries. especially point in this case
