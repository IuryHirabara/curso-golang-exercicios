package main

import "fmt"

func main() {
	number := 15
	// não precisa de parenteses
	if number > 10 {
		fmt.Println("greater than 10")
	} else {
		fmt.Println("less of equal to 10")
	}

	// if init
	// declarando variável na mesma linha em que há a estrutura de controle
	// fazendo deste modo, a variável limitada ao escopo do if else (ou estrutura de controle em que ela foi definida)
	if another_number := number; another_number < 0 {
		fmt.Println("greater than 0", another_number)
	} else if another_number == 15 {
		fmt.Println("It is equal to 15")
	}
}
