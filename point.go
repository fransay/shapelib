/*
A point is the most fundamental object in geometry with a zero-dimension.
In classical eclidean geometry, a point is a primitive notion that models 
an exact location in space.
*/
package siffa  

import "wkt.go"

// Create point type 
type Point struct{ 
	X float64
	Y float64
}

// A point has no size
func Size(point Point) float64{
	return 0.0
}

// A point has no length 
func length(point Point) float64{
	return 0.0
}

// A point has no width 
func width(point Point) float64{
	return 0.0
}

func Centriod(point Point) float64{
	return 0.0
}

// returns the dimensions of a point in a well known text format.
func Dim(point Point){

}