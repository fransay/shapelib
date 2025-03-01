package geom

type SRID struct {
	// unique identifier
	ID int
}

// NewSRID constructs a new SRID
func NewSRID(id int) *SRID {
	return &SRID{
		ID: id,
	}
}

// SetID set sr object's id to new id
func (sr *SRID) SetID(id int) {
	sr.ID = id
}
