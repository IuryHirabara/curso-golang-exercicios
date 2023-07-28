package main

import "fmt"

// go é uma linguagem fortemente tipada, sua tipagem pode ser feita de maneira implicita ou explicita
func main() {
	var variable1 string = "variable 1" // declarando tipo explicitamente
	variable2 := "variable 2"           // implicitamente, o nome técnico é "inferencia de tipo"
	fmt.Println(variable1, variable2)

	// é possível também declarar multiplas variaveis, sem a necessidade de usar o var em cada uma delas
	// elas podem tem um valor ou não
	var (
		variable3 string
		variable4 string = "variable 4"
	)
	fmt.Println(variable3, variable4)

	// declaração de várias variáveis por inferência
	variable5, variable6 := "variable 5", "variable 6"
	fmt.Println(variable5, variable6)

	// constantes
	const constant1 string = "constant 1"
	fmt.Println(constant1)

	variable5, variable6 = variable6, variable5
}
