package xrun

import (
	"fmt"
	"sample_solid_go/solid/lsp/features"
	"sample_solid_go/solid/lsp/impls"
	"sample_solid_go/solid/lsp/utils"
)

func LSP() {

	fmt.Println("Result of LSP:")

	var shape features.Shape = impls.Rectangle{Width: 6, Height: 5}

	// shape = impls.Rectangle{Width: 6, Height: 5}
	fmt.Println("Shape R", shape.Area())

	shape = impls.Circle{Radius: 7}
	fmt.Println("Shape C", shape.Area())

	rectangle1 := impls.Rectangle{Width: 5, Height: 5}
	ac := utils.AreaCalculator{Shape: rectangle1}
	fmt.Println("Rectangle1", ac.Area())
	// utils.ShowArea("Rectangle1", ac.Shape)

	rectangle2 := impls.Rectangle{Width: 10, Height: 5}
	utils.ShowArea("Rectangle2", rectangle2)

	circle1 := impls.Circle{Radius: 5}
	ac.Shape = circle1
	utils.ShowArea("Circle1", ac.Shape)

	circle2 := impls.Circle{Radius: 10}
	utils.ShowArea("Circle2", circle2)

	fmt.Println()
}
