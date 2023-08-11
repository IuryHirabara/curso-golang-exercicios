package interfaces

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rec := Rectangle{10, 12}
		expectedArea := float64(120)
		receivedArea := rec.Area()

		if expectedArea != receivedArea {
			t.Errorf("The received area: %f doesn't match with the expected area: %f", receivedArea, receivedArea)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circ := Circle{10}
		expectedArea := float64(math.Pi * 100)
		receivedArea := circ.Area()

		if expectedArea != receivedArea {
			t.Fatal("Error")
		}
	})
}
