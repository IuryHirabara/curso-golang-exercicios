package main

import "fmt"

type user struct {
	name    string
	years   uint8
	address address
}

type address struct {
	street string
	number uint8
}

func main() {
	// o valor zero de uma struct será correspondente ao tipo dos seus campos
	var u user
	u.name = "Cleberson"
	u.years = 20
	fmt.Println(u)

	// por inferencia
	address1 := address{"Jonesvaldo", 23}
	u2 := user{"Mococa", 23, address1}
	fmt.Println(u2)

	// passando somente um parâmetro
	u3 := user{years: 27}
	fmt.Println(u3)
	u3.name = "Davi"
	fmt.Println(u3)
}
