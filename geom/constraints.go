package geom

// IntFloat type constraints for all numeric data representation.
type IntFloat interface {
	int8 | int16 | int32 | int64 | float32 | float64
}
