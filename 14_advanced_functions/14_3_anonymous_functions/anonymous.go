package main

import "fmt"

func main() {
	func() {
		fmt.Println("TEST")
	}()

	// com parametros
	func(text string) {
		fmt.Println(text)
	}("this is a text")

	// armazenando retorno
	text := func(text string) string {
		return fmt.Sprintf("Received -> %s", text)
	}("PARAMETER")
	fmt.Println(text)
}
