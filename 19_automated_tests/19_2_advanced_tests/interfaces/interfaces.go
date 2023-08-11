package interfaces

import (
	"math"
)

type Form interface {
	area() float64
}

type Rectangle struct {
	Heigth float64
	Width  float64
}

func (ret Rectangle) Area() float64 {
	return ret.Heigth * ret.Width
}

type Circle struct {
	Radius float64
}

func (circ Circle) Area() float64 {
	return math.Pi * math.Pow(circ.Radius, 2)
}
