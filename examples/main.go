package main

import (
	"fmt"

	"github.com/fransay/shapelib/src/geom"
	"github.com/fransay/shapelib/src/shape"
)

func main() {
	// shapes
	circle := shape.Circle{Radius: 20, Centroid: geom.Point2D{X: 10, Y: 50}}
	square := shape.Square{Length: 5}
	fmt.Println("Area of a circle", circle.Area())
	fmt.Println("Area of a square", square.Area())
}
