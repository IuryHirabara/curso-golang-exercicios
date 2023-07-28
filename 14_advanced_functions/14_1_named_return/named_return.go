package main

import "fmt"

// multiplos retornos
// retorno nomeado
// com o retorno nomeado, não é necessário passar os parametros na clausula do return
func calc(number1 int, number2 int) (sum int, subtraction int) {
	sum = number1 + number2
	subtraction = number1 - number2
	return
}

func main() {
	sum, subtraction := calc(10, 20)
	fmt.Println(sum, subtraction)
}
