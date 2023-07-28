package interfaces

import (
	"fmt"
	"math"
)

type Form interface {
	area() float64
}

func WriteArea(receivedForm Form) {
	fmt.Printf("the area of form is %0.2f\n", receivedForm.area())
}

type Rectangle struct {
	Heigth float64
	Width  float64
}

func (ret Rectangle) area() float64 {
	return ret.Heigth * ret.Width
}

type Circle struct {
	radius float64
}

func (circ Circle) area() float64 {
	return math.Pi * math.Pow(circ.radius, 2)
}
