package utils

import (
	"fmt"
	"sample_solid_go/solid/lsp/features"
)

type AreaCalculator struct {
	Shape features.Shape
}

func (ac AreaCalculator) Area() float64 {
	return ac.Shape.Area()
}

func _Fill(s string) string {

	switch s {
	case "":
		return "Unknown"

	default:
		return s
	}
}

func ShowArea(shapeType string, shape features.Shape) {
	fmt.Println("Area of", _Fill(shapeType), "is", shape.Area())
}
