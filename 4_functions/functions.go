package main

import "fmt"

func sum(number1 int8, number2 int8) int8 {
	return number1 + number2
}

// podemos declarar o tipo dos parametros no ultimo parametro se os anteriores forem do mesmo tipo
// podemos ter multiplos retornos também
func multiCalc(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	subtraction := n1 - n2
	return sum, subtraction
}

func main() {
	result := sum(4, 5)
	fmt.Println(result)

	// declaração de variável do tipo função
	var function1 = func(text string) string {
		fmt.Println(text)
		return text
	}

	new_result := function1("This is a function into a variable!")
	fmt.Println(new_result)

	// a quantidade variáveis precisa ser igual a quantidade de retorno da função
	// para não usar uma variável, coloca-se um "_" no lugar da variável
	sum, _ := multiCalc(3, 6)
	fmt.Println(sum)
}
