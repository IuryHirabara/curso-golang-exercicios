package main

import (
	"fmt"
)

func main() {
	// com contador
	// i := 0
	// for i < 10 {
	// time.Sleep(time.Second)
	// fmt.Println("adding to i")
	// i++
	// }
	// fmt.Println(i)

	// for com init
	// visivel só no for
	// for j := 0; j < 10; j++ {
	// fmt.Println("adding to j")
	// time.Sleep(time.Second)
	// }

	// for com clausula RANGE

	// iterando array
	names := [3]string{"Josh", "John", "Joseph"}
	for index, name := range names {
		fmt.Println("index: ", index, "  name: ", name)
	}

	// iterando só o nome
	for _, name := range names {
		fmt.Println(name)
	}

	// iterando slice
	other_names := []string{"Clodonildo", "Clebelsvaldo", "Roselvado", "Jurasdir"}
	for index, name := range other_names {
		fmt.Println(index, name)
	}

	// iterando palavras
	// o "name" será o número que corresponde a letra na tabela ASCII
	for index, name := range "WORD" {
		// usamos a função string() para converter o name que está em ASCII para seu respectivo caracter
		fmt.Println(index, name, string(name))
	}

	// iterando map
	user := map[string]string{
		"name":     "Leonhard",
		"lastname": "Godding",
	}
	for index, value := range user {
		fmt.Println(index, value)
	}

	// não há como fazer range em struct
}
