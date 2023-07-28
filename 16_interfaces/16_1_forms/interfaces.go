package main

import (
	"fmt"
	"math"
)

// declarando interface
// não pode passar propriedades somente assinaturas de métodos (assinaturas de métodos descrevem como eles devem ser)
type form interface {
	// declarando método obrigatório para ser implementado nas sctructs rectangle e circle
	// elas terão um método chamado "area" que não irá receber nenhum parâmetros e irá retornar um float64
	area() float64
}

// com isso, é possível declarar uma só vez a função writeArea e ela ficará disponível tanto no rectangle quanto no circle
func writeArea(receivedForm form) {
	fmt.Printf("the area of form is %0.2f\n", receivedForm.area())
}

type rectangle struct {
	heigth float64
	width  float64
}

// implementando método area() para a struct "rectangle"
func (ret rectangle) area() float64 {
	return ret.heigth * ret.width
}

type circle struct {
	radius float64
}

// implementando método area() para a struct do "circle"
func (circ circle) area() float64 {
	return math.Pi * math.Pow(circ.radius, 2)
}

func main() {
	ret := rectangle{10, 15}
	// para chamar a função writeArea, é necessário implementar o método area
	writeArea(ret)

	circ := circle{10}
	writeArea(circ)
}
