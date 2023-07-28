package main

import "fmt"

// declarando interface genérica
func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("test")
	generic(123)
	generic(12.4)

	// a função Println do pacote fmt recebe várias interfaces vazias

	map1 := map[interface{}]interface{}{
		1:     "var",
		123.3: true,
	}
	fmt.Println(map1)
}
