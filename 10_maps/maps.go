package main

import "fmt"

func main() {
	// criando map
	user := map[string]string{
		"name":     "Jonesvaldo",
		"lastname": "Cledenilson",
	}
	fmt.Println(user)
	fmt.Println(user["name"]) // acessando name

	// map aninhado
	user2 := map[string]map[string]string{
		"complete_name": {
			"name":     "Claudio",
			"lastname": "Abc",
		},
		"graduation": {
			"course": "science",
		},
	}
	fmt.Println(user2)
	fmt.Println(user2["complete_name"]["lastname"])

	// para deletar chaves
	delete(user2, "graduation")
	fmt.Println(user2)

	// adicionando valor no map
	user2["sign"] = map[string]string{
		"sign": "sagitarius",
	}
	fmt.Println(user2)
}
