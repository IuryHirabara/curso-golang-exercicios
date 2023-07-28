package main

import "fmt"

func invertOperand(number int) int {
	return number * -1
}

// não precisa de retorno pois está alterando a variável original
func invertOperandWithPointer(number *int) {
	*number *= -1
}

func main() {
	// number := 20
	// passando parâmetro por cópia
	// invertedNumber := invertOperand(number)
	// fmt.Println(number, invertedNumber)

	newNumber := 40
	fmt.Println(newNumber)
	// passando parâmetro por referência
	invertOperandWithPointer(&newNumber)
	fmt.Println(newNumber)
}
